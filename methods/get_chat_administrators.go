package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type GetChatAdministratorsParams struct {
	ChatID any `json:"chat_id" rules:"required,type:string|int"`
}

// GetChatAdministrators https://core.telegram.org/bots/api#getchatadministrators
func GetChatAdministrators(ctx context.Context, b *bot.Bot, params *GetChatAdministratorsParams) ([]models.ChatMember, error) {
	var res []models.ChatMember

	err := bot.RawRequest(ctx, b, "getChatAdministrators", params, &res)

	return res, err
}
