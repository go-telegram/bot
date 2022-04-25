package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type PinChatMessageParams struct {
	ChatID              string `json:"chat_id"`
	MessageID           int    `json:"message_id"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
}

func (p PinChatMessageParams) Validate() error {
	if p.ChatID == "" {
		return bot.ErrEmptyChatID
	}
	return nil
}

func PinChatMessage(ctx context.Context, b *bot.Bot, params *PinChatMessageParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "pinChatMessage", params, &result)

	return result, err
}
