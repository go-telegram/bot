package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SetWebhookParams struct {
	URL                string           `json:"url"`
	Certificate        models.InputFile `json:"certificate,omitempty"`
	IPAddress          string           `json:"ip_address,omitempty"`
	MaxConnections     int              `json:"max_connections,omitempty"`
	AllowedUpdates     []string         `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool             `json:"drop_pending_updates,omitempty"`
}

// SetWebhook https://core.telegram.org/bots/api#setwebhook
func SetWebhook(ctx context.Context, b *bot.Bot, params *SetWebhookParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "setWebhook", params, &result)

	return result, err
}
