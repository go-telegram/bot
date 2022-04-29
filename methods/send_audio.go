package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SendAudioParams struct {
	ChatID                   any                    `json:"chat_id" rules:"required,chat_id"`
	Audio                    models.InputFile       `json:"audio" rules:"required"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	Duration                 int                    `json:"duration,omitempty"`
	Performer                string                 `json:"performer,omitempty"`
	Title                    string                 `json:"title,omitempty"`
	Thumb                    string                 `json:"thumb,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

// SendAudio https://core.telegram.org/bots/api#sendaudio
func SendAudio(ctx context.Context, b *bot.Bot, params *SendAudioParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendAudio", params, result)

	return result, err
}
