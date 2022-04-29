package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SendMediaGroupParams struct {
	ChatID                   any                 `json:"chat_id" rules:"required,chat_id"`
	Media                    []models.InputMedia `json:"media" rules:"required,min:1"`
	DisableNotification      bool                `json:"disable_notification,omitempty"`
	ProtectContent           bool                `json:"protect_content,omitempty"`
	ReplyToMessageID         int                 `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                `json:"allow_sending_without_reply,omitempty"`
}

// SendMediaGroup https://core.telegram.org/bots/api#sendmediagroup
func SendMediaGroup(ctx context.Context, b *bot.Bot, params *SendMediaGroupParams) ([]*models.Message, error) {
	var res []*models.Message

	err := bot.RawRequest(ctx, b, "sendMediaGroup", params, &res)

	return res, err
}
