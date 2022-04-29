package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SendVenueParams struct {
	ChatID                   any                `json:"chat_id" rules:"required,chat_id"`
	Latitude                 float64            `json:"latitude" rules:"required"`
	Longitude                float64            `json:"longitude" rules:"required"`
	Title                    string             `json:"title" rules:"required"`
	Address                  string             `json:"address" rules:"required"`
	FoursquareID             string             `json:"foursquare_id,omitempty"`
	FoursquareType           string             `json:"foursquare_type,omitempty"`
	GooglePlaceID            string             `json:"google_place_id,omitempty"`
	GooglePlaceType          string             `json:"google_place_type,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendVenue https://core.telegram.org/bots/api#sendvenue
func SendVenue(ctx context.Context, b *bot.Bot, params *SendVenueParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendVenue", params, result)

	return result, err
}
