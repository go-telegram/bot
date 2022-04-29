package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SetMyDefaultAdministratorRightsParams struct {
	Rights      *models.ChatAdministratorRights `json:"rights,omitempty"`
	ForChannels bool                            `json:"for_channels,omitempty"`
}

// SetMyDefaultAdministratorRights https://core.telegram.org/bots/api#setmydefaultadministratorrights
func SetMyDefaultAdministratorRights(ctx context.Context, b *bot.Bot, params *SetMyDefaultAdministratorRightsParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "setMyDefaultAdministratorRights", params, &result)

	return result, err
}
