package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CopyMessageParams struct {
	ChatID                   any                    `json:"chat_id"`
	FromChatID               string                 `json:"from_chat_id"`
	MessageID                int                    `json:"message_id"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

// CopyMessage https://core.telegram.org/bots/api#copymessage
func CopyMessage(ctx context.Context, b *bot.Bot, params *CopyMessageParams) (*models.MessageID, error) {
	result := &models.MessageID{}

	err := bot.RawRequest(ctx, b, "forwardMessage", params, result)

	return result, err
}
