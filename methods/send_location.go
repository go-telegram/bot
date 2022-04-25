package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// https://core.telegram.org/bots/api#sendlocation

type SendLocationParams struct {
	ChatID                   string             `json:"chat_id"`
	Latitude                 float64            `json:"latitude"`
	Longitude                float64            `json:"longitude"`
	HorizontalAccuracy       float64            `json:"horizontal_accuracy,omitempty"`
	LivePeriod               int                `json:"live_period,omitempty"`
	Heading                  int                `json:"heading,omitempty"`
	ProximityAlertRadius     int                `json:"proximity_alert_radius,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (p SendLocationParams) Validate() error {
	if p.ChatID == "" {
		return bot.ErrEmptyChatID
	}
	if p.Latitude == 0 {
		return bot.ErrEmptyLatitude
	}
	if p.Longitude == 0 {
		return bot.ErrEmptyLongitude
	}
	return nil
}

func (p SendLocationParams) GetReplyMarkup() models.ReplyMarkup {
	return p.ReplyMarkup
}

func SendLocation(ctx context.Context, b *bot.Bot, params *SendLocationParams) (*models.Message, error) {
	mes := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendLocation", params, mes)

	return mes, err
}
