package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type UnpinChatMessageParams struct {
	ChatID    any `json:"chat_id" rules:"required,chat_id"`
	MessageID int `json:"message_id,omitempty"`
}

// UnpinChatMessage https://core.telegram.org/bots/api#unpinchatmessage
func UnpinChatMessage(ctx context.Context, b *bot.Bot, params *UnpinChatMessageParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "unpinChatMessage", params, &result)

	return result, err
}
