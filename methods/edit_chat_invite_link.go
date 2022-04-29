package methods

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type EditChatInviteLinkParams struct {
	ChatID             any    `json:"chat_id" rules:"required,type:string|int"`
	InviteLink         string `json:"invite_link" rules:"required"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int    `json:"expire_date,omitempty"`
	MemberLimit        int    `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

// EditChatInviteLink https://core.telegram.org/bots/api#editchatinvitelink
func EditChatInviteLink(ctx context.Context, b *bot.Bot, params *EditChatInviteLinkParams) (*models.ChatInviteLink, error) {
	result := &models.ChatInviteLink{}

	err := bot.RawRequest(ctx, b, "editChatInviteLink", params, &result)

	return result, err
}
