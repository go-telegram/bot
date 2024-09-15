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
	h := b.defaultHandlerFunc

	defer func() {
		applyMiddlewares(h, b.middlewares...)(ctx, b, upd)
	}()

	h = b.findHandler(upd)
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
