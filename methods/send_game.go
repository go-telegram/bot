package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type SendGameParams struct {
	ChatID                   any                `json:"chat_id"`
	GameShorName             string             `json:"game_short_name"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendGame https://core.telegram.org/bots/api#sendgame
func SendGame(ctx context.Context, b *bot.Bot, params *SendGameParams) (*models.Message, error) {
	result := &models.Message{}

	err := bot.RawRequest(ctx, b, "sendGame", params, result)

	return result, err
}
