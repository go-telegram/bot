package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type PinChatMessageParams struct {
	ChatID              any  `json:"chat_id" rules:"required,type:string|int"`
	MessageID           int  `json:"message_id" rules:"required"`
	DisableNotification bool `json:"disable_notification,omitempty"`
}

// PinChatMessage https://core.telegram.org/bots/api#pinchatmessage
func PinChatMessage(ctx context.Context, b *bot.Bot, params *PinChatMessageParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "pinChatMessage", params, &result)

	return result, err
}
