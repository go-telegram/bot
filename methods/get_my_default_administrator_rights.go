package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type GetMyDefaultAdministratorRightsParams struct {
	ForChannels bool `json:"for_channels,omitempty"`
}

// GetMyDefaultAdministratorRights https://core.telegram.org/bots/api#getmydefaultadministratorrights
func GetMyDefaultAdministratorRights(ctx context.Context, b *bot.Bot, params *GetMyDefaultAdministratorRightsParams) (*models.ChatAdministratorRights, error) {
	result := &models.ChatAdministratorRights{}

	err := bot.RawRequest(ctx, b, "setMyDefaultAdministratorRights", params, &result)

	return result, err
}
