package bot

import (
	"time"
)

// Option is a function that configures a bot.
type Option func(b *Bot)

// WithCheckInitTimeout sets the timeout for the GetMe call that New makes.
func WithCheckInitTimeout(timeout time.Duration) Option {
	return func(b *Bot) {
		b.checkInitTimeout = timeout
	}
}

// WithMiddlewares adds middlewares that wrap the handler of each incoming update.
// A middleware can run code before and after the next handler, or skip it.
// The first added middleware is the outermost, so they start in the order added.
func WithMiddlewares(middlewares ...Middleware) Option {
	return func(b *Bot) {
		b.middlewares = append(b.middlewares, middlewares...)
	}
}

// WithMessageTextHandler registers a handler for messages whose text matches
// pattern and matchType. To add a handler after creating the bot, use
// Bot.RegisterHandler.
func WithMessageTextHandler(pattern string, matchType MatchType, handler HandlerFunc) Option {
	return func(b *Bot) {
		b.RegisterHandler(HandlerTypeMessageText, pattern, matchType, handler)
	}
}

// WithCallbackQueryDataHandler registers a handler for callback queries whose
// data matches pattern and matchType. To add a handler after creating the bot,
// use Bot.RegisterHandler.
func WithCallbackQueryDataHandler(pattern string, matchType MatchType, handler HandlerFunc) Option {
	return func(b *Bot) {
		b.RegisterHandler(HandlerTypeCallbackQueryData, pattern, matchType, handler)
	}
}

// WithPhotoCaptionHandler registers a handler for messages whose caption
// matches pattern and matchType. It checks the caption of any message type, not
// only photos. To add a handler after creating the bot, use Bot.RegisterHandler.
func WithPhotoCaptionHandler(pattern string, matchType MatchType, handler HandlerFunc) Option {
	return func(b *Bot) {
		b.RegisterHandler(HandlerTypePhotoCaption, pattern, matchType, handler)
	}
}

// WithDefaultHandler sets the handler for updates that no registered handler matches.
func WithDefaultHandler(handler HandlerFunc) Option {
	return func(b *Bot) {
		b.defaultHandlerFunc = handler
	}
}

// WithDebug enables debug mode. The debug handler then receives API requests,
// API responses and webhook payloads, except getUpdates requests and responses
// whose result is an empty array.
func WithDebug() Option {
	return func(b *Bot) {
		b.isDebug = true
	}
}

// WithErrorsHandler sets the callback for errors reported internally by the
// library, such as polling and webhook processing errors. Errors from API
// methods such as SendMessage are returned to the caller and must be handled
// there.
func WithErrorsHandler(handler ErrorsHandler) Option {
	return func(b *Bot) {
		b.errorsHandler = handler
	}
}

// WithDebugHandler sets the handler for debug messages. The bot emits them only
// in debug mode, see WithDebug.
func WithDebugHandler(handler DebugHandler) Option {
	return func(b *Bot) {
		b.debugHandler = handler
	}
}

// WithHTTPClient sets the HTTP client for requests to the Telegram API and the
// polling timeout. A getUpdates request asks Telegram to wait up to pollTimeout
// minus one second.
func WithHTTPClient(pollTimeout time.Duration, client HttpClient) Option {
	return func(b *Bot) {
		b.pollTimeout = pollTimeout
		b.client = client
	}
}

// WithServerURL sets the Bot API server URL, such as that of a local Bot API server.
func WithServerURL(serverURL string) Option {
	return func(b *Bot) {
		b.url = serverURL
	}
}

// WithSkipGetMe allows skip call GetMe on bot init. A MatchTypeCommand or
// MatchTypeCommandStartOnly handler for a bare command such as start then
// matches /start@bot_username only after a GetMe call succeeds.
func WithSkipGetMe() Option {
	return func(b *Bot) {
		b.skipGetMe = true
	}
}

// WithAllowedUpdates sets the update types requested from getUpdates. It applies
// only to polling.
func WithAllowedUpdates(params AllowedUpdates) Option {
	return func(b *Bot) {
		b.allowedUpdates = params
	}
}

// WithUpdatesChannelCap sets the capacity of the queue between the update
// source and the workers. Start fetches the next batch only after the current
// one is dispatched, and a webhook request waits until a worker takes its
// update, so the capacity only bounds how many updates wait for a worker.
// New returns an error for a negative capacity.
func WithUpdatesChannelCap(cap int) Option {
	return func(b *Bot) {
		b.updatesChanCap = cap
	}
}

// WithWebhookSecretToken sets the expected value of the
// X-Telegram-Bot-Api-Secret-Token header. Webhook requests with a different
// value are rejected. An empty token disables the check.
func WithWebhookSecretToken(webhookSecretToken string) Option {
	return func(b *Bot) {
		b.webhookSecretToken = webhookSecretToken
	}
}

// WithWorkers sets the number of workers that take updates from the queue. By
// default a worker starts each handler in a new goroutine, so this number does
// not limit concurrent handlers. With WithNotAsyncHandlers, it limits how many
// queued updates are handled at once. Start and StartWebhook report an error
// and return if workers is not positive.
func WithWorkers(workers int) Option {
	return func(b *Bot) {
		b.workers = workers
	}
}

// UseTestEnvironment sends API requests to the Telegram test environment.
func UseTestEnvironment() Option {
	return func(b *Bot) {
		b.testEnvironment = true
	}
}

// WithNotAsyncHandlers runs each handler synchronously in the calling
// goroutine: the worker for queued updates, or the caller of ProcessUpdate.
// Each worker handles its updates sequentially, but several workers can still
// run handlers concurrently. A WebhookHandler request may wait for a busy worker.
func WithNotAsyncHandlers() Option {
	return func(b *Bot) {
		b.notAsyncHandlers = true
	}
}

// WithInitialOffset sets the ID of the last processed update. The first
// getUpdates request uses offset+1 as its offset.
func WithInitialOffset(offset int64) Option {
	return func(b *Bot) {
		b.lastUpdateID = offset
	}
}
