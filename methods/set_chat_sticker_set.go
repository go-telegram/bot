package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type SetChatStickerSetParams struct {
	ChatID         any    `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

// SetChatStickerSet https://core.telegram.org/bots/api#setchatstickerset
func SetChatStickerSet(ctx context.Context, b *bot.Bot, params *SetChatStickerSetParams) (bool, error) {
	var res bool

	err := bot.RawRequest(ctx, b, "setChatStickerSet", params, &res)

	return res, err
}
