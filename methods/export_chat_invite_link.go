package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type ExportChatInviteLinkParams struct {
	ChatID any `json:"chat_id"`
}

// ExportChatInviteLink https://core.telegram.org/bots/api#exportchatinvitelink
func ExportChatInviteLink(ctx context.Context, b *bot.Bot, params *ExportChatInviteLinkParams) (string, error) {
	var result string

	err := bot.RawRequest(ctx, b, "exportChatInviteLink", params, &result)

	return result, err
}
