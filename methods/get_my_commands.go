package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type GetMyCommandsParams struct {
	Scope        BotCommandScope `json:"scope,omitempty"`
	LanguageCode string          `json:"language_code,omitempty"`
}

// GetMyCommands https://core.telegram.org/bots/api#getmycommands
func GetMyCommands(ctx context.Context, b *bot.Bot, params *GetMyCommandsParams) ([]models.BotCommand, error) {
	var result []models.BotCommand

	err := bot.RawRequest(ctx, b, "getMyCommands", params, &result)

	return result, err
}
