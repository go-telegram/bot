package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type AnswerPreCheckoutQueryParams struct {
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`
	OK                 bool   `json:"ok"`
	ErrorMessage       string `json:"error_message,omitempty"`
}

// AnswerPreCheckoutQuery https://core.telegram.org/bots/api#answerprecheckoutquery
func AnswerPreCheckoutQuery(ctx context.Context, b *bot.Bot, params *AnswerPreCheckoutQueryParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "answerPreCheckoutQuery", params, &result)

	return result, err
}
