package bot

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-telegram/bot/models"
)

type getUpdatesClientFunc func(*http.Request) (*http.Response, error)

func (f getUpdatesClientFunc) Do(req *http.Request) (*http.Response, error) {
	// keep the body readable for req.FormValue
	body, _ := io.ReadAll(req.Body)
	_ = req.Body.Close()
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
			item.dispatched <- struct{}{}
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

// A batch canceled before every update is dispatched stays unconfirmed, so the
// whole batch, including handled updates, comes back.
func TestBot_StartReplaysBatchCanceledBeforeDispatch(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resumeCtx, cancelResume := context.WithCancel(context.Background())
	defer cancelResume()

	batch := `{"ok":true,"result":[{"update_id":40},{"update_id":42}]}`
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
				return getUpdatesJSONResponse(batch), nil
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
	if got := waitForUpdateTest(t, offsets, "resumed poll"); got != "1" {
		t.Errorf("resumed poll offset = %q, want 1", got)
	}
	if got := waitForUpdateTest(t, offsets, "poll after the replayed batch"); got != "43" {
		t.Errorf("offset after the replayed batch = %q, want 43", got)
	}
	cancelResume()
	waitForUpdateTest(t, resumed, "resumed polling to stop")
	if want := []int64{40, 40, 42}; !reflect.DeepEqual(handled, want) {
		t.Errorf("handled updates = %v, want %v", handled, want)
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
