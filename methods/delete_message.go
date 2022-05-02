package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type DeleteMessageParams struct {
	ChatID    any `json:"chat_id"`
	MessageID int `json:"message_id"`
}

// DeleteMessage https://core.telegram.org/bots/api#deletemessage
func DeleteMessage(ctx context.Context, b *bot.Bot, params *DeleteMessageParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "deleteMessage", params, &result)

	return result, err
}
