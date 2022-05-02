package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SendLocationParams struct {
	ChatID                   any                `json:"chat_id"`
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

// SendLocation https://core.telegram.org/bots/api#sendlocation
func SendLocation(ctx context.Context, b *bot.Bot, params *SendLocationParams) (*models.Message, error) {
	mes := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendLocation", params, mes)

	return mes, err
}
