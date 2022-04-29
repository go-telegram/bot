package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type ApproveChatJoinRequestParams struct {
	ChatID any `json:"chat_id" rules:"required,chat_id"`
	UserID int `json:"user_id" rules:"required"`
}

// ApproveChatJoinRequest https://core.telegram.org/bots/api#approvechatjoinrequest
func ApproveChatJoinRequest(ctx context.Context, b *bot.Bot, params *ApproveChatJoinRequestParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "approveChatJoinRequest", params, &result)

	return result, err
}
