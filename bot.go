package bot

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-telegram/bot/models"
)

const (
	defaultPollTimeout      = time.Minute
	defaultUpdatesChanCap   = 1024
	defaultCheckInitTimeout = time.Second * 5
	defaultWorkers          = 1
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type ErrorsHandler func(err error)
type DebugHandler func(format string, args ...any)
type Middleware func(next HandlerFunc) HandlerFunc
type HandlerFunc func(ctx context.Context, bot *Bot, update *models.Update)
type MatchFunc func(update *models.Update) bool

// Bot represents Telegram Bot main object
type Bot struct {
	lastUpdateID int64

	url                string
	token              string
	pollTimeout        time.Duration
	skipGetMe          bool
	webhookSecretToken string
	testEnvironment    bool
	workers            int
	notAsyncHandlers   bool

	defaultHandlerFunc HandlerFunc

	errorsHandler ErrorsHandler
	debugHandler  DebugHandler

	middlewares []Middleware

	handlersMx sync.RWMutex
	handlers   []handler

	client           HttpClient
	isDebug          bool
	checkInitTimeout time.Duration

	allowedUpdates AllowedUpdates

	updatesChanCap int
	updatesMx      sync.Mutex
	updatesRun     *updatesRun
}

// New creates new Bot instance
func New(token string, options ...Option) (*Bot, error) {
	if strings.TrimSpace(token) == "" {
		return nil, fmt.Errorf("empty token")
	}

	b := &Bot{
		url:         "https://api.telegram.org",
		token:       token,
		pollTimeout: defaultPollTimeout,
		client: &http.Client{
			Timeout: defaultPollTimeout,
		},
		defaultHandlerFunc: defaultHandler,
		errorsHandler:      defaultErrorsHandler,
		debugHandler:       defaultDebugHandler,
		checkInitTimeout:   defaultCheckInitTimeout,
		workers:            defaultWorkers,

		updatesChanCap: defaultUpdatesChanCap,
	}

	for _, o := range options {
		o(b)
	}

	ctx, cancel := context.WithTimeout(context.Background(), b.checkInitTimeout)
	defer cancel()

	if !b.skipGetMe {
		_, err := b.GetMe(ctx)
		if err != nil {
			return nil, fmt.Errorf("error call getMe, %w", err)
		}
	}

	return b, nil
}

// ID returns the bot user id from the token prefix ("<id>:<secret>").
// It returns 0 if the prefix cannot be parsed as an int64.
// ID does not validate or authenticate the token.
func (b *Bot) ID() int64 {
	i, err := strconv.ParseInt(strings.Split(b.token, ":")[0], 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// SetToken sets the bot token
func (b *Bot) SetToken(token string) {
	b.token = token
}

// Token returns the bot token
func (b *Bot) Token() string {
	return b.token
}

// StartWebhook processes updates from WebhookHandler until ctx is canceled.
// Handlers receive ctx. On cancellation, WebhookHandler rejects new updates
// with 503 so Telegram retries them, workers stop dispatching, and
// StartWebhook returns once running handlers return. Accepted updates that
// were still queued are dropped and reported to the errors handler. Handlers
// should return when ctx is canceled: StartWebhook does not return until they
// do. If another Start or StartWebhook is running on the same Bot, it waits for
// that run to end first. Calling it from a handler of the running Bot therefore
// blocks until the ctx passed to that call is canceled.
func (b *Bot) StartWebhook(ctx context.Context) {
	run := b.beginUpdates(ctx, true)
	if run == nil {
		return
	}
	defer b.endUpdates()

	workers := b.startUpdateWorkers(ctx, run)
	<-ctx.Done()

	b.updatesMx.Lock()
	run.webhookCtx = nil
	b.updatesMx.Unlock()
	close(run.queue)
	workers.Wait()

	// A worker that exited early, such as through runtime.Goexit, can leave items behind.
	for range run.queue {
		atomic.AddInt64(&run.dropped, 1)
	}

	if dropped := atomic.LoadInt64(&run.dropped); dropped > 0 {
		b.error("webhook stopped before dispatching accepted updates, dropped %d", dropped)
	}
}

// Start receives and processes updates by polling until ctx is canceled.
// A batch is confirmed once all its updates have been dispatched to handlers,
// so polling continues while asynchronous handlers run. With
// WithNotAsyncHandlers, a batch is dispatched only as fast as its handlers
// run. On cancellation, Start waits for dispatched handlers before returning.
// If cancellation comes before every update of a batch is dispatched, the
// whole batch, including updates already handled, may be delivered again when
// polling resumes. Handlers should return when ctx is canceled: Start does not
// return until they do. If another Start or StartWebhook is running on the same
// Bot, it waits for that run to end first. Calling it from a handler of the
// running Bot therefore blocks until the ctx passed to that call is canceled.
func (b *Bot) Start(ctx context.Context) {
	run := b.beginUpdates(ctx, false)
	if run == nil {
		return
	}
	defer b.endUpdates()

	workers := b.startUpdateWorkers(ctx, run)
	b.getUpdates(ctx, run.queue)
	close(run.queue)
	workers.Wait()
}

func defaultErrorsHandler(err error) {
	log.Printf("[TGBOT] [ERROR] %v", err)
}

func defaultDebugHandler(format string, args ...any) {
	log.Printf("[TGBOT] [DEBUG] "+format, args...)
}

func defaultHandler(_ context.Context, _ *Bot, update *models.Update) {
	log.Printf("[TGBOT] [UPDATE] %+v", update)
}

func (b *Bot) error(format string, args ...any) {
	b.errorsHandler(fmt.Errorf(format, args...))
}

// True and False returns the pointer to bool
func True() *bool {
	b := true
	return &b
}

// False and True returns the pointer to bool
func False() *bool {
	b := false
	return &b
}

// FileDownloadLink returns the file download link
func (b *Bot) FileDownloadLink(f *models.File) string {
	return fmt.Sprintf("%s/file/bot%s/%s", b.url, b.token, f.FilePath)
}
