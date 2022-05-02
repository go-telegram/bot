package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type DeleteWebhookParams struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

// DeleteWebhook https://core.telegram.org/bots/api#deletewebhook
func DeleteWebhook(ctx context.Context, b *bot.Bot, params *DeleteWebhookParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "deleteWebhook", params, &result)

	return result, err
}
