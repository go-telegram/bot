package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type EditMessageMediaParams struct {
	ChatID          any                `json:"chat_id"`
	MessageID       int                `json:"message_id,omitempty"`
	InlineMessageID string             `json:"inline_message_id,omitempty"`
	Media           models.InputMedia  `json:"media"`
	ReplyMarkup     models.ReplyMarkup `json:"reply_markup,omitempty"`
}

// EditMessageMedia https://core.telegram.org/bots/api#editmessagemedia
func EditMessageMedia(ctx context.Context, b *bot.Bot, params *EditMessageMediaParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "editMessageMedia", params, result)

	return result, err
}
