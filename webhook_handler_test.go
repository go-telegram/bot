package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/go-telegram/bot/models"
)

type mockDebugHandler struct {
	messages []string
}

func (h *mockDebugHandler) Handle(format string, args ...any) {
	h.messages = append(h.messages, format)
}

type mockErrorsHandler struct {
	errors []error
}

func (h *mockErrorsHandler) Handle(err error) {
	h.errors = append(h.errors, err)
}

type errReaderStruct struct {
	err error
}

func (e *errReaderStruct) Read(p []byte) (int, error) {
	return 0, e.err
}

func (e *errReaderStruct) Close() error {
	return nil
}

func errReader(err error) io.ReadCloser {
	return &errReaderStruct{err: err}
}

func TestWebhookHandler_Success(t *testing.T) {
	debugHandler := &mockDebugHandler{}
	errorsHandler := &mockErrorsHandler{}
	delivered := make(chan *models.Update, 1)

	bot := &Bot{
		workers: 1,
		isDebug: true,
		defaultHandlerFunc: func(_ context.Context, _ *Bot, update *models.Update) {
			delivered <- update
		},
		debugHandler: func(format string, args ...any) {
			debugHandler.Handle(format, args...)
		},
		errorsHandler: func(err error) {
			errorsHandler.Handle(err)
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stopped := startWebhookForTest(t, bot, ctx)

	update := &models.Update{
		ID: 12345,
	}
	updateBody, _ := json.Marshal(update)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(updateBody))
	w := httptest.NewRecorder()

	handler := bot.WebhookHandler()
	handler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("webhook status %d, want 200", w.Code)
	}
	upd := waitForUpdateTest(t, delivered, "update to be delivered")
	if upd.ID != update.ID {
		t.Fatalf("Expected update ID %d, got %d", update.ID, upd.ID)
	}
	cancel()
	waitForUpdateTest(t, stopped, "webhook to stop")
}

func TestWebhookHandler_ReadBodyError(t *testing.T) {
	errorsHandler := &mockErrorsHandler{}
	delivered := make(chan *models.Update, 1)

	bot := &Bot{
		workers: 1,
		defaultHandlerFunc: func(_ context.Context, _ *Bot, update *models.Update) {
			delivered <- update
		},
		errorsHandler: func(err error) {
			errorsHandler.Handle(err)
		},
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stopped := startWebhookForTest(t, bot, ctx)

	req := httptest.NewRequest(http.MethodPost, "/", errReader(errors.New("read error")))
	w := httptest.NewRecorder()

	handler := bot.WebhookHandler()
	handler(w, req)
	cancel()
	waitForUpdateTest(t, stopped, "webhook to stop")

	if w.Code != http.StatusBadRequest {
		t.Errorf("webhook status %d, want %d", w.Code, http.StatusBadRequest)
	}
	if len(delivered) != 0 {
		t.Error("webhook delivered an update after the body read failed")
	}

	if len(errorsHandler.errors) == 0 {
		t.Fatal("Expected an error, but none occurred")
	}

	if capturedError := errorsHandler.errors[0]; capturedError == nil || !containsString(capturedError.Error(), "read error") {
		t.Fatalf("Expected read body error, got %v", capturedError)
	}
}

func TestWebhookHandler_DecodeError(t *testing.T) {
	debugHandler := &mockDebugHandler{}
	errorsHandler := &mockErrorsHandler{}

	bot := &Bot{
		debugHandler: func(format string, args ...any) {
			debugHandler.Handle(format, args...)
		},
		errorsHandler: func(err error) {
			errorsHandler.Handle(err)
		},
	}

	invalidJSON := []byte("{invalid json}")
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(invalidJSON))
	w := httptest.NewRecorder()

	handler := bot.WebhookHandler()
	handler(w, req)

	if len(errorsHandler.errors) == 0 {
		t.Fatal("Expected an error, but none occurred")
	}

	if capturedError := errorsHandler.errors[0]; capturedError == nil || !containsString(capturedError.Error(), "error decode request body") {
		t.Fatalf("Expected decode error, got %v", capturedError)
	}
}

func TestWebhookHandler_ContextDone(t *testing.T) {
	delivered := make(chan struct{}, 1)
	bot := &Bot{
		workers:        1,
		updatesChanCap: 1,
		defaultHandlerFunc: func(context.Context, *Bot, *models.Update) {
			delivered <- struct{}{}
		},
	}
	runCtx, stop := context.WithCancel(context.Background())
	defer stop()
	stopped := startWebhookForTest(t, bot, runCtx)

	update := &models.Update{
		ID: 12345,
	}
	updateBody, _ := json.Marshal(update)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(updateBody)).WithContext(ctx)
	w := httptest.NewRecorder()

	handler := bot.WebhookHandler()
	handler(w, req)

	stop()
	waitForUpdateTest(t, stopped, "webhook to stop")
	if w.Code != http.StatusServiceUnavailable {
		t.Errorf("webhook status %d, want 503", w.Code)
	}
	if len(delivered) != 0 {
		t.Fatal("processed an update from a canceled request")
	}
}

