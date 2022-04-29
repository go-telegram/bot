package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type UnbanChatSenderChatParams struct {
	ChatID       any `json:"chat_id" rules:"required,type:string|int"`
	SenderChatID int `json:"sender_chat_id" rules:"required"`
}

// UnbanChatSenderChat https://core.telegram.org/bots/api#unbanchatsenderchat
func UnbanChatSenderChat(ctx context.Context, b *bot.Bot, params *UnbanChatSenderChatParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "unbanChatSenderChat", params, &result)

	return result, err
}
