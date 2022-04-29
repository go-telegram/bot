package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type DeleteMessageParams struct {
	ChatID    any `json:"chat_id" rules:"required,type:string|int"`
	MessageID int `json:"message_id" rules:"required"`
}

// DeleteMessage https://core.telegram.org/bots/api#deletemessage
func DeleteMessage(ctx context.Context, b *bot.Bot, params *DeleteMessageParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "deleteMessage", params, &result)

	return result, err
}
