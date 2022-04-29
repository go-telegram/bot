package methods

import (
	"context"

	"github.com/go-telegram/bot"
)

type PromoteChatMemberParams struct {
	ChatID             any  `json:"chat_id" rules:"required,type:string|int"`
	UserID             int  `json:"user_id" rules:"required"`
	IsAnonymous        bool `json:"is_anonymous,omitempty"`
	CanManageChat      bool `json:"can_manage_chat,omitempty"`
	CanPostMessages    bool `json:"can_post_messages,omitempty"`
	CanEditMessages    bool `json:"can_edit_messages,omitempty"`
	CanDeleteMessages  bool `json:"can_delete_messages,omitempty"`
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
	CanPromoteMembers  bool `json:"can_promote_members,omitempty"`
	CanChangeInfo      bool `json:"can_change_info,omitempty"`
	CanInviteUsers     bool `json:"can_invite_users,omitempty"`
	CanPinMessages     bool `json:"can_pin_messages,omitempty"`
}

// PromoteChatMember https://core.telegram.org/bots/api#promotechatmember
func PromoteChatMember(ctx context.Context, b *bot.Bot, params *PromoteChatMemberParams) (bool, error) {
	var result bool

	err := bot.RawRequest(ctx, b, "promoteChatMember", params, &result)

	return result, err
}
