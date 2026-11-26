package bot

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-telegram/bot/models"
)

type getUpdatesClientFunc func(*http.Request) (*http.Response, error)

func (f getUpdatesClientFunc) Do(req *http.Request) (*http.Response, error) {
	// keep the body readable for req.FormValue
	body, err := io.ReadAll(req.Body)
	_ = req.Body.Close()
	if err != nil {
		return nil, err
	}
	req.Body = io.NopCloser(bytes.NewReader(body))
	return f(req)
}

// runGetUpdates polls like Start does and acts as the worker: it records every
// queued update and signals that it was dispatched.
func runGetUpdates(ctx context.Context, b *Bot) []*models.Update {
	queue := make(chan queuedUpdate)
	var delivered []*models.Update
	done := make(chan struct{})
	go func() {
		defer close(done)
		for item := range queue {
			delivered = append(delivered, item.update)
			item.claim.take()
		}
	}()
	b.getUpdates(ctx, queue)
	close(queue)
	<-done
	return delivered
}

func getUpdatesJSONResponse(body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

// One undecodable update must not poison the batch: the others are delivered,
// the offset moves past it and the failure is reported with the update id.
func Test_getUpdates_skipsUndecodableUpdate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	batch := `{"ok":true,"result":[
		{"update_id":100,"message":{"message_id":1,"text":"first"}},
		{"update_id":101,"message":"not an object"},
		{"update_id":102,"message":{"message_id":2,"text":"third"}}
	]}`

	var calls int32
	var errs []string
	b := &Bot{
		token:         "XXX",
		errorsHandler: func(handlerErr error) { errs = append(errs, handlerErr.Error()) },
		debugHandler:  func(string, ...any) {},
		client: getUpdatesClientFunc(func(*http.Request) (*http.Response, error) {
			if atomic.AddInt32(&calls, 1) > 1 {
				cancel()
				return nil, ctx.Err()
			}
			return getUpdatesJSONResponse(batch), nil
		}),
	}

	delivered := runGetUpdates(ctx, b)

	if got := atomic.LoadInt64(&b.lastUpdateID); got != 102 {
		t.Fatalf("offset must move past the bad update, lastUpdateID=%d", got)
	}
	if len(delivered) != 2 {
		t.Fatalf("expected 2 delivered updates, got %d", len(delivered))
	}
	if first, third := delivered[0], delivered[1]; first.ID != 100 || third.ID != 102 {
		t.Fatalf("expected updates 100 and 102, got %d and %d", first.ID, third.ID)
	}
	if len(errs) != 1 || !strings.Contains(errs[0], "101") || !strings.Contains(errs[0], "not an object") {
		t.Fatalf("undecodable update must be reported once with its id and raw payload, got %v", errs)
	}
}

// An update whose id cannot be read leaves the offset where it is, so the very same
// batch comes back on the next request. Without a backoff that is a tight loop.
func Test_getUpdates_backsOffWhenUpdateIDUnreadable(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	batch := `{"ok":true,"result":[{"update_id":"abc","message":"x"}]}`

	var calls int
	var firstCallAt time.Time
	var gap time.Duration

	b := &Bot{
		token:         "XXX",
		errorsHandler: func(error) {},
		debugHandler:  func(string, ...any) {},
		client: getUpdatesClientFunc(func(*http.Request) (*http.Response, error) {
			calls++
			switch calls {
			case 1:
				firstCallAt = time.Now()
			case 2:
				gap = time.Since(firstCallAt)
				cancel()
				return nil, ctx.Err()
			}
			return getUpdatesJSONResponse(batch), nil
		}),
	}

	runGetUpdates(ctx, b)

	if gap < 100*time.Millisecond {
		t.Fatalf("expected a backoff before the next request, got %v", gap)
	}
}

