package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type SetChatAdministratorCustomTitleParams struct {
	ChatID      any    `json:"chat_id" rules:"required,chat_id"`
	UserID      int    `json:"user_id" rules:"required"`
	CustomTitle string `json:"custom_title" rules:"required"`
}

// SetChatAdministratorCustomTitle https://core.telegram.org/bots/api#setchatadministratorcustomtitle
func SetChatAdministratorCustomTitle(ctx context.Context, b *bot.Bot, params *SetChatAdministratorCustomTitleParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "setChatAdministratorCustomTitle", params, &result)

	return result, err
}
