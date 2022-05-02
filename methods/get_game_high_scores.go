package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type GetGameHighScoresParams struct {
	UserID          int `json:"user_id"`
	ChatID          int `json:"chat_id,omitempty"`
	MessageID       int `json:"message_id,omitempty"`
	InlineMessageID int `json:"inline_message_id,omitempty"`
}

// GetGameHighScores https://core.telegram.org/bots/api#getgamehighscores
func GetGameHighScores(ctx context.Context, b *bot.Bot, params *GetGameHighScoresParams) (*models.GameHighScore, error) {
	result := &models.GameHighScore{}

	err := bot.RawRequest(ctx, b, "getGameHighScores", params, result)

	return result, err
}
