package bot

import (
	"context"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-telegram/bot/models"
)

type getUpdatesClientFunc func(*http.Request) (*http.Response, error)

func (f getUpdatesClientFunc) Do(req *http.Request) (*http.Response, error) {
	// drain the multipart body so the request pipe writer can finish
	_, _ = io.Copy(io.Discard, req.Body)
	_ = req.Body.Close()
	return f(req)
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
		updates:       make(chan *models.Update, 10),
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

	wg := sync.WaitGroup{}
	wg.Add(1)
	b.getUpdates(ctx, &wg)
	wg.Wait()

	if got := atomic.LoadInt64(&b.lastUpdateID); got != 102 {
		t.Fatalf("offset must move past the bad update, lastUpdateID=%d", got)
	}
	if len(b.updates) != 2 {
		t.Fatalf("expected 2 delivered updates, got %d", len(b.updates))
	}
	if first, third := <-b.updates, <-b.updates; first.ID != 100 || third.ID != 102 {
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
		updates:       make(chan *models.Update, 10),
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

	wg := sync.WaitGroup{}
	wg.Add(1)
	b.getUpdates(ctx, &wg)
	wg.Wait()

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
		updates:       make(chan *models.Update, 100),
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

	wg := sync.WaitGroup{}
	wg.Add(1)
	b.getUpdates(ctx, &wg)
	wg.Wait()

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
		updates:       make(chan *models.Update, 10),
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

	wg := sync.WaitGroup{}
	wg.Add(1)
	b.getUpdates(ctx, &wg)
	wg.Wait()

	if got := atomic.LoadInt64(&b.lastUpdateID); got != 500 {
		t.Fatalf("offset must stay at the last real update, lastUpdateID=%d", got)
	}
	if len(b.updates) != 1 {
		t.Fatalf("expected only the real update to be delivered, got %d", len(b.updates))
	}
}