func TestWebhookHandler_StopsDispatchingOnShutdown(t *testing.T) {
	type contextKey struct{}
	parent := context.WithValue(context.Background(), contextKey{}, "value")
	ctx, cancel := context.WithCancel(parent)
	defer cancel()
	firstStarted := make(chan struct{})
	delivered := make(chan int64, 6)
	reported := make(chan error, 6)
	b, err := New("xxx", WithSkipGetMe(), WithUpdatesChannelCap(1), WithNotAsyncHandlers(),
		WithErrorsHandler(func(err error) { reported <- err }),
		WithDefaultHandler(func(ctx context.Context, _ *Bot, update *models.Update) {
			if ctx.Value(contextKey{}) != "value" {
				t.Errorf("handler context lost its values: %v", ctx)
			}
			if update.ID == 1 {
				close(firstStarted)
				// A handler waiting for cancellation must not block shutdown.
				<-ctx.Done()
			}
			delivered <- update.ID
		}))
	if err != nil {
		t.Fatal(err)
	}
	handler := b.WebhookHandler()
	request := func(body string, ctx context.Context) int {
		w := httptest.NewRecorder()
		handler(w, httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body)).WithContext(ctx))
		return w.Code
	}

	if got := request(`{"update_id":99}`, context.Background()); got != http.StatusServiceUnavailable {
		t.Fatalf("status before starting webhook = %d, want 503", got)
	}
	if got := waitForUpdateTest(t, reported, "report before start"); !strings.Contains(got.Error(), "StartWebhook is not running") {
		t.Errorf("reported %v, want the not running report", got)
	}
	stopped := startWebhookForTest(t, b, ctx)
	if got := request(`{"update_id":1}`, context.Background()); got != http.StatusOK {
		t.Fatalf("status for first update = %d, want 200", got)
	}
	waitForUpdateTest(t, firstStarted, "first webhook handler to start")
	if got := request(`{"update_id":2}`, context.Background()); got != http.StatusOK {
		t.Fatalf("status for queued update = %d, want 200", got)
	}

	// The worker is busy and its only queue slot is occupied.
	blockedCtx, cancelBlocked := context.WithCancel(context.Background())
	defer cancelBlocked()
	blocked := &observedDoneContext{Context: blockedCtx, observed: make(chan struct{})}
	blockedStatus := make(chan int, 1)
	go func() {
		blockedStatus <- request(`{"update_id":3}`, blocked)
	}()
	waitForUpdateTest(t, blocked.observed, "request to wait for queue space")
	cancel()
	if got := waitForUpdateTest(t, blockedStatus, "blocked request to return"); got != http.StatusServiceUnavailable {
		t.Errorf("status for blocked update = %d, want 503", got)
	}
	if got := request(`{"update_id":4}`, context.Background()); got != http.StatusServiceUnavailable {
		t.Errorf("status during shutdown = %d, want 503", got)
	}
	waitForUpdateTest(t, stopped, "webhook to stop after cancellation")
	if got := waitForUpdateTest(t, delivered, "first handler to return"); got != 1 {
		t.Errorf("delivered update %d, want 1", got)
	}
	if got := waitForUpdateTest(t, reported, "dropped update report"); !strings.Contains(got.Error(), "dropped 1") {
		t.Errorf("reported %v, want the dropped queued update", got)
	}

	if got := request(`{"update_id":5}`, context.Background()); got != http.StatusServiceUnavailable {
		t.Errorf("status after shutdown = %d, want 503", got)
	}
	if got := waitForUpdateTest(t, reported, "report after shutdown"); !strings.Contains(got.Error(), "StartWebhook is not running") {
		t.Errorf("reported %v, want the not running report", got)
	}
	restartCtx, cancelRestart := context.WithCancel(parent)
	defer cancelRestart()
	restarted := startWebhookForTest(t, b, restartCtx)
	if got := request(`{"update_id":6}`, context.Background()); got != http.StatusOK {
		t.Fatalf("status after restart = %d, want 200", got)
	}
	if got := waitForUpdateTest(t, delivered, "update after restart"); got != 6 {
		t.Errorf("delivered update %d, want 6", got)
	}
	cancelRestart()
	waitForUpdateTest(t, restarted, "restarted webhook to stop")
	if len(delivered) != 0 {
		t.Error("delivered a dropped or rejected webhook update")
	}
	// A request that arrives after shutdown finished reports that no webhook is
	// running, which the request during shutdown may or may not hit.
	for len(reported) != 0 {
		if got := <-reported; !strings.Contains(got.Error(), "StartWebhook is not running") {
			t.Errorf("unexpected error report: %v", got)
		}
	}
}

