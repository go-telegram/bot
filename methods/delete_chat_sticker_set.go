package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type DeleteChatStickerSetParams struct {
	ChatID         any    `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

// DeleteChatStickerSet https://core.telegram.org/bots/api#deletechatstickerset
func DeleteChatStickerSet(ctx context.Context, b *bot.Bot, params *DeleteChatStickerSetParams) (bool, error) {
	var res bool

	err := bot.RawRequest(ctx, b, "deleteChatStickerSet", params, &res)

	return res, err
}
