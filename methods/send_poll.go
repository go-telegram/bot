package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SendPollParams struct {
	ChatID                   any                    `json:"chat_id" rules:"required,chat_id"`
	Question                 string                 `json:"question" rules:"required"`
	Options                  []string               `json:"options" rules:"required,min:2,max:10"`
	IsAnonymous              bool                   `json:"is_anonymous,omitempty"`
	Type                     string                 `json:"type,omitempty"`
	AllowsMultipleAnswers    bool                   `json:"allows_multiple_answers,omitempty"`
	CorrectOptionID          int                    `json:"correct_option_id"`
	Explanation              string                 `json:"explanation,omitempty"`
	ExplanationParseMode     string                 `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities      []models.MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod               int                    `json:"open_period,omitempty"`
	CloseDate                int                    `json:"close_date,omitempty"`
	IsClosed                 bool                   `json:"is_closed,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

// SendPoll https://core.telegram.org/bots/api#sendpoll
func SendPoll(ctx context.Context, b *bot.Bot, params *SendPollParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendPoll", params, result)

	return result, err
}
