package bot

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-telegram/bot/models"
)

const (
	defaultPollTimeout = time.Minute
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type ErrorsHandler func(err error)
type Middleware func(next HandlerFunc) HandlerFunc
type HandlerFunc func(ctx context.Context, bot *Bot, update *models.Update)

// Bot represents Telegram Bot main object
type Bot struct {
	url         string
	token       string
	pollTimeout time.Duration

	defaultHandlerFunc HandlerFunc

	errorsHandler ErrorsHandler

	middlewaresMx *sync.RWMutex
	middlewares   []Middleware

	handlersMx *sync.RWMutex
	handlers   map[string]handler

	client       HttpClient
	lastUpdateID int64
	isDebug      bool

	updates chan *models.Update
}

// New creates new Bot instance
func New(ctx context.Context, token string, options ...Option) *Bot {
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

		updates: make(chan *models.Update),
	}

	for _, o := range options {
		o(b)
	}

	go b.waitUpdates(ctx)

	return b
}

func defaultErrorsHandler(err error) {
	log.Printf("[TGBOT] [ERROR] %v", err)
}

func defaultHandler(_ context.Context, _ *Bot, upd *models.Update) {
	log.Printf("[TGBOT] [UPDATE] %+v", upd)
}

func (b *Bot) error(format string, args ...interface{}) {
	b.errorsHandler(fmt.Errorf(format, args...))
}

func (b *Bot) debug(format string, args ...interface{}) {
	if !b.isDebug {
		return
	}

	log.Printf("[TGBOT] [DEBUG] "+format, args...)
}
