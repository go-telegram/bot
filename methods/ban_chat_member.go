package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type BanChatMemberParams struct {
	ChatID         any  `json:"chat_id" rules:"required,type:string|int"`
	UserID         int  `json:"user_id" rules:"required"`
	UntilDate      int  `json:"until_date,omitempty"`
	RevokeMessages bool `json:"revoke_messages,omitempty"`
}

// BanChatMember https://core.telegram.org/bots/api#banchatmember
func BanChatMember(ctx context.Context, b *bot.Bot, params *BanChatMemberParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "banChatMember", params, &result)

	return result, err
}
