package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type RestrictChatMemberParams struct {
	ChatID      any                     `json:"chat_id" rules:"required,chat_id"`
	UserID      int                     `json:"user_id" rules:"required"`
	Permissions *models.ChatPermissions `json:"permissions,omitempty" rules:"required"`
	UntilDate   int                     `json:"until_date,omitempty"`
}

// RestrictChatMember https://core.telegram.org/bots/api#restrictchatmember
func RestrictChatMember(ctx context.Context, b *bot.Bot, params *RestrictChatMemberParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "restrictChatMember", params, &result)

	return result, err
}
