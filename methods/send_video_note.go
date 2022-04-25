package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// https://core.telegram.org/bots/api#sendvideonote

type SendVideoNoteParams struct {
	ChatID                   string             `json:"chat_id"`
	VideoNote                string             `json:"video_note"`
	Duration                 int                `json:"duration,omitempty"`
	Length                   int                `json:"length,omitempty"`
	Thumb                    string             `json:"thumb,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p SendVideoNoteParams) Validate() error {
	if p.ChatID == "" {
		return bot.ErrEmptyChatID
	}
	if p.VideoNote == "" {
		return bot.ErrEmptyVideoNote
	}
	return nil
}

func (p SendVideoNoteParams) GetReplyMarkup() models.ReplyMarkup {
	return p.ReplyMarkup
}

func SendVideoNote(ctx context.Context, b *bot.Bot, params *SendVideoNoteParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendVideoNote", params, result)

	return result, err
}
