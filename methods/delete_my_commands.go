package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type DeleteMyCommandsParams struct {
	Scope        models.BotCommandScope `json:"scope,omitempty"`
	LanguageCode string                 `json:"language_code,omitempty"`
}

// DeleteMyCommands https://core.telegram.org/bots/api#deletemycommands
func DeleteMyCommands(ctx context.Context, b *bot.Bot, params *DeleteMyCommandsParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "deleteMyCommands", params, &result)

	return result, err
}
