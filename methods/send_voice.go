package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// https://core.telegram.org/bots/api#sendaudio

type SendVoiceParams struct {
	ChatID                   string                 `json:"chat_id"`
	Voice                    string                 `json:"voice"`
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

func (p SendVoiceParams) Validate() error {
	if p.ChatID == "" {
		return bot.ErrEmptyChatID
	}
	if p.Voice == "" {
		return bot.ErrEmptyVoice
	}
	return nil
}

func SendVoice(ctx context.Context, b *bot.Bot, params *SendVoiceParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendVoice", params, result)

	return result, err
}