// pollGaps runs getUpdates against a client that answers every call with batch and
// returns the delays between consecutive requests, cancelling after calls requests.
func pollGaps(t *testing.T, batch string, calls int) []time.Duration {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var seen int
	var last time.Time
	var gaps []time.Duration

	b := &Bot{
		token:         "XXX",
		errorsHandler: func(error) {},
		debugHandler:  func(string, ...any) {},
		client: getUpdatesClientFunc(func(*http.Request) (*http.Response, error) {
			seen++
			if seen > 1 {
				gaps = append(gaps, time.Since(last))
			}
			last = time.Now()
			if seen == calls {
				cancel()
				return nil, ctx.Err()
			}
			return getUpdatesJSONResponse(batch), nil
		}),
	}

	runGetUpdates(ctx, b)

	return gaps
}

// While the offset stays stuck on the same unreadable batch the backoff must grow
// like it does for failed requests, not sit at the first step forever.
func Test_getUpdates_backoffGrowsWhileOffsetStuck(t *testing.T) {
	gaps := pollGaps(t, `{"ok":true,"result":[{"update_id":"abc","message":"x"}]}`, 4)

	if len(gaps) != 3 {
		t.Fatalf("expected 3 gaps, got %v", gaps)
	}
	for i, want := range []time.Duration{100 * time.Millisecond, 200 * time.Millisecond, 400 * time.Millisecond} {
		if gaps[i] < want {
			t.Fatalf("gap %d must be at least %v, got %v (all: %v)", i, want, gaps[i], gaps)
		}
	}
}

// Several unreadable elements in one batch are one stuck poll, not several: the
// backoff takes a single step per request.
func Test_getUpdates_backoffStepsOncePerBatch(t *testing.T) {
	batch := `{"ok":true,"result":[{"update_id":"a"},{"update_id":"b"},{"update_id":"c"},{"update_id":"d"}]}`
	gaps := pollGaps(t, batch, 2)

	if len(gaps) != 1 || gaps[0] >= 200*time.Millisecond {
		t.Fatalf("expected a single 100ms step, got %v", gaps)
	}
}

// When a later element of the batch advances the offset the unreadable one can never
// come back, so there is nothing to back off from.
func Test_getUpdates_noBackoffWhenLaterUpdateAdvancesOffset(t *testing.T) {
	batch := `{"ok":true,"result":[{"update_id":"abc","message":"x"},{"update_id":50,"message":{"message_id":1}}]}`
	gaps := pollGaps(t, batch, 2)

	if len(gaps) != 1 || gaps[0] >= 50*time.Millisecond {
		t.Fatalf("expected the next request without delay, got %v", gaps)
	}
}

// A null element decodes into an empty Update without an error. It must not be
// delivered and must not drag the offset back to zero.
func Test_getUpdates_ignoresUpdateWithoutID(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	batch := `{"ok":true,"result":[{"update_id":500,"message":{"message_id":1}},null,{}]}`

	var calls int32
	b := &Bot{
		token:         "XXX",
		errorsHandler: func(error) {},
		debugHandler:  func(string, ...any) {},
		client: getUpdatesClientFunc(func(*http.Request) (*http.Response, error) {
			if atomic.AddInt32(&calls, 1) > 1 {
				cancel()
				return nil, ctx.Err()
			}
			return getUpdatesJSONResponse(batch), nil
		}),
	}

	delivered := runGetUpdates(ctx, b)

	if got := atomic.LoadInt64(&b.lastUpdateID); got != 500 {
		t.Fatalf("offset must stay at the last real update, lastUpdateID=%d", got)
	}
	if len(delivered) != 1 {
		t.Fatalf("expected only the real update to be delivered, got %d", len(delivered))
	}
}

func waitForUpdateTest[T any](t *testing.T, ch <-chan T, event string) T {
	t.Helper()
	select {
	case value := <-ch:
		return value
	case <-time.After(5 * time.Second):
		t.Fatalf("timed out waiting for %s", event)
		var zero T
		return zero
	}
}

