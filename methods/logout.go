package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

// Logout https://core.telegram.org/bots/api#logout
func Logout(ctx context.Context, b *bot.Bot) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "logout", nil, &result)

	return result, err
}
