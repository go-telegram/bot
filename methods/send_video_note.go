package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SendVideoNoteParams struct {
	ChatID                   any                `json:"chat_id"`
	VideoNote                models.InputFile   `json:"video_note"`
	Duration                 int                `json:"duration,omitempty"`
	Length                   int                `json:"length,omitempty"`
	Thumb                    string             `json:"thumb,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendVideoNote https://core.telegram.org/bots/api#sendvideonote
func SendVideoNote(ctx context.Context, b *bot.Bot, params *SendVideoNoteParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendVideoNote", params, result)

	return result, err
}
