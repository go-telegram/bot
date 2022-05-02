package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SendAnimationParams struct {
	ChatID                   any                    `json:"chat_id"`
	Animation                models.InputFile       `json:"animation"`
	Duration                 int                    `json:"duration,omitempty"`
	Width                    int                    `json:"width,omitempty"`
	Height                   int                    `json:"height,omitempty"`
	Thumb                    string                 `json:"thumb,omitempty"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

// SendAnimation https://core.telegram.org/bots/api#sendanimation
func SendAnimation(ctx context.Context, b *bot.Bot, params *SendAnimationParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendAnimation", params, result)

	return result, err
}