// A batch is confirmed once its asynchronous handlers start, so a handler can
// wait for an update from a later batch.
func TestBot_StartPollsWhileHandlerRuns(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	secondHandled := make(chan struct{})
	firstSawSecond := make(chan bool, 1)
	offsets := make(chan string, 3)
	var calls int
	b, err := New("xxx", WithSkipGetMe(),
		WithDefaultHandler(func(_ context.Context, _ *Bot, update *models.Update) {
			if update.ID == 2 {
				close(secondHandled)
				return
			}
			select {
			case <-secondHandled:
				firstSawSecond <- true
			case <-time.After(2 * time.Second):
				firstSawSecond <- false
			}
		}),
		WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
			calls++
			offsets <- req.FormValue("offset")
			switch calls {
			case 1:
				return getUpdatesJSONResponse(`{"ok":true,"result":[{"update_id":1}]}`), nil
			case 2:
				return getUpdatesJSONResponse(`{"ok":true,"result":[{"update_id":2}]}`), nil
			}
			<-req.Context().Done()
			return nil, req.Context().Err()
		})),
	)
	if err != nil {
		t.Fatal(err)
	}

	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		b.Start(ctx)
	}()
	if !waitForUpdateTest(t, firstSawSecond, "first handler to finish") {
		t.Error("the second update was not dispatched while the first handler was running")
	}
	cancel()
	waitForUpdateTest(t, stopped, "polling to stop")
	if first, second := <-offsets, <-offsets; first != "1" || second != "2" {
		t.Errorf("poll offsets = %q, %q, want 1, 2", first, second)
	}
}

// Cancellation confirms the updates before the first undispatched one, so
// polling resumes there without replaying the handled update.
func TestBot_StartResumesFromFirstUndispatchedUpdate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resumeCtx, cancelResume := context.WithCancel(context.Background())
	defer cancelResume()

	firstStarted := make(chan struct{})
	offsets := make(chan string, 4)
	var handled []int64
	var calls int
	var restarting bool
	b, err := New("xxx", WithSkipGetMe(), WithNotAsyncHandlers(),
		WithDefaultHandler(func(ctx context.Context, _ *Bot, update *models.Update) {
			handled = append(handled, update.ID)
			if !restarting && update.ID == 40 {
				close(firstStarted)
				<-ctx.Done()
			}
		}),
		WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
			calls++
			offsets <- req.FormValue("offset")
			if calls == 1 {
				if req.FormValue("offset") == "41" {
					return getUpdatesJSONResponse(`{"ok":true,"result":[{"update_id":42}]}`), nil
				}
				return getUpdatesJSONResponse(`{"ok":true,"result":[{"update_id":40},{"update_id":42}]}`), nil
			}
			<-req.Context().Done()
			return nil, req.Context().Err()
		})),
	)
	if err != nil {
		t.Fatal(err)
	}

	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		b.Start(ctx)
	}()
	waitForUpdateTest(t, firstStarted, "first handler to start")
	cancel()
	waitForUpdateTest(t, stopped, "polling to stop")
	if got := waitForUpdateTest(t, offsets, "first poll"); got != "1" {
		t.Errorf("first poll offset = %q, want 1", got)
	}

	restarting = true
	calls = 0
	resumed := make(chan struct{})
	go func() {
		defer close(resumed)
		b.Start(resumeCtx)
	}()
	if got := waitForUpdateTest(t, offsets, "resumed poll"); got != "41" {
		t.Errorf("resumed poll offset = %q, want 41", got)
	}
	if got := waitForUpdateTest(t, offsets, "poll after the resumed batch"); got != "43" {
		t.Errorf("offset after the resumed batch = %q, want 43", got)
	}
	cancelResume()
	waitForUpdateTest(t, resumed, "resumed polling to stop")
	if want := []int64{40, 42}; !reflect.DeepEqual(handled, want) {
		t.Errorf("handled updates = %v, want %v", handled, want)
	}
}

