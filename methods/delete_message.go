package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type DeleteMessageParams struct {
	ChatID    string `json:"chat_id"`
	MessageID int    `json:"message_id"`
}

func (p *DeleteMessageParams) Validate() error {
	if p.ChatID == "" {
		return bot.ErrEmptyChatID
	}
	if p.MessageID == 0 {
		return bot.ErrEmptyMessageID
	}
	return nil
}

func DeleteMessage(ctx context.Context, b *bot.Bot, params *DeleteMessageParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "deleteMessage", params, &result)

	return result, err
}
