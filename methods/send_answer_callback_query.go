package methods

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
)

type AnswerCallbackQueryParams struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	URL             string `json:"url,omitempty"`
	CacheTime       int    `json:"cache_time,omitempty"`
}

func (a AnswerCallbackQueryParams) Validate() error {
	if a.CallbackQueryID == "" {
		return fmt.Errorf("callback_query_id is empty")
	}
	return nil
}

func SendAnswerCallbackQuery(ctx context.Context, b *bot.Bot, params *AnswerCallbackQueryParams) (bool, error) {
	var res bool

	err := bot.RawRequest(ctx, b, "answerCallbackQuery", params, &res)

	return res, err
}
