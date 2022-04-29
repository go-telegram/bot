package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type AnswerCallbackQueryParams struct {
	CallbackQueryID string `json:"callback_query_id" rules:"required"`
	Text            string `json:"text,omitempty" rules:"max:200"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	URL             string `json:"url,omitempty"`
	CacheTime       int    `json:"cache_time,omitempty"`
}

// AnswerCallbackQuery https://core.telegram.org/bots/api#answercallbackquery
func AnswerCallbackQuery(ctx context.Context, b *bot.Bot, params *AnswerCallbackQueryParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "answerCallbackQuery", params, &result)

	return result, err
}
