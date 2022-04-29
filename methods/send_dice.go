package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SendDiceParams struct {
	ChatID                   any                `json:"chat_id" rules:"required,chat_id"`
	Emoji                    string             `json:"emoji,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendDice https://core.telegram.org/bots/api#senddice
func SendDice(ctx context.Context, b *bot.Bot, params *SendDiceParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendDice", params, &result)

	return result, err
}
