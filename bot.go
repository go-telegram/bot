package bot

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-telegram/bot/models"
)

const (
	defaultPollTimeout      = time.Minute
	defaultUpdatesChanCap   = 1024
	defaultCheckInitTimeout = time.Second * 5
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
	url         string
	token       string
	pollTimeout time.Duration
	skipGetMe   bool

	defaultHandlerFunc HandlerFunc

	errorsHandler ErrorsHandler
	debugHandler  DebugHandler

	middlewaresMx *sync.RWMutex
	middlewares   []Middleware

	handlersMx *sync.RWMutex
	handlers   map[string]handler

	client           HttpClient
	lastUpdateID     int64
	isDebug          bool
	checkInitTimeout time.Duration

	updates chan *models.Update
}

// New creates new Bot instance
func New(token string, options ...Option) (*Bot, error) {
	if strings.TrimSpace(token) == "" {
		return nil, fmt.Errorf("empty token")
	}

	b := &Bot{
		url:           "https://api.telegram.org",
		token:         token,
		pollTimeout:   defaultPollTimeout,
		middlewaresMx: &sync.RWMutex{},
		middlewares:   []Middleware{},
		handlersMx:    &sync.RWMutex{},
		handlers:      map[string]handler{},
		client: &http.Client{
			Timeout: defaultPollTimeout,
		},
		defaultHandlerFunc: defaultHandler,
		errorsHandler:      defaultErrorsHandler,
		debugHandler:       defaultDebugHandler,
		checkInitTimeout:   defaultCheckInitTimeout,

		updates: make(chan *models.Update, defaultUpdatesChanCap),
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

// StartWebhook starts the Bot with webhook mode
func (b *Bot) StartWebhook(ctx context.Context) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go b.waitUpdates(ctx, wg)

	wg.Wait()
}

// Start the bot
func (b *Bot) Start(ctx context.Context) {
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go b.waitUpdates(ctx, wg)
	go b.getUpdates(ctx, wg)

	wg.Wait()
}

func defaultErrorsHandler(err error) {
	log.Printf("[TGBOT] [ERROR] %v", err)
}

func defaultDebugHandler(format string, args ...interface{}) {
	log.Printf("[TGBOT] [DEBUG] "+format, args...)
}

func defaultHandler(_ context.Context, _ *Bot, update *models.Update) {
	log.Printf("[TGBOT] [UPDATE] %+v", update)
}

func (b *Bot) error(format string, args ...interface{}) {
	b.errorsHandler(fmt.Errorf(format, args...))
}

func True() *bool {
	b := true
	return &b
}

func False() *bool {
	b := false
	return &b
}
