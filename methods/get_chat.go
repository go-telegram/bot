package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type GetChatParams struct {
	ChatID string `json:"chat_id"`
}

func (p *GetChatParams) Validate() error {
	if p.ChatID == "" {
		return bot.ErrEmptyChatID
	}
	return nil
}

func GetChat(ctx context.Context, b *bot.Bot, params *GetChatParams) (*models.Chat, error) {
	var res *models.Chat

	err := bot.RawRequest(ctx, b, "getChat", params, &res)

	return res, err
}