func TestBot_StartKeepsOffsetWhenLastDispatchedHandlerCancels(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	offsets := make(chan string, 2)
	var calls int
	b, err := New("xxx", WithSkipGetMe(), WithNotAsyncHandlers(),
		WithDefaultHandler(func(context.Context, *Bot, *models.Update) {
			cancel()
		}),
		WithErrorsHandler(func(error) {
			// The malformed record keeps the poller from reading the dispatch
			// signal until after the valid update's handler cancels the run.
			<-ctx.Done()
		}),
		WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
			offsets <- req.FormValue("offset")
			calls++
			if calls == 1 {
				return getUpdatesJSONResponse(`{"ok":true,"result":[{"update_id":42},{"update_id":"invalid"}]}`), nil
			}
			return nil, context.Canceled
		})))
	if err != nil {
		t.Fatal(err)
	}

	b.Start(ctx)
	if got := waitForUpdateTest(t, offsets, "initial poll"); got != "1" {
		t.Errorf("initial offset = %q, want 1", got)
	}

	b.Start(context.Background())
	if got := waitForUpdateTest(t, offsets, "poll after restart"); got != "43" {
		t.Errorf("offset after restart = %q, want 43", got)
	}
}

func TestBot_StartWaitsForDispatchedHandlers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	started := make(chan struct{})
	release := make(chan struct{})
	var calls int
	b, err := New("xxx", WithSkipGetMe(),
		WithDefaultHandler(func(context.Context, *Bot, *models.Update) {
			close(started)
			// The handler ignores cancellation on purpose.
			<-release
		}),
		WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
			calls++
			if calls == 1 {
				return getUpdatesJSONResponse(`{"ok":true,"result":[{"update_id":1}]}`), nil
			}
			<-req.Context().Done()
			return nil, req.Context().Err()
		})),
	)
	if err != nil {
		t.Fatal(err)
	}

	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		b.Start(ctx)
	}()
	waitForUpdateTest(t, started, "handler to start")
	cancel()
	select {
	case <-stopped:
		t.Fatal("Start returned while a dispatched handler was running")
	case <-time.After(100 * time.Millisecond):
	}
	close(release)
	waitForUpdateTest(t, stopped, "Start to return after the handler")
}

// An update whose preparation exits counts as dispatched, so polling goes on
// with the remaining worker instead of replaying it.
func TestBot_StartConfirmsUpdateWhosePreparationExits(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	offsets := make(chan string, 4)
	var calls int32
	b, err := New("xxx", WithSkipGetMe(), WithWorkers(2), WithErrorsHandler(func(error) {}),
		WithDefaultHandler(func(context.Context, *Bot, *models.Update) {}),
		WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
			offsets <- req.FormValue("offset")
			if atomic.AddInt32(&calls, 1) == 1 {
				return getUpdatesJSONResponse(`{"ok":true,"result":[{"update_id":7},{"update_id":8}]}`), nil
			}
			<-req.Context().Done()
			return nil, req.Context().Err()
		})))
	if err != nil {
		t.Fatal(err)
	}
	b.RegisterHandlerMatchFunc(func(update *models.Update) bool {
		if update.ID == 8 {
			runtime.Goexit()
		}
		return false
	}, b.defaultHandlerFunc)
	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		b.Start(ctx)
	}()
	for _, want := range []string{"1", "9"} {
		if got := waitForUpdateTest(t, offsets, "poll"); got != want {
			t.Errorf("poll offset = %q, want %q", got, want)
		}
	}
	cancel()
	waitForUpdateTest(t, stopped, "polling to stop")
}

