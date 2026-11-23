package bot

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
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

	// usernameMx guards username, the username from the last successful GetMe
	// call for the bot the token belongs to. SetToken holds it while changing the
	// token, so checking the bot ID and clearing the username happen together;
	// other code still reads the token without it.
	usernameMx sync.RWMutex
	username   string

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
	if b.updatesChanCap < 0 {
		return nil, fmt.Errorf("updates channel capacity must be non-negative")
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

// SetToken sets the bot token. If the token belongs to a different bot, by its
// ID prefix, the username recorded by GetMe is cleared until GetMe succeeds
// again. A new token for the same bot, as after revoking the old one, keeps it.
func (b *Bot) SetToken(token string) {
	b.usernameMx.Lock()
	defer b.usernameMx.Unlock()
	previousID := b.ID()
	b.token = token
	if b.ID() != previousID {
		b.username = ""
	}
}

// setUsername records the username GetMe returned, unless the token now
// belongs to another bot, as when it was switched during the request. The
// token starts with the bot's ID, which GetMe returns as id.
func (b *Bot) setUsername(id int64, username string) {
	b.usernameMx.Lock()
	defer b.usernameMx.Unlock()
	if id == b.ID() {
		b.username = username
	}
}

func (b *Bot) botUsername() string {
	b.usernameMx.RLock()
	defer b.usernameMx.RUnlock()
	return b.username
}

// Token returns the bot token
func (b *Bot) Token() string {
	return b.token
}

// StartWebhook processes updates from WebhookHandler until ctx is canceled.
// WebhookHandler answers a valid update with HTTP 200 once a worker takes it,
// and with 503 if the run stops first. Handlers receive a context derived from
// ctx and canceled when the run stops. StartWebhook waits for running handlers
// to return. If every worker exits early, such as through runtime.Goexit in a
// matcher or in a handler run with WithNotAsyncHandlers, the run stops. If
// another Start or StartWebhook is running on the same Bot, it waits for that
// run to end first. Calling it from a handler of the running Bot therefore
// blocks until the ctx passed to that call is canceled.
func (b *Bot) StartWebhook(ctx context.Context) {
	b.start(ctx, true)
}

// Start receives and processes updates by polling until ctx is canceled. A
// batch is confirmed once all its updates have been dispatched to handlers, so
// polling continues while asynchronous handlers run. With WithNotAsyncHandlers,
// a batch is dispatched only as fast as its handlers run. On cancellation,
// Start waits for dispatched handlers before returning. When polling resumes
// after cancellation, it starts from the first update of the batch that was not
// dispatched, so updates handled after it may be delivered again. Telegram
// learns the new offset only from the next poll, so after a process restart it
// may deliver the last batch again: delivery is at least once. Handlers receive
// a context derived from ctx and canceled when the run stops. They should
// return when it is canceled. If every worker exits early, such as through
// runtime.Goexit in a matcher or in a handler run with WithNotAsyncHandlers,
// the run stops. If another Start or StartWebhook is running on the same Bot,
// it waits for that run to end first. Calling it from a handler of the running
// Bot therefore blocks until the ctx passed to that call is canceled.
func (b *Bot) Start(ctx context.Context) {
	b.start(ctx, false)
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
