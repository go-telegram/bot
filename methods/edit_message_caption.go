package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type EditMessageCaptionParams struct {
	ChatID                any                    `json:"chat_id"`
	MessageID             int                    `json:"message_id,omitempty"`
	InlineMessageID       string                 `json:"inline_message_id,omitempty"`
	Caption               string                 `json:"caption,omitempty"`
	ParseMode             models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableWebPagePreview bool                   `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

// EditMessageCaption https://core.telegram.org/bots/api#editmessagecaption
func EditMessageCaption(ctx context.Context, b *bot.Bot, params *EditMessageCaptionParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "editMessageCaption", params, result)

	return result, err
}