// When the only worker exits during preparation the run stops, and the next
// run resumes after that update instead of replaying it forever.
func TestBot_StartResumesAfterUpdateWhosePreparationExits(t *testing.T) {
	b, err := New("xxx", WithSkipGetMe(), WithWorkers(1), WithNotAsyncHandlers(), WithErrorsHandler(func(error) {}),
		WithDefaultHandler(func(context.Context, *Bot, *models.Update) {}),
		WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
			if req.FormValue("offset") == "1" {
				return getUpdatesJSONResponse(`{"ok":true,"result":[{"update_id":7},{"update_id":8},{"update_id":9}]}`), nil
			}
			<-req.Context().Done()
			return nil, req.Context().Err()
		})))
	if err != nil {
		t.Fatal(err)
	}
	b.RegisterHandlerMatchFunc(func(update *models.Update) bool {
		if update.ID == 8 {
			runtime.Goexit()
		}
		return false
	}, b.defaultHandlerFunc)
	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		b.Start(context.Background())
	}()
	waitForUpdateTest(t, stopped, "polling to stop after its only worker exits")
	if got := atomic.LoadInt64(&b.lastUpdateID); got != 8 {
		t.Errorf("last update ID = %d, want 8", got)
	}
}

// A negative offset makes Telegram drop older updates. Polling must resume
// from the batch it got, so an update not dispatched before shutdown is
// delivered after a restart.
func TestBot_StartWithNegativeOffsetKeepsUndispatchedUpdate(t *testing.T) {
	var mx sync.Mutex
	pending := []int64{100, 101}
	var handled []int64
	firstStarted := make(chan struct{})
	restarted := false
	b, err := New("xxx", WithSkipGetMe(), WithNotAsyncHandlers(), WithInitialOffset(-3), WithErrorsHandler(func(error) {}),
		WithDefaultHandler(func(ctx context.Context, _ *Bot, update *models.Update) {
			mx.Lock()
			handled = append(handled, update.ID)
			first := !restarted && update.ID == 100
			mx.Unlock()
			if first {
				close(firstStarted)
				<-ctx.Done()
			}
		}),
		WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
			offset, _ := strconv.ParseInt(req.FormValue("offset"), 10, 64)
			mx.Lock()
			// Follow the getUpdates offset rules of the Bot API.
			if offset < 0 && len(pending) > int(-offset) {
				pending = pending[len(pending)-int(-offset):]
			}
			var result []string
			kept := pending[:0]
			for _, id := range pending {
				if offset <= 0 || id >= offset {
					kept = append(kept, id)
					result = append(result, fmt.Sprintf(`{"update_id":%d}`, id))
				}
			}
			pending = kept
			mx.Unlock()
			if len(result) == 0 {
				<-req.Context().Done()
				return nil, req.Context().Err()
			}
			return getUpdatesJSONResponse(`{"ok":true,"result":[` + strings.Join(result, ",") + `]}`), nil
		})))
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		b.Start(ctx)
	}()
	waitForUpdateTest(t, firstStarted, "first handler to start")
	cancel()
	waitForUpdateTest(t, stopped, "polling to stop")

	mx.Lock()
	restarted = true
	pending = append(pending, 102, 103)
	mx.Unlock()
	resumeCtx, cancelResume := context.WithCancel(context.Background())
	defer cancelResume()
	resumed := make(chan struct{})
	go func() {
		defer close(resumed)
		b.Start(resumeCtx)
	}()
	deadline := time.Now().Add(5 * time.Second)
	for {
		mx.Lock()
		n := len(handled)
		mx.Unlock()
		if n >= 4 || time.Now().After(deadline) {
			break
		}
		time.Sleep(time.Millisecond)
	}
	cancelResume()
	waitForUpdateTest(t, resumed, "resumed polling to stop")
	mx.Lock()
	defer mx.Unlock()
	if want := []int64{100, 101, 102, 103}; !reflect.DeepEqual(handled, want) {
		t.Errorf("handled updates = %v, want %v", handled, want)
	}
}

