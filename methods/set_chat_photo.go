package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SetChatPhotoParams struct {
	ChatID any              `json:"chat_id" rules:"required,type:string|int"`
	Photo  models.InputFile `json:"photo" rules:"required"`
}

// SetChatPhoto https://core.telegram.org/bots/api#setchatphoto
func SetChatPhoto(ctx context.Context, b *bot.Bot, params *SetChatPhotoParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "setChatPhoto", params, &result)

	return result, err
}
