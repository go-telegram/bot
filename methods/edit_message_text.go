package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type EditMessageTextParams struct {
	ChatID                any                    `json:"chat_id" rules:"required_if_empty:InlineMessageID,chat_id"`
	MessageID             int                    `json:"message_id,omitempty" rules:"required_if_empty:InlineMessageID"`
	InlineMessageID       string                 `json:"inline_message_id,omitempty" rules:"required_if_empty:ChatID|MessageID"`
	Text                  string                 `json:"text" rules:"required"`
	ParseMode             models.ParseMode       `json:"parse_mode,omitempty"`
	Entities              []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableWebPagePreview bool                   `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

// EditMessageText https://core.telegram.org/bots/api#editmessagetext
func EditMessageText(ctx context.Context, b *bot.Bot, params *EditMessageTextParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "editMessageText", params, result)

	return result, err
}