// A batch fetched with a negative offset moves the cursor to the batch before
// any update is dispatched.
func Test_getUpdates_negativeOffsetResumesFromBatch(t *testing.T) {
	requested := make(chan struct{})
	b, err := New("xxx", WithSkipGetMe(), WithInitialOffset(-3),
		WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
			if req.FormValue("offset") != "-2" {
				<-req.Context().Done()
				return nil, req.Context().Err()
			}
			close(requested)
			return getUpdatesJSONResponse(`{"ok":true,"result":[{"update_id":100},{"update_id":101}]}`), nil
		})))
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		// Nobody reads the queue, so no update is dispatched.
		b.getUpdates(ctx, make(chan queuedUpdate))
	}()
	waitForUpdateTest(t, requested, "poll with the negative offset")
	cancel()
	waitForUpdateTest(t, stopped, "polling to stop")
	if got := atomic.LoadInt64(&b.lastUpdateID); got != 99 {
		t.Errorf("last update ID = %d, want 99", got)
	}
}

func TestBot_StartWithoutWorkersReportsError(t *testing.T) {
	reported := make(chan error, 2)
	b, err := New("xxx", WithSkipGetMe(), WithWorkers(0),
		WithErrorsHandler(func(err error) { reported <- err }))
	if err != nil {
		t.Fatal(err)
	}

	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		b.Start(context.Background())
	}()
	waitForUpdateTest(t, stopped, "Start without workers to return")
	if got := waitForUpdateTest(t, reported, "workers error"); !strings.Contains(got.Error(), "without update workers") {
		t.Errorf("reported %v, want the missing workers error", got)
	}
}

// Cancellation while a matcher runs must not leave the update unconfirmed:
// its handler still runs, so replaying it would run it twice.
func TestBot_StartConfirmsUpdateCanceledDuringPreparation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	matching := make(chan struct{})
	release := make(chan struct{})
	var handled int32
	b, err := New("xxx", WithSkipGetMe(), WithWorkers(1), WithErrorsHandler(func(error) {}),
		WithDefaultHandler(func(context.Context, *Bot, *models.Update) { atomic.AddInt32(&handled, 1) }),
		WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
			if req.FormValue("offset") == "1" {
				return getUpdatesJSONResponse(`{"ok":true,"result":[{"update_id":7}]}`), nil
			}
			<-req.Context().Done()
			return nil, req.Context().Err()
		})))
	if err != nil {
		t.Fatal(err)
	}
	b.RegisterHandlerMatchFunc(func(*models.Update) bool {
		close(matching)
		<-release
		return false
	}, b.defaultHandlerFunc)
	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		b.Start(ctx)
	}()
	waitForUpdateTest(t, matching, "matcher to run")
	cancel()
	close(release)
	waitForUpdateTest(t, stopped, "polling to stop")
	if got := atomic.LoadInt32(&handled); got != 1 {
		t.Errorf("handled %d times, want 1", got)
	}
	if got := atomic.LoadInt64(&b.lastUpdateID); got != 7 {
		t.Errorf("last update ID = %d, want 7", got)
	}
}

// Cancellation while the poller waits to queue the next update still confirms
// the updates already dispatched.
func TestBot_StartConfirmsDispatchedUpdatesWhenCanceledWhileQueuing(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	firstStarted := make(chan struct{})
	b, err := New("xxx", WithSkipGetMe(), WithWorkers(1), WithNotAsyncHandlers(), WithUpdatesChannelCap(0),
		WithDefaultHandler(func(ctx context.Context, _ *Bot, update *models.Update) {
			if update.ID == 1 {
				close(firstStarted)
				<-ctx.Done()
			}
		}),
		WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
			if req.FormValue("offset") == "1" {
				return getUpdatesJSONResponse(`{"ok":true,"result":[{"update_id":1},{"update_id":2}]}`), nil
			}
			<-req.Context().Done()
			return nil, req.Context().Err()
		})))
	if err != nil {
		t.Fatal(err)
	}
	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		b.Start(ctx)
	}()
	waitForUpdateTest(t, firstStarted, "first handler to start")
	cancel()
	waitForUpdateTest(t, stopped, "polling to stop")
	if got := atomic.LoadInt64(&b.lastUpdateID); got != 1 {
		t.Errorf("last update ID = %d, want 1", got)
	}
}

