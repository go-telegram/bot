package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type GetChatParams struct {
	ChatID any `json:"chat_id" rules:"required,chat_id"`
}

// GetChat https://core.telegram.org/bots/api#getchat
func GetChat(ctx context.Context, b *bot.Bot, params *GetChatParams) (*models.Chat, error) {
	var res *models.Chat

	err := bot.RawRequest(ctx, b, "getChat", params, &res)

	return res, err
}
