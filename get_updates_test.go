package bot

import (
	"context"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/go-telegram/bot/models"
)

type clientFunc func(*http.Request) (*http.Response, error)

func (f clientFunc) Do(req *http.Request) (*http.Response, error) {
	// drain the multipart body so the request pipe writer can finish
	_, _ = io.Copy(io.Discard, req.Body)
	_ = req.Body.Close()
	return f(req)
}

func jsonResponse(body string) *http.Response {
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
		errorsHandler: func(err error) { errs = append(errs, err.Error()) },
		debugHandler:  func(string, ...any) {},
		client: clientFunc(func(*http.Request) (*http.Response, error) {
			if atomic.AddInt32(&calls, 1) > 1 {
				cancel()
				return nil, ctx.Err()
			}
			return jsonResponse(batch), nil
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
	if len(errs) != 1 || !strings.Contains(errs[0], "101") {
		t.Fatalf("undecodable update must be reported once with its id, got %v", errs)
	}
}
