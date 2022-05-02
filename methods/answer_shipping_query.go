package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type AnswerShippingQueryParams struct {
	ShippingQueryID string                  `json:"shipping_query_id"`
	OK              bool                    `json:"ok"`
	ShippingOptions []models.ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    string                  `json:"error_message,omitempty"`
}

// AnswerShippingQuery https://core.telegram.org/bots/api#answershippingquery
func AnswerShippingQuery(ctx context.Context, b *bot.Bot, params *AnswerShippingQueryParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "answerShippingQuery", params, &result)

	return result, err
}