func TestBot_StartShutdownGraceAllowsInFlightOutboundCall(t *testing.T) {
	type ctxKey struct{}
	parent := context.WithValue(context.Background(), ctxKey{}, "trace-123")
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	started := make(chan struct{})
	allowSend := make(chan struct{})
	secondPollStarted := make(chan struct{})
	secondPollCanceled := make(chan struct{})
	sendErrCh := make(chan error, 1)
	ctxValCh := make(chan any, 1)
	var pollCalls int32

	b, err := New("xxx", WithSkipGetMe(), WithShutdownGrace(500*time.Millisecond),
		WithDefaultHandler(func(hCtx context.Context, bot *Bot, _ *models.Update) {
			ctxValCh <- hCtx.Value(ctxKey{})
			close(started)
			<-allowSend
			_, sendErr := bot.SendMessage(hCtx, &SendMessageParams{ChatID: int64(1), Text: "done"})
			sendErrCh <- sendErr
		}),
		WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
			if strings.HasSuffix(req.URL.Path, "/sendMessage") {
				if err := req.Context().Err(); err != nil {
					return nil, err
				}
				return getUpdatesJSONResponse(`{"ok":true,"result":{"message_id":10}}`), nil
			}
			if atomic.AddInt32(&pollCalls, 1) == 1 {
				return getUpdatesJSONResponse(`{"ok":true,"result":[{"update_id":1}]}`), nil
			}
			close(secondPollStarted)
			<-req.Context().Done()
			close(secondPollCanceled)
			return nil, req.Context().Err()
		})))
	if err != nil {
		t.Fatal(err)
	}

	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		b.Start(ctx)
	}()

	waitForUpdateTest(t, started, "handler to start")
	waitForUpdateTest(t, secondPollStarted, "second poll to start")
	cancel()
	waitForUpdateTest(t, secondPollCanceled, "polling request to cancel immediately")

	select {
	case <-stopped:
		t.Fatal("Start returned before in-flight handler completed within shutdown grace")
	case <-time.After(30 * time.Millisecond):
	}

	close(allowSend)
	waitForUpdateTest(t, stopped, "Start to return after handler finished")

	if gotErr := waitForUpdateTest(t, sendErrCh, "sendMessage result"); gotErr != nil {
		t.Fatalf("SendMessage during shutdown grace failed: %v", gotErr)
	}
	if gotVal := waitForUpdateTest(t, ctxValCh, "handler context value"); gotVal != "trace-123" {
		t.Fatalf("handler context value = %v, want trace-123", gotVal)
	}
}

func TestBot_StartShutdownGraceCancelsHandlerAfterGrace(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	started := make(chan struct{})
	handlerExited := make(chan error, 1)
	var pollCalls int32

	b, err := New("xxx", WithSkipGetMe(), WithShutdownGrace(50*time.Millisecond),
		WithDefaultHandler(func(hCtx context.Context, _ *Bot, _ *models.Update) {
			close(started)
			<-hCtx.Done()
			handlerExited <- hCtx.Err()
		}),
		WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
			if atomic.AddInt32(&pollCalls, 1) == 1 {
				return getUpdatesJSONResponse(`{"ok":true,"result":[{"update_id":1}]}`), nil
			}
			<-req.Context().Done()
			return nil, req.Context().Err()
		})))
	if err != nil {
		t.Fatal(err)
	}

	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		b.Start(ctx)
	}()

	waitForUpdateTest(t, started, "handler to start")
	cancel()

	select {
	case err := <-handlerExited:
		t.Fatalf("handler context canceled before grace period elapsed: %v", err)
	case <-time.After(15 * time.Millisecond):
	}

	if gotErr := waitForUpdateTest(t, handlerExited, "handler context to cancel after grace"); gotErr == nil {
		t.Fatal("expected non-nil context error after grace period")
	}
	waitForUpdateTest(t, stopped, "Start to return after grace")
}

