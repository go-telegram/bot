package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strconv"
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

	req := httptest.NewRequest(http.MethodPost, "/", errReader(errors.New("read error")))
	w := httptest.NewRecorder()

	handler := bot.WebhookHandler()
	handler(w, req)

	if got := w.Result().StatusCode; got != http.StatusBadRequest {
		t.Errorf("status for unreadable request body = %d, want %d", got, http.StatusBadRequest)
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
	release := make(chan struct{})
	var releaseOnce sync.Once
	releaseHandler := func() { releaseOnce.Do(func() { close(release) }) }
	// A failed check must not leave the handler blocked.
	defer releaseHandler()
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
				// Pending requests must be released even while this handler is busy.
				<-release
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
	queued := make(chan int, 1)
	go func() { queued <- request(`{"update_id":2}`, context.Background()) }()
	b.updatesMx.Lock()
	queue := b.updatesRun.queue
	b.updatesMx.Unlock()
	deadline := time.Now().Add(5 * time.Second)
	for len(queue) == 0 {
		if time.Now().After(deadline) {
			t.Fatal("second update did not enter the queue")
		}
		time.Sleep(time.Millisecond)
	}

	// The worker is busy and its only queue slot is occupied.
	blocked := &observedDoneContext{Context: context.Background(), observed: make(chan struct{})}
	blockedStatus := make(chan int, 1)
	go func() { blockedStatus <- request(`{"update_id":3}`, blocked) }()
	waitForUpdateTest(t, blocked.observed, "request to wait for queue space")
	cancel()
	if got := waitForUpdateTest(t, queued, "queued request to return"); got != http.StatusServiceUnavailable {
		t.Errorf("status for queued update = %d, want 503", got)
	}
	if got := waitForUpdateTest(t, blockedStatus, "blocked request to return"); got != http.StatusServiceUnavailable {
		t.Errorf("status for blocked update = %d, want 503", got)
	}
	if got := request(`{"update_id":4}`, context.Background()); got != http.StatusServiceUnavailable {
		t.Errorf("status during shutdown = %d, want 503", got)
	}
	releaseHandler()
	waitForUpdateTest(t, stopped, "webhook to stop after cancellation")
	if got := waitForUpdateTest(t, delivered, "first handler to return"); got != 1 {
		t.Errorf("delivered update %d, want 1", got)
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
	if len(reported) != 0 {
		t.Errorf("unexpected error report: %v", <-reported)
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

// A worker that takes an update has dispatched it, even if its matcher then
// exits, so Telegram does not keep resending an update that kills a worker.
func TestWebhookHandler_AcknowledgesUpdateWhosePreparationExits(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	delivered := make(chan int64, 2)
	b, err := New("xxx", WithSkipGetMe(), WithWorkers(2), WithUpdatesChannelCap(0),
		WithDefaultHandler(func(_ context.Context, _ *Bot, update *models.Update) {
			delivered <- update.ID
		}))
	if err != nil {
		t.Fatal(err)
	}
	b.RegisterHandlerMatchFunc(func(update *models.Update) bool {
		if update.ID != 2 {
			runtime.Goexit()
		}
		return false
	}, b.defaultHandlerFunc)
	stopped := startWebhookForTest(t, b, ctx)
	request := func(body string) int {
		w := httptest.NewRecorder()
		b.WebhookHandler()(w, httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body)))
		return w.Code
	}
	if got := request(`{"update_id":1}`); got != http.StatusOK {
		t.Errorf("status for interrupted preparation = %d, want 200", got)
	}
	if got := request(`{"update_id":2}`); got != http.StatusOK {
		t.Errorf("status with surviving worker = %d, want 200", got)
	}
	if got := waitForUpdateTest(t, delivered, "surviving worker to handle update"); got != 2 {
		t.Errorf("delivered update = %d, want 2", got)
	}
	if got := request(`{"update_id":3}`); got != http.StatusOK {
		t.Errorf("status for last worker's interrupted preparation = %d, want 200", got)
	}
	waitForUpdateTest(t, stopped, "webhook to stop after its last worker exits")
}

func TestWebhookHandler_CompletesClaimedMiddlewareAfterCancellation(t *testing.T) {
	for _, stopRun := range []bool{false, true} {
		ctx, stop := context.WithCancel(context.Background())
		requestCtx, cancelRequest := context.WithCancel(context.Background())
		preparing := make(chan struct{})
		releasePreparation := make(chan struct{})
		resource := make(chan struct{}, 1)
		b, err := New("xxx", WithSkipGetMe(), WithUpdatesChannelCap(0),
			WithMiddlewares(func(next HandlerFunc) HandlerFunc {
				resource <- struct{}{}
				close(preparing)
				<-releasePreparation
				return func(ctx context.Context, b *Bot, update *models.Update) {
					defer func() { <-resource }()
					next(ctx, b, update)
				}
			}),
			WithDefaultHandler(func(context.Context, *Bot, *models.Update) {}))
		if err != nil {
			stop()
			cancelRequest()
			t.Fatal(err)
		}
		stopped := startWebhookForTest(t, b, ctx)
		status := make(chan int, 1)
		go func() {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"update_id":1}`)).WithContext(requestCtx)
			b.WebhookHandler()(w, req)
			status <- w.Code
		}()
		waitForUpdateTest(t, preparing, "middleware to acquire the resource")
		if stopRun {
			stop()
		} else {
			cancelRequest()
		}
		close(releasePreparation)
		if got := waitForUpdateTest(t, status, "claimed request to complete"); got != http.StatusOK {
			t.Errorf("stopRun=%t: status = %d, want 200", stopRun, got)
		}
		stop()
		cancelRequest()
		waitForUpdateTest(t, stopped, "claimed handler to finish")
		if len(resource) != 0 {
			t.Errorf("stopRun=%t: middleware resource was not released", stopRun)
		}
	}
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

func startWebhookForTest(t *testing.T, b *Bot, ctx context.Context) <-chan struct{} {
	t.Helper()
	stopped := make(chan struct{})
	go func() {
		defer close(stopped)
		b.StartWebhook(ctx)
	}()
	deadline := time.Now().Add(5 * time.Second)
	for {
		b.updatesMx.Lock()
		started := b.updatesRun != nil && b.updatesRun.webhook
		b.updatesMx.Unlock()
		if started {
			return stopped
		}
		if time.Now().After(deadline) {
			t.Fatal("webhook did not start")
		}
		time.Sleep(time.Millisecond)
	}
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
	var releaseOnce sync.Once
	releaseHandler := func() { releaseOnce.Do(func() { close(release) }) }
	// A failed check must not leave the handler blocked.
	defer releaseHandler()
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

	releaseHandler()
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

// Requests canceled while their updates wait in the queue get 503, and the
// worker skips those updates instead of handling them. A worker that could
// complete the handoff without its request would handle some of them, so
// many requests make that failure reliable.
func TestWebhookHandler_SkipsUpdatesOfCanceledRequests(t *testing.T) {
	const canceledRequests = 20
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	firstStarted := make(chan struct{})
	release := make(chan struct{})
	var releaseOnce sync.Once
	releaseHandler := func() { releaseOnce.Do(func() { close(release) }) }
	// A failed check must not leave the handler blocked.
	defer releaseHandler()
	delivered := make(chan int64, canceledRequests+2)
	b, err := New("xxx", WithSkipGetMe(), WithWorkers(1), WithUpdatesChannelCap(canceledRequests), WithNotAsyncHandlers(),
		WithDefaultHandler(func(_ context.Context, _ *Bot, update *models.Update) {
			if update.ID == 1 {
				close(firstStarted)
				<-release
			}
			delivered <- update.ID
		}))
	if err != nil {
		t.Fatal(err)
	}
	request := func(id int, ctx context.Context) int {
		w := httptest.NewRecorder()
		body := strings.NewReader(`{"update_id":` + strconv.Itoa(id) + `}`)
		b.WebhookHandler()(w, httptest.NewRequest(http.MethodPost, "/", body).WithContext(ctx))
		return w.Code
	}
	stopped := startWebhookForTest(t, b, ctx)
	if got := request(1, context.Background()); got != http.StatusOK {
		t.Fatalf("status for first update = %d, want 200", got)
	}
	waitForUpdateTest(t, firstStarted, "first handler to start")

	requestCtx, cancelRequests := context.WithCancel(context.Background())
	defer cancelRequests()
	statuses := make(chan int, canceledRequests)
	for id := 2; id < canceledRequests+2; id++ {
		go func(id int) { statuses <- request(id, requestCtx) }(id)
	}
	b.updatesMx.Lock()
	queue := b.updatesRun.queue
	b.updatesMx.Unlock()
	deadline := time.Now().Add(5 * time.Second)
	for len(queue) < canceledRequests {
		if time.Now().After(deadline) {
			t.Fatal("canceled updates did not enter the queue")
		}
		time.Sleep(time.Millisecond)
	}
	cancelRequests()
	for i := 0; i < canceledRequests; i++ {
		if got := waitForUpdateTest(t, statuses, "canceled request"); got != http.StatusServiceUnavailable {
			t.Errorf("status for canceled request = %d, want 503", got)
		}
	}

	releaseHandler()
	const lastID = 1000
	if got := request(lastID, context.Background()); got != http.StatusOK {
		t.Fatalf("status after the canceled requests = %d, want 200", got)
	}
	for _, want := range []int64{1, lastID} {
		if got := waitForUpdateTest(t, delivered, "handler"); got != want {
			t.Errorf("delivered update %d, want %d", got, want)
		}
	}
	cancel()
	waitForUpdateTest(t, stopped, "webhook to stop")
}

// A webhook request that arrives while Start is polling is rejected.
func TestWebhookHandler_RejectsUpdateWhilePolling(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	polling := make(chan struct{})
	var once sync.Once
	handled := make(chan struct{}, 1)
	b, err := New("xxx", WithSkipGetMe(),
		WithDefaultHandler(func(context.Context, *Bot, *models.Update) { handled <- struct{}{} }),
		WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
			once.Do(func() { close(polling) })
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
	waitForUpdateTest(t, polling, "polling to start")
	w := httptest.NewRecorder()
	b.WebhookHandler()(w, httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"update_id":1}`)))
	if w.Code != http.StatusServiceUnavailable {
		t.Errorf("webhook status while polling = %d, want 503", w.Code)
	}
	cancel()
	waitForUpdateTest(t, stopped, "polling to stop")
	if len(handled) != 0 {
		t.Error("handled a webhook update while polling")
	}
}

