package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SetGameScoreParams struct {
	UserID             int  `json:"user_id"`
	Score              int  `json:"score"`
	Force              bool `json:"force,omitempty"`
	DisableEditMessage bool `json:"disable_edit_message,omitempty"`
	ChatID             int  `json:"chat_id,omitempty"`
	MessageID          int  `json:"message_id,omitempty"`
	InlineMessageID    int  `json:"inline_message_id,omitempty"`
}

// SetGameScore https://core.telegram.org/bots/api#setgamescore
func SetGameScore(ctx context.Context, b *bot.Bot, params *SetGameScoreParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "setGameScore", params, result)

	return result, err
}
