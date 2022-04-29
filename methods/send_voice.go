package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SendVoiceParams struct {
	ChatID                   any                    `json:"chat_id" rules:"required,type:string|int"`
	Voice                    models.InputFile       `json:"voice" rules:"required"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	Duration                 int                    `json:"duration,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

// SendVoice https://core.telegram.org/bots/api#sendvoice
func SendVoice(ctx context.Context, b *bot.Bot, params *SendVoiceParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendVoice", params, result)

	return result, err
}
