package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SendPhotoParams struct {
	ChatID                   any                    `json:"chat_id" rules:"required,type:string|int"`
	Photo                    models.InputFile       `json:"photo" rules:"required"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

// SendPhoto https://core.telegram.org/bots/api#sendphoto
func SendPhoto(ctx context.Context, b *bot.Bot, params *SendPhotoParams) (*models.Message, error) {
	mes := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendPhoto", params, mes)

	return mes, err
}
