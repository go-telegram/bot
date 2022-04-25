package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// https://core.telegram.org/bots/api#forwardmessage

type ForwardMessageParams struct {
	ChatID              string `json:"chat_id"`
	FromChatID          string `json:"from_chat_id"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ProtectContent      bool   `json:"protect_content,omitempty"`
	MessageID           int    `json:"message_id"`
}

func (p ForwardMessageParams) Validate() error {
	if p.ChatID == "" {
		return bot.ErrEmptyChatID
	}
	if p.FromChatID == "" {
		return bot.ErrEmptyFromChatID
	}
	if p.MessageID == 0 {
		return bot.ErrEmptyMessageID
	}
	return nil
}

func (ForwardMessageParams) GetReplyMarkup() models.ReplyMarkup {
	return nil
}

func ForwardMessage(ctx context.Context, b *bot.Bot, params *ForwardMessageParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "forwardMessage", params, result)

	return result, err
}
