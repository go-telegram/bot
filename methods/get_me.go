package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// GetMe https://core.telegram.org/bots/api#getme
func GetMe(ctx context.Context, b *bot.Bot) (*models.User, error) {
	result := &models.User{}

	err := bot.RawRequest(ctx, b, "getMe", nil, result)

	return result, err
}
