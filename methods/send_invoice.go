package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SendInvoiceParams struct {
	ChatID                    any                   `json:"chat_id" rules:"required,chat_id"`
	Title                     string                `json:"title" rules:"required,min:1,max:32"`
	Description               string                `json:"description" rules:"required,min:1,max:255"`
	Payload                   string                `json:"payload" rules:"required,min:1,max:128"`
	ProviderToken             string                `json:"provider_token" rules:"required"`
	Currency                  string                `json:"currency"`
	Prices                    []models.LabeledPrice `json:"prices"`
	MaxTipAmount              int                   `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int                 `json:"suggested_tip_amounts,omitempty"`
	StartParameter            string                `json:"start_parameter,omitempty"`
	ProviderData              string                `json:"provider_data,omitempty"`
	PhotoURL                  string                `json:"photo_url,omitempty"`
	PhotoSize                 int                   `json:"photo_size,omitempty"`
	PhotoWidth                int                   `json:"photo_width,omitempty"`
	PhotoHeight               int                   `json:"photo_height,omitempty"`
	NeedName                  bool                  `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                  `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                  `json:"need_email,omitempty"`
	NeedShippingAddress       bool                  `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                  `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                  `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool                  `json:"is_flexible,omitempty"`
	DisableNotification       bool                  `json:"disable_notification,omitempty"`
	ProtectContent            bool                  `json:"protect_content,omitempty"`
	ReplyToMessageID          int                   `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply  bool                  `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup               models.ReplyMarkup    `json:"reply_markup,omitempty"`
}

// SendInvoice https://core.telegram.org/bots/api#sendinvoice
func SendInvoice(ctx context.Context, b *bot.Bot, params *SendInvoiceParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendInvoice", params, result)

	return result, err
}