func TestStartWebhook_ShutdownGraceAllowsInFlightHandler(t *testing.T) {
	type ctxKey struct{}
	parent := context.WithValue(context.Background(), ctxKey{}, "wh-trace")
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	started := make(chan struct{})
	allowFinish := make(chan struct{})
	var releaseOnce sync.Once
	releaseHandler := func() { releaseOnce.Do(func() { close(allowFinish) }) }
	defer releaseHandler()
	sendErrCh := make(chan error, 1)
	ctxValCh := make(chan any, 1)

	b, err := New("xxx", WithSkipGetMe(), WithShutdownGrace(500*time.Millisecond),
		WithErrorsHandler(func(error) {}),
		WithHTTPClient(time.Second, getUpdatesClientFunc(func(req *http.Request) (*http.Response, error) {
			if err := req.Context().Err(); err != nil {
				return nil, err
			}
			return getUpdatesJSONResponse(`{"ok":true,"result":{"message_id":20}}`), nil
		})),
		WithDefaultHandler(func(hCtx context.Context, bot *Bot, _ *models.Update) {
			ctxValCh <- hCtx.Value(ctxKey{})
			close(started)
			<-allowFinish
			_, sendErr := bot.SendMessage(hCtx, &SendMessageParams{ChatID: int64(1), Text: "wh-done"})
			sendErrCh <- sendErr
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

	waitForUpdateTest(t, started, "webhook handler to start")
	cancel()

	wAfter := httptest.NewRecorder()
	b.WebhookHandler()(wAfter, httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"update_id":2}`)))
	if wAfter.Code != http.StatusServiceUnavailable {
		t.Fatalf("webhook status after cancel = %d, want 503", wAfter.Code)
	}

	select {
	case <-stopped:
		t.Fatal("StartWebhook returned before in-flight handler completed within grace")
	case <-time.After(30 * time.Millisecond):
	}

	releaseHandler()
	waitForUpdateTest(t, stopped, "StartWebhook to return after handler")

	if gotErr := waitForUpdateTest(t, sendErrCh, "webhook sendMessage result"); gotErr != nil {
		t.Fatalf("SendMessage during webhook shutdown grace failed: %v", gotErr)
	}
	if gotVal := waitForUpdateTest(t, ctxValCh, "webhook handler context value"); gotVal != "wh-trace" {
		t.Fatalf("webhook handler context value = %v, want wh-trace", gotVal)
	}
}
