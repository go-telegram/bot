package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type SendChatActionParams struct {
	ChatID any    `json:"chat_id" rules:"required,type:string|int"`
	Action string `json:"action" rules:"required"`
}

// SendChatAction https://core.telegram.org/bots/api#sendchataction
func SendChatAction(ctx context.Context, b *bot.Bot, params *SendChatActionParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "sendDice", params, &result)

	return result, err
}
