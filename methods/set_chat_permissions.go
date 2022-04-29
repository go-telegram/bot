package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SetChatPermissionsParams struct {
	ChatID      any                    `json:"chat_id" rules:"required,chat_id"`
	Permissions models.ChatPermissions `json:"permissions"`
}

// SetChatPermissions https://core.telegram.org/bots/api#setchatpermissions
func SetChatPermissions(ctx context.Context, b *bot.Bot, params *SetChatPermissionsParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "setChatPermissions", params, &result)

	return result, err
}
