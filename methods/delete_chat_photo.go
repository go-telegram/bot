package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type DeleteChatPhotoParams struct {
	ChatID any `json:"chat_id" rules:"required,type:string|int"`
}

// DeleteChatPhoto https://core.telegram.org/bots/api#deletechatphoto
func DeleteChatPhoto(ctx context.Context, b *bot.Bot, params *DeleteChatPhotoParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "deleteChatPhoto", params, &result)

	return result, err
}
