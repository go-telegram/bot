package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type AnswerInlineQueryParams struct {
	InlineQueryID     string                     `json:"inline_query_id" rules:"required"`
	Results           []models.InlineQueryResult `json:"results" rules:"required"`
	CacheTime         int                        `json:"cache_time,omitempty"`
	IsPersonal        bool                       `json:"is_personal,omitempty"`
	NextOffset        string                     `json:"next_offset,omitempty"`
	SwitchPmText      string                     `json:"switch_pm_text,omitempty"`
	SwitchPmParameter string                     `json:"switch_pm_parameter,omitempty"`
}

// AnswerInlineQuery https://core.telegram.org/bots/api#answerinlinequery
func AnswerInlineQuery(ctx context.Context, b *bot.Bot, params *AnswerInlineQueryParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "answerInlineQuery", params, &result)

	return result, err
}
