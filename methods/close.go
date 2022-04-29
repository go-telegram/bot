package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

// Close https://core.telegram.org/bots/api#close
func Close(ctx context.Context, b *bot.Bot) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "close", nil, &result)

	return result, err
}
