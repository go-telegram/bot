package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type BanChatSenderChatParams struct {
	ChatID       any `json:"chat_id" rules:"required,type:string|int"`
	SenderChatID int `json:"sender_chat_id" rules:"required"`
}

// BanChatSenderChat https://core.telegram.org/bots/api#banchatsenderchat
func BanChatSenderChat(ctx context.Context, b *bot.Bot, params *BanChatSenderChatParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "banChatSenderChat", params, &result)

	return result, err
}
