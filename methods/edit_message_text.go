package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// https://core.telegram.org/bots/api#editmessagetext

type EditMessageTextParams struct {
	ChatID                string                 `json:"chat_id"`
	MessageID             int                    `json:"message_id,omitempty"`
	InlineMessageID       string                 `json:"inline_message_id,omitempty"`
	Text                  string                 `json:"text"`
	ParseMode             models.ParseMode       `json:"parse_mode,omitempty"`
	Entities              []models.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview bool                   `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

func (p EditMessageTextParams) Validate() error {
	if p.ChatID == "" {
		return bot.ErrEmptyChatID
	}
	if p.Text == "" {
		return bot.ErrEmptyText
	}
	return nil
}

func EditMessageText(ctx context.Context, b *bot.Bot, params *EditMessageTextParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "editMessageText", params, result)

	return result, err
}
