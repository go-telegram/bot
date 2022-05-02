package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SetMyCommandsParams struct {
	Commands     []models.BotCommand    `json:"commands"`
	Scope        models.BotCommandScope `json:"scope,omitempty"`
	LanguageCode string                 `json:"language_code,omitempty"`
}

// SetMyCommands https://core.telegram.org/bots/api#setmycommands
func SetMyCommands(ctx context.Context, b *bot.Bot, params *SetMyCommandsParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "setMyCommands", params, &result)

	return result, err
}
