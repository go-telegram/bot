package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type GetChatMemberParams struct {
	ChatID any `json:"chat_id" rules:"required,chat_id"`
	UserID int `json:"user_id" rules:"required"`
}

// GetChatMember https://core.telegram.org/bots/api#getchatmember
func GetChatMember(ctx context.Context, b *bot.Bot, params *GetChatMemberParams) (*models.ChatMember, error) {
	res := &models.ChatMember{}

	err := bot.RawRequest(ctx, b, "getChatMember", params, &res)

	return res, err
}
