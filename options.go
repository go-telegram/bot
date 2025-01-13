package bot

import (
	"time"

	"github.com/go-telegram/bot/models"
)

// Option is a function that configures a bot.
type Option func(b *Bot)

// WithCheckInitTimeout allows to redefine CheckInitTimeout
func WithCheckInitTimeout(timeout time.Duration) Option {
	return func(b *Bot) {
		b.checkInitTimeout = timeout
	}
}

// WithMiddlewares allows to set middlewares for each incoming request
func WithMiddlewares(middlewares ...Middleware) Option {
	return func(b *Bot) {
		b.middlewares = append(b.middlewares, middlewares...)
	}
}

// WithMessageTextHandler allows to set handler for incoming text messages
// Also you can use *bot.RegisterHandler function after bot creation
func WithMessageTextHandler(pattern string, matchType MatchType, handler HandlerFunc) Option {
	return func(b *Bot) {
		b.RegisterHandler(HandlerTypeMessageText, pattern, matchType, handler)
	}
}

// WithCallbackQueryDataHandler allows to set handler for incoming callback query
// Also you can use *bot.RegisterHandler function after bot creation
func WithCallbackQueryDataHandler(pattern string, matchType MatchType, handler HandlerFunc) Option {
	return func(b *Bot) {
		b.RegisterHandler(HandlerTypeCallbackQueryData, pattern, matchType, handler)
	}
}

// WithPhotoCaptionHandler allows to set handler for incoming photos with caption
// Also you can use *bot.RegisterHandler function after bot creation
func WithPhotoCaptionHandler(pattern string, matchType MatchType, handler HandlerFunc) Option {
	return func(b *Bot) {
		b.RegisterHandler(HandlerTypePhotoCaption, pattern, matchType, handler)
	}
}

// WithDefaultHandler allows to set default handler for incoming updates
func WithDefaultHandler(handler HandlerFunc) Option {
	return func(b *Bot) {
		b.defaultHandlerFunc = handler
	}
}

// WithDebug allows to enable debug mode. In debug mode, all requests and responses are logged by debug handler
func WithDebug() Option {
	return func(b *Bot) {
		b.isDebug = true
	}
}

// WithErrorsHandler allows to set handler for errors
func WithErrorsHandler(handler ErrorsHandler) Option {
	return func(b *Bot) {
		b.errorsHandler = handler
	}
}

// WithDebugHandler allows to set handler for debug messages
func WithDebugHandler(handler DebugHandler) Option {
	return func(b *Bot) {
		b.debugHandler = handler
	}
}

// WithHTTPClient allows to set custom http client
func WithHTTPClient(pollTimeout time.Duration, client HttpClient) Option {
	return func(b *Bot) {
		b.pollTimeout = pollTimeout
		b.client = client
	}
}

// WithServerURL allows to set custom server url
func WithServerURL(serverURL string) Option {
	return func(b *Bot) {
		b.url = serverURL
	}
}

// WithSkipGetMe allows skip call GetMe on bot init
func WithSkipGetMe() Option {
	return func(b *Bot) {
		b.skipGetMe = true
	}
}

// WithAllowedUpdates allows to set custom params for getUpdates method
func WithAllowedUpdates(params AllowedUpdates) Option {
	return func(b *Bot) {
		b.allowedUpdates = params
	}
}

// WithUpdatesChannelCap allows setting custom capacity for the Updates channel
func WithUpdatesChannelCap(cap int) Option {
	return func(b *Bot) {
		b.updates = make(chan *models.Update, cap)
	}
}

// WithWebhookSecretToken allows setting X-Telegram-Bot-Api-Secret-Token sent from Telegram servers
func WithWebhookSecretToken(webhookSecretToken string) Option {
	return func(b *Bot) {
		b.webhookSecretToken = webhookSecretToken
	}
}

// WithWorkers allows setting the number of workers that are processing the Updates channel
func WithWorkers(workers int) Option {
	return func(b *Bot) {
		b.workers = workers
	}
}

// UseTestEnvironment allows to use test environment
func UseTestEnvironment() Option {
	return func(b *Bot) {
		b.testEnvironment = true
	}
}

// WithNotAsyncHandlers allows to run handlers in the main goroutine
func WithNotAsyncHandlers() Option {
	return func(b *Bot) {
		b.notAsyncHandlers = true
	}
}

// WithInitialOffset allows to set initial offset for getUpdates method
func WithInitialOffset(offset int64) Option {
	return func(b *Bot) {
		b.lastUpdateID = offset
	}
}
