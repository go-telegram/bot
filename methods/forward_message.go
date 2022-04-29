package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type ForwardMessageParams struct {
	ChatID              any    `json:"chat_id" rules:"required,chat_id"`
	FromChatID          string `json:"from_chat_id" rules:"required"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ProtectContent      bool   `json:"protect_content,omitempty"`
	MessageID           int    `json:"message_id" rules:"required"`
}

// ForwardMessage https://core.telegram.org/bots/api#forwardmessage
func ForwardMessage(ctx context.Context, b *bot.Bot, params *ForwardMessageParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "forwardMessage", params, result)

	return result, err
}
