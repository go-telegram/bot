package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type SetChatDescriptionParams struct {
	ChatID      any    `json:"chat_id" rules:"required,type:string|int"`
	Description string `json:"title" rules:"max:255"`
}

// SetChatDescription https://core.telegram.org/bots/api#setchatdescription
func SetChatDescription(ctx context.Context, b *bot.Bot, params *SetChatDescriptionParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "setChatDescription", params, &result)

	return result, err
}
