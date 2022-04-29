package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SendContactParams struct {
	ChatID                   any                `json:"chat_id" rules:"required,type:string|int"`
	PhoneNumber              string             `json:"phone_number" rules:"required"`
	FirstName                string             `json:"first_name" rules:"required"`
	LastName                 string             `json:"last_name,omitempty"`
	VCard                    string             `json:"vcard,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendContact https://core.telegram.org/bots/api#sendcontact
func SendContact(ctx context.Context, b *bot.Bot, params *SendContactParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendContact", params, result)

	return result, err
}
