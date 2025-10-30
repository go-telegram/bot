package bot

import (
	"time"

	"github.com/go-telegram/bot/models"
)

// Option is a function that configures a bot.
type Option func(b *Bot)

// WithCheckInitTimeout returns Option that redefine CheckInitTimeout.
// This timeout is used to check if bot initialization was successful.
func WithCheckInitTimeout(timeout time.Duration) Option {
	return func(b *Bot) {
		b.checkInitTimeout = timeout
	}
}

// WithMiddlewares returns Option that sets middlewares for each incoming update.
// Middlewares are functions that execute before handlers.
// They are called in the order in which they are added.
func WithMiddlewares(middlewares ...Middleware) Option {
	return func(b *Bot) {
		b.middlewares = append(b.middlewares, middlewares...)
	}
}

// WithMessageTextHandler returns Option that sets handler for incoming text messages.
// This handler will be called when a text message matches the specified pattern and match type.
// Also you can use *bot.RegisterHandler function after bot creation.
func WithMessageTextHandler(pattern string, matchType MatchType, handler HandlerFunc) Option {
	return func(b *Bot) {
		b.RegisterHandler(HandlerTypeMessageText, pattern, matchType, handler)
	}
}

// WithCallbackQueryDataHandler returns Option that sets handler for incoming callback query.
// This handler will be called when a callback query data matches the specified pattern and match type.
// Also you can use *bot.RegisterHandler function after bot creation.
func WithCallbackQueryDataHandler(pattern string, matchType MatchType, handler HandlerFunc) Option {
	return func(b *Bot) {
		b.RegisterHandler(HandlerTypeCallbackQueryData, pattern, matchType, handler)
	}
}

// WithPhotoCaptionHandler returns Option that sets handler for incoming photos with caption.
// This handler will be called when a photo caption matches the specified pattern and match type.
// Also you can use *bot.RegisterHandler function after bot creation.
func WithPhotoCaptionHandler(pattern string, matchType MatchType, handler HandlerFunc) Option {
	return func(b *Bot) {
		b.RegisterHandler(HandlerTypePhotoCaption, pattern, matchType, handler)
	}
}

// WithDefaultHandler returns Option that sets default handler for incoming updates.
// This handler will be called when no other handlers match the incoming update.
func WithDefaultHandler(handler HandlerFunc) Option {
	return func(b *Bot) {
		b.defaultHandlerFunc = handler
	}
}

// WithDebug returns Option that enables debug mode.
// In debug mode, all requests and responses are logged by debug handler.
func WithDebug() Option {
	return func(b *Bot) {
		b.isDebug = true
	}
}

// WithErrorsHandler returns Option that sets handler for errors.
// This handler will be called whenever an error occurs during bot operation.
func WithErrorsHandler(handler ErrorsHandler) Option {
	return func(b *Bot) {
		b.errorsHandler = handler
	}
}

// WithDebugHandler returns Option that sets handler for debug messages.
// This handler will be called to process debug information when debug mode is enabled.
func WithDebugHandler(handler DebugHandler) Option {
	return func(b *Bot) {
		b.debugHandler = handler
	}
}

// WithHTTPClient returns Option that sets custom http client.
// This enables using custom HTTP configuration for all requests to Telegram API.
func WithHTTPClient(pollTimeout time.Duration, client HttpClient) Option {
	return func(b *Bot) {
		b.pollTimeout = pollTimeout
		b.client = client
	}
}

// WithServerURL returns Option that sets custom server url.
// This is useful for using local Bot API server or alternative endpoints.
func WithServerURL(serverURL string) Option {
	return func(b *Bot) {
		b.url = serverURL
	}
}

// WithSkipGetMe returns Option that skips call GetMe on bot init.
// This can speed up bot initialization when bot information is not needed.
func WithSkipGetMe() Option {
	return func(b *Bot) {
		b.skipGetMe = true
	}
}

// WithAllowedUpdates returns Option that sets custom params for getUpdates method.
// This helps to filter which update types the bot should receive.
func WithAllowedUpdates(params AllowedUpdates) Option {
	return func(b *Bot) {
		b.allowedUpdates = params
	}
}

// WithUpdatesChannelCap returns Option that sets custom capacity for the Updates channel.
// Higher capacity can improve performance under high load at the cost of memory usage.
func WithUpdatesChannelCap(cap int) Option {
	return func(b *Bot) {
		b.updates = make(chan *models.Update, cap)
	}
}

// WithWebhookSecretToken returns Option that sets X-Telegram-Bot-Api-Secret-Token sent from Telegram servers.
// This token is used to verify that webhook requests actually come from Telegram.
func WithWebhookSecretToken(webhookSecretToken string) Option {
	return func(b *Bot) {
		b.webhookSecretToken = webhookSecretToken
	}
}

// WithWorkers returns Option that sets the number of workers that are processing the Updates channel.
// More workers can improve concurrent update processing performance.
func WithWorkers(workers int) Option {
	return func(b *Bot) {
		b.workers = workers
	}
}

// UseTestEnvironment returns Option that enables test environment.
// This switches the bot to use Telegram's test servers for development and testing.
func UseTestEnvironment() Option {
	return func(b *Bot) {
		b.testEnvironment = true
	}
}

// WithNotAsyncHandlers returns Option that runs handlers in the main goroutine.
// This ensures handlers execute sequentially instead of concurrently.
func WithNotAsyncHandlers() Option {
	return func(b *Bot) {
		b.notAsyncHandlers = true
	}
}

// WithInitialOffset returns Option that sets initial offset for getUpdates method.
// This determines which updates to skip when starting to receive new updates.
func WithInitialOffset(offset int64) Option {
	return func(b *Bot) {
		b.lastUpdateID = offset
	}
}
