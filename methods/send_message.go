package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SendMessageParams struct {
	ChatID                   any                    `json:"chat_id" rules:"required,chat_id"`
	Text                     string                 `json:"text" rules:"required"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	Entities                 []models.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    bool                   `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

// SendMessage https://core.telegram.org/bots/api#sendmessage
func SendMessage(ctx context.Context, b *bot.Bot, params *SendMessageParams) (*models.Message, error) {
	mes := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendMessage", params, mes)

	return mes, err
}
