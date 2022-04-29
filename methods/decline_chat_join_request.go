package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type DeclineChatJoinRequestParams struct {
	ChatID any `json:"chat_id" rules:"required,type:string|int"`
	UserID int `json:"user_id" rules:"required"`
}

// DeclineChatJoinRequest https://core.telegram.org/bots/api#declinechatjoinrequest
func DeclineChatJoinRequest(ctx context.Context, b *bot.Bot, params *DeclineChatJoinRequestParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "declineChatJoinRequest", params, &result)

	return result, err
}
