package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type EditMessageReplyMarkupParams struct {
	ChatID          any                `json:"chat_id"`
	MessageID       int                `json:"message_id,omitempty"`
	InlineMessageID string             `json:"inline_message_id,omitempty"`
	ReplyMarkup     models.ReplyMarkup `json:"reply_markup,omitempty"`
}

// EditMessageReplyMarkup https://core.telegram.org/bots/api#editmessagereplymarkup
func EditMessageReplyMarkup(ctx context.Context, b *bot.Bot, params *EditMessageReplyMarkupParams) (*models.Message, error) {
	mes := &models.Message{}

	err := bot.RawRequest(ctx, b, "editMessageReplyMarkup", params, mes)

	return mes, err
}
