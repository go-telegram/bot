package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SetChatMenuButtonParams struct {
	ChatID     any                    `json:"chat_id" rules:"type:string|int"`
	MenuButton models.InputMenuButton `json:"menu_button"`
}

// SetChatMenuButton https://core.telegram.org/bots/api#setchatmenubutton
func SetChatMenuButton(ctx context.Context, b *bot.Bot, params *SetChatMenuButtonParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "setChatMenuButton", params, &result)

	return result, err
}
