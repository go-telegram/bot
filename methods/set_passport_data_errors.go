package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SetPassportDataErrorsParams struct {
	UserID int                           `json:"user_id"`
	Errors []models.PassportElementError `json:"errors"`
}

// SetPassportDataErrors https://core.telegram.org/bots/api#setpassportdataerrors
func SetPassportDataErrors(ctx context.Context, b *bot.Bot, params *SetPassportDataErrorsParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "setPassportDataErrors", params, &result)

	return result, err
}
