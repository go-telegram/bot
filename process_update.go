package bot

import (
	"context"

	"github.com/go-telegram/bot/models"
)

func applyMiddlewares(h HandlerFunc, m ...Middleware) HandlerFunc {
	if len(m) < 1 {
		return h
	}
	wrapped := h
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}
	return wrapped
}

// ProcessUpdate finds the handler for upd, applies the bot middlewares and runs
// it. The handler runs in a new goroutine, or in the calling goroutine with
// WithNotAsyncHandlers.
func (b *Bot) ProcessUpdate(ctx context.Context, upd *models.Update) {
	h := b.findHandler(upd)

	r := applyMiddlewares(h, b.middlewares...)

	if b.notAsyncHandlers {
		r(ctx, b, upd)
		return
	}

	go r(ctx, b, upd)
}

func (b *Bot) findHandler(upd *models.Update) HandlerFunc {
	b.handlersMx.RLock()
	defer b.handlersMx.RUnlock()

	usernames := b.botUsernames()
	for _, h := range b.handlers {
		if h.match(upd, usernames...) {
			return h.handler
		}
	}

	return b.defaultHandlerFunc
}
