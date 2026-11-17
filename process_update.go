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

// ProcessUpdate allows you to process update
func (b *Bot) ProcessUpdate(ctx context.Context, upd *models.Update) {
	h := b.findHandler(upd)

	r := applyMiddlewares(h, b.middlewares...)

	if b.notAsyncHandlers {
		r(ctx, b, upd)
		return
	}

	go r(ctx, b, upd)
}

// dispatchUpdate runs the handler of a queued update and calls done when it
// returns.
func (b *Bot) dispatchUpdate(ctx context.Context, upd *models.Update, done func()) {
	r := applyMiddlewares(b.findHandler(upd), b.middlewares...)
	run := func() {
		defer done()
		r(ctx, b, upd)
	}

	if b.notAsyncHandlers {
		run()
		return
	}

	go run()
}

func (b *Bot) findHandler(upd *models.Update) HandlerFunc {
	b.handlersMx.RLock()
	defer b.handlersMx.RUnlock()

	for _, h := range b.handlers {
		if h.match(upd) {
			return h.handler
		}
	}

	return b.defaultHandlerFunc
}
