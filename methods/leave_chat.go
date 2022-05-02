package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type LeaveChatParams struct {
	ChatID any `json:"chat_id"`
}

// LeaveChat https://core.telegram.org/bots/api#leavechat
func LeaveChat(ctx context.Context, b *bot.Bot, params *LeaveChatParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "leaveChat", params, &result)

	return result, err
}
