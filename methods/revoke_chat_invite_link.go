package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type RevokeChatInviteLinkParams struct {
	ChatID     any    `json:"chat_id" rules:"required,chat_id"`
	InviteLink string `json:"invite_link" rules:"required"`
}

// RevokeChatInviteLink https://core.telegram.org/bots/api#revokechatinvitelink
func RevokeChatInviteLink(ctx context.Context, b *bot.Bot, params *RevokeChatInviteLinkParams) (*models.ChatInviteLink, error) {
	result := &models.ChatInviteLink{}

	err := bot.RawRequest(ctx, b, "revokeChatInviteLink", params, &result)

	return result, err
}
