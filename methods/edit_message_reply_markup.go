package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type EditMessageReplyMarkupParams struct {
	ChatID          string             `json:"chat_id,omitempty"`
	MessageID       int                `json:"message_id,omitempty"`
	InlineMessageID string             `json:"inline_message_id,omitempty"`
	ReplyMarkup     models.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p EditMessageReplyMarkupParams) Validate() error {
	if p.InlineMessageID == "" {
		if p.ChatID == "" || p.MessageID == 0 {
			return bot.ErrEmptyChatID
		}
	}
	return nil
}

func EditMessageReplyMarkup(ctx context.Context, b *bot.Bot, params *EditMessageReplyMarkupParams) (*models.Message, error) {
	mes := &models.Message{}

	err := bot.RawRequest(ctx, b, "editMessageReplyMarkup", params, mes)

	return mes, err
}
