package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type EditMessageLiveLocationParams struct {
	ChatID               any                `json:"chat_id" rules:"type:string|int"`
	MessageID            int                `json:"message_id,omitempty"`
	InlineMessageID      string             `json:"inline_message_id,omitempty"`
	Latitude             float64            `json:"latitude" rules:"required"`
	Longitude            float64            `json:"longitude" rules:"required"`
	HorizontalAccuracy   float64            `json:"horizontal_accuracy,omitempty"`
	Heading              int                `json:"heading,omitempty"`
	ProximityAlertRadius int                `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          models.ReplyMarkup `json:"reply_markup,omitempty"`
}

// EditMessageLiveLocation https://core.telegram.org/bots/api#editmessagelivelocation
func EditMessageLiveLocation(ctx context.Context, b *bot.Bot, params *EditMessageLiveLocationParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "editMessageLiveLocation", params, result)

	return result, err
}
