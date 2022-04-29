package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type StopPollParams struct {
	ChatID      any                `json:"chat_id" rules:"required,type:string|int"`
	MessageID   int                `json:"message_id" rules:"required"`
	ReplyMarkup models.ReplyMarkup `json:"reply_markup,omitempty"`
}

// StopPoll https://core.telegram.org/bots/api#stoppoll
func StopPoll(ctx context.Context, b *bot.Bot, params *StopPollParams) (*models.Poll, error) {
	result := &models.Poll{}

	err := bot.RawRequest(ctx, b, "stopPoll", params, &result)

	return result, err
}
