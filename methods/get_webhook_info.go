package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// GetWebhookInfo https://core.telegram.org/bots/api#getwebhookinfo
func GetWebhookInfo(ctx context.Context, b *bot.Bot) (*models.WebhookInfo, error) {
	result := &models.WebhookInfo{}

	err := bot.RawRequest(ctx, b, "getWebhookInfo", nil, &result)

	return result, err
}
