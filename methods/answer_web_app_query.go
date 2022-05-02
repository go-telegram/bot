package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type AnswerWebAppQueryParams struct {
	WebAppQueryID string                   `json:"web_app_query_id"`
	Result        models.InlineQueryResult `json:"result"`
}

// AnswerWebAppQuery https://core.telegram.org/bots/api#answerwebappquery
func AnswerWebAppQuery(ctx context.Context, b *bot.Bot, params *AnswerWebAppQueryParams) (*models.SentWebAppMessage, error) {
	result := &models.SentWebAppMessage{}

	err := bot.RawRequest(ctx, b, "answerWebAppQuery", params, result)

	return result, err
}
