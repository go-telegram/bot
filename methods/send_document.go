package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SendDocumentParams struct {
	ChatID                      any                    `json:"chat_id" rules:"required,type:string|int"`
	Document                    string                 `json:"document" rules:"required"`
	Thumb                       string                 `json:"thumb,omitempty"`
	Caption                     string                 `json:"caption,omitempty"`
	ParseMode                   models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities             []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                   `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool                   `json:"disable_notification,omitempty"`
	ProtectContent              bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID            int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply    bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup                 models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

// SendDocument https://core.telegram.org/bots/api#senddocument
func SendDocument(ctx context.Context, b *bot.Bot, params *SendDocumentParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendDocument", params, result)

	return result, err
}