func TestStartWebhook_ReturnsWhenAsyncHandlerWaitsForCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	started := make(chan struct{})
	b, err := New("xxx", WithSkipGetMe(),
		WithDefaultHandler(func(ctx context.Context, _ *Bot, _ *models.Update) {
			close(started)
			<-ctx.Done()
		}))
	if err != nil {
		t.Fatal(err)
	}
	stopped := startWebhookForTest(t, b, ctx)
	w := httptest.NewRecorder()
	b.WebhookHandler()(w, httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"update_id":1}`)))
	if w.Code != http.StatusOK {
		t.Fatalf("webhook status = %d, want 200", w.Code)
	}
	waitForUpdateTest(t, started, "handler to start")
	cancel()
	waitForUpdateTest(t, stopped, "webhook to stop after cancellation")
}

func TestWebhookHandler_NoWorkers(t *testing.T) {
	b, err := New("xxx", WithSkipGetMe(), WithWorkers(0), WithErrorsHandler(func(error) {}))
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		b.StartWebhook(ctx)
	}()
	waitForUpdateTest(t, stopped, "webhook with no workers to stop")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"update_id":1}`))
	b.WebhookHandler().ServeHTTP(w, req)
	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("webhook status without workers = %d, want 503", w.Code)
	}
}

type observedDoneContext struct {
	context.Context
	observed chan struct{}
	once     sync.Once
}

func (ctx *observedDoneContext) Done() <-chan struct{} {
	ctx.once.Do(func() { close(ctx.observed) })
	return ctx.Context.Done()
}

// startWebhookForTest relies on StartWebhook calling ctx.Done only after it
// registered the run, which holds unless it waits for a previous run. Start
// such a second run in a plain goroutine instead.
func startWebhookForTest(t *testing.T, b *Bot, ctx context.Context) <-chan struct{} {
	t.Helper()
	observed := &observedDoneContext{Context: ctx, observed: make(chan struct{})}
	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		b.StartWebhook(observed)
	}()
	waitForUpdateTest(t, observed.observed, "webhook to start")
	return stopped
}

func containsString(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}

func TestStartWebhook_RestartWaitsForPreviousRun(t *testing.T) {
	ctx1, cancel1 := context.WithCancel(context.Background())
	defer cancel1()
	ctx2, cancel2 := context.WithCancel(context.Background())
	defer cancel2()
	started := make(chan struct{})
	release := make(chan struct{})
	delivered := make(chan int64, 4)
	reported := make(chan error, 8)
	b, err := New("xxx", WithSkipGetMe(),
		WithErrorsHandler(func(err error) {
			// Many 503s can pile up in the handover window; never block a request.
			select {
			case reported <- err:
			default:
			}
		}),
		WithDefaultHandler(func(_ context.Context, _ *Bot, update *models.Update) {
			if update.ID == 1 {
				close(started)
				// The handler ignores cancellation on purpose.
				<-release
				return
			}
			delivered <- update.ID
		}))
	if err != nil {
		t.Fatal(err)
	}
	request := func(body string) int {
		w := httptest.NewRecorder()
		b.WebhookHandler()(w, httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body)))
		return w.Code
	}

	firstStopped := startWebhookForTest(t, b, ctx1)
	if got := request(`{"update_id":1}`); got != http.StatusOK {
		t.Fatalf("status for first update = %d, want 200", got)
	}
	waitForUpdateTest(t, started, "first handler to start")
	cancel1()
	secondStopped := make(chan struct{})
	go func() {
		defer close(secondStopped)
		b.StartWebhook(ctx2)
	}()
	select {
	case <-secondStopped:
		t.Fatal("second StartWebhook returned while the previous run was still shutting down")
	case err := <-reported:
		t.Fatalf("unexpected error report: %v", err)
	case <-time.After(100 * time.Millisecond):
	}

	close(release)
	waitForUpdateTest(t, firstStopped, "first webhook to stop")
	deadline := time.Now().Add(5 * time.Second)
	for request(`{"update_id":2}`) != http.StatusOK {
		if time.Now().After(deadline) {
			t.Fatal("second run never accepted an update")
		}
		time.Sleep(10 * time.Millisecond)
	}
	if got := waitForUpdateTest(t, delivered, "update from the second run"); got != 2 {
		t.Errorf("delivered update %d, want 2", got)
	}
	cancel2()
	waitForUpdateTest(t, secondStopped, "second webhook to stop")
}
