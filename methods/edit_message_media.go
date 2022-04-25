package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type EditMessageMediaParams struct {
	ChatID          string             `json:"chat_id,omitempty"`
	MessageID       int                `json:"message_id,omitempty"`
	InlineMessageID int                `json:"inline_message_id,omitempty"`
	Media           InputMedia         `json:"media"`
	ReplyMarkup     models.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p EditMessageMediaParams) Validate() error {
	if p.Media == nil {
		return bot.ErrEmptyMedia
	}
	return nil
}

func EditMessageMedia(ctx context.Context, b *bot.Bot, params *EditMessageMediaParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "editMessageMedia", params, result)

	return result, err
}
