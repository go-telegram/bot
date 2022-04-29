package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type StopMessageLiveLocationParams struct {
	ChatID          string             `json:"chat_id,omitempty"`
	MessageID       int                `json:"message_id,omitempty"`
	InlineMessageID string             `json:"inline_message_id,omitempty"`
	ReplyMarkup     models.ReplyMarkup `json:"reply_markup,omitempty"`
}

// StopMessageLiveLocation https://core.telegram.org/bots/api#stopmessagelivelocation
func StopMessageLiveLocation(ctx context.Context, b *bot.Bot, params *StopMessageLiveLocationParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "stopMessageLiveLocation", params, result)

	return result, err
}