func TestBot_StartShutdownGraceBoundsStuckHandler(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	started := make(chan struct{})
	release := make(chan struct{})
	var releaseOnce sync.Once
	releaseHandler := func() { releaseOnce.Do(func() { close(release) }) }
	defer releaseHandler()
	handlerCtxCh := make(chan context.Context, 1)
	var pollCalls int32

	b, err := New("xxx", WithSkipGetMe(), WithShutdownGrace(50*time.Millisecond),
		WithDefaultHandler(func(hCtx context.Context, _ *Bot, update *models.Update) {
			if update.ID == 1 {
				handlerCtxCh <- hCtx
				close(started)
				<-release
			}
		}),
		WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
			if atomic.AddInt32(&pollCalls, 1) == 1 {
				return getUpdatesJSONResponse(`{"ok":true,"result":[{"update_id":1}]}`), nil
			}
			<-req.Context().Done()
			return nil, req.Context().Err()
		})))
	if err != nil {
		t.Fatal(err)
	}

	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		b.Start(ctx)
	}()

	waitForUpdateTest(t, started, "stuck handler to start")
	hCtx := waitForUpdateTest(t, handlerCtxCh, "handler context")
	cancel()

	waitForUpdateTest(t, stopped, "Start to return after grace timeout even with stuck handler")
	if hCtx.Err() == nil {
		t.Fatal("expected handler context to be canceled when shutdown grace expired")
	}

	// Ensure a subsequent Start run can begin cleanly after the previous run hit grace timeout.
	ctx2, cancel2 := context.WithCancel(context.Background())
	stopped2 := make(chan struct{})
	go func() {
		defer close(stopped2)
		b.Start(ctx2)
	}()
	cancel2()
	waitForUpdateTest(t, stopped2, "second Start to return cleanly")
}

func TestBot_StartShutdownGraceWithNotAsyncHandlers(t *testing.T) {
	for _, stuck := range []bool{false, true} {
		name := "completes"
		if stuck {
			name = "bounds_stuck"
		}
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			started := make(chan struct{})
			release := make(chan struct{})
			var releaseOnce sync.Once
			releaseHandler := func() { releaseOnce.Do(func() { close(release) }) }
			defer releaseHandler()
			sendErrCh := make(chan error, 1)
			var pollCalls int32

			b, err := New("xxx", WithSkipGetMe(), WithNotAsyncHandlers(), WithShutdownGrace(50*time.Millisecond),
				WithDefaultHandler(func(hCtx context.Context, bot *Bot, _ *models.Update) {
					close(started)
					<-release
					if !stuck {
						_, sendErr := bot.SendMessage(hCtx, &SendMessageParams{ChatID: int64(1), Text: "sync"})
						sendErrCh <- sendErr
					}
				}),
				WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
					if strings.HasSuffix(req.URL.Path, "/sendMessage") {
						if err := req.Context().Err(); err != nil {
							return nil, err
						}
						return getUpdatesJSONResponse(`{"ok":true,"result":{"message_id":11}}`), nil
					}
					if atomic.AddInt32(&pollCalls, 1) == 1 {
						return getUpdatesJSONResponse(`{"ok":true,"result":[{"update_id":1}]}`), nil
					}
					<-req.Context().Done()
					return nil, req.Context().Err()
				})))
			if err != nil {
				t.Fatal(err)
			}

			stopped := make(chan struct{})
			go func() {
				defer close(stopped)
				b.Start(ctx)
			}()

			waitForUpdateTest(t, started, "sync handler to start")
			cancel()

			if !stuck {
				time.Sleep(15 * time.Millisecond)
				releaseHandler()
				waitForUpdateTest(t, stopped, "Start to return after sync handler completed")
				if gotErr := waitForUpdateTest(t, sendErrCh, "sync sendMessage result"); gotErr != nil {
					t.Fatalf("SendMessage during grace failed: %v", gotErr)
				}
			} else {
				waitForUpdateTest(t, stopped, "Start to return after grace timeout with stuck sync handler")
			}
		})
	}
}
