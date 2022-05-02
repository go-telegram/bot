package methods

import (
	"context"
	"github.com/go-telegram/bot/models"

	"github.com/go-telegram/bot"
)

type GetChatMenuButtonParams struct {
	ChatID any `json:"chat_id"`
}

// GetChatMenuButton https://core.telegram.org/bots/api#getchatmenubutton
func GetChatMenuButton(ctx context.Context, b *bot.Bot, params *GetChatMenuButtonParams) (models.MenuButton, error) {
	var result models.MenuButton

	err := bot.RawRequest(ctx, b, "getChatMenuButton", params, &result)

	return result, err
}
