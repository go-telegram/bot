package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type UnbanChatMemberParams struct {
	ChatID       any  `json:"chat_id" rules:"required,chat_id"`
	UserID       int  `json:"user_id" rules:"required"`
	OnlyIfBanned bool `json:"only_if_banned,omitempty"`
}

// UnbanChatMember https://core.telegram.org/bots/api#unbanchatmember
func UnbanChatMember(ctx context.Context, b *bot.Bot, params *UnbanChatMemberParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "unbanChatMember", params, &result)

	return result, err
}
