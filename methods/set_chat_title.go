package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type SetChatTitleParams struct {
	ChatID any    `json:"chat_id"`
	Title  string `json:"title"`
}

// SetChatTitle https://core.telegram.org/bots/api#setchattitle
func SetChatTitle(ctx context.Context, b *bot.Bot, params *SetChatTitleParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "setChatTitle", params, &result)

	return result, err
}
