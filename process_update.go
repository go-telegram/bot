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

func (b *Bot) processUpdate(ctx context.Context, upd *models.Update) {
	h := b.defaultHandlerFunc

	defer func() {
		applyMiddlewares(h, b.middlewares...)(ctx, b, upd)
	}()

	// Message
	if upd.Message != nil {
		h = b.findHandler(HandlerTypeMessageText, upd.Message.Text)
		return
	}
	// EditedMessage
	// ChannelPost
	// EditedChannelPost
	// InlineQuery
	// ChosenInlineResult
	// CallbackQuery
	if upd.CallbackQuery != nil {
		h = b.findHandler(HandlerTypeCallbackQueryData, upd.CallbackQuery.Data)
		return
	}
	// ShippingQuery
	// PreCheckoutQuery
	// Poll
	// PollAnswer
	// MyChatMember
	// ChatMember
	// ChatJoinRequest
}

func (b *Bot) findHandler(handlerType HandlerType, pattern string) HandlerFunc {
	b.handlersMx.RLock()
	defer b.handlersMx.RUnlock()

	for _, h := range b.handlers {
		if h.handlerType == handlerType {
			if h.match(pattern) {
				return h.handler
			}
		}
	}

	return b.defaultHandlerFunc
}
