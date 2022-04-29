package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type GetChatMemberCountParams struct {
	ChatID any `json:"chat_id" rules:"required,type:string|int"`
}

// GetChatMemberCount https://core.telegram.org/bots/api#getchatmembercount
func GetChatMemberCount(ctx context.Context, b *bot.Bot, params *GetChatMemberCountParams) (int, error) {
	var res int

	err := bot.RawRequest(ctx, b, "getChatMemberCount", params, &res)

	return res, err
}
