package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CreateChatInviteLinkParams struct {
	ChatID             any    `json:"chat_id" rules:"required,chat_id"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int    `json:"expire_date,omitempty"`
	MemberLimit        int    `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

// ExportChatInviteLink https://core.telegram.org/bots/api#createchatinvitelink
func CreateChatInviteLink(ctx context.Context, b *bot.Bot, params *CreateChatInviteLinkParams) (*models.ChatInviteLink, error) {
	result := &models.ChatInviteLink{}

	err := bot.RawRequest(ctx, b, "createChatInviteLink", params, &result)

	return result, err
}
