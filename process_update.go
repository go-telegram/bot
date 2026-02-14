package bot

import (
	"context"

	"github.com/go-telegram/bot/internal/machine"
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
	h := b.findHandler(ctx, upd)

	r := applyMiddlewares(h, b.middlewares...)

	if b.notAsyncHandlers {
		r(ctx, b, upd)
		return
	}

	go r(ctx, b, upd)
}

func (b *Bot) findHandler(ctx context.Context, upd *models.Update) HandlerFunc {
	b.handlersMx.RLock()
	defer b.handlersMx.RUnlock()

	var (
		s   = states{def: machine.DefaultState, cur: machine.DefaultState}
		err error
	)

	uid, ok := ExtractUserID(upd)

	if b.machine != nil && ok {
		s.def = b.machine.DefaultState()
		s.cur, err = b.machine.Current(ctx, uid)
		if err != nil {
			return b.defaultHandlerFunc
		}
	}

	for _, h := range b.handlers {
		if h.match(upd, s) {
			return h.handler
		}
	}

	return b.defaultHandlerFunc
}
