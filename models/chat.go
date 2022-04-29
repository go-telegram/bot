package models

// ChatJoinRequest https://core.telegram.org/bots/api#chatjoinrequest
type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`
	From       User            `json:"from"`
	Date       int             `json:"date"`
	Bio        string          `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

// ChatPhoto https://core.telegram.org/bots/api#chatphoto
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

// ChatInviteLink https://core.telegram.org/bots/api#chatinvitelink
type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`
	Creator                 User   `json:"creator"`
	CreatesJoinRequest      bool   `json:"creates_join_request"`
	IsPrimary               bool   `json:"is_primary"`
	IsRevoked               bool   `json:"is_revoked"`
	Name                    string `json:"name,omitempty"`
	ExpireDate              int    `json:"expire_date,omitempty"`
	MemberLimit             int    `json:"member_limit,omitempty"`
	PendingJoinRequestCount int    `json:"pending_join_request_count,omitempty"`
}

// ChatAdministratorRights https://core.telegram.org/bots/api#chatadministratorrights
type ChatAdministratorRights struct {
	IsAnonymous         bool `json:"is_anonymous"`
	CanManageChat       bool `json:"can_manage_chat"`
	CanDeleteMessages   bool `json:"can_delete_messages"`
	CanManageVideoChats bool `json:"can_manage_video_chats"`
	CanRestrictMembers  bool `json:"can_restrict_members"`
	CanPromoteMembers   bool `json:"can_promote_members"`
	CanChangeInfo       bool `json:"can_change_info"`
	CanInviteUsers      bool `json:"can_invite_users"`
	CanPostMessages     bool `json:"can_post_messages,omitempty"`
	CanEditMessages     bool `json:"can_edit_messages,omitempty"`
	CanPinMessages      bool `json:"can_pin_messages,omitempty"`
}

// ChatPermissions https://core.telegram.org/bots/api#chatpermissions
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
}

// ChatLocation https://core.telegram.org/bots/api#chatlocation
type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}

// Chat https://core.telegram.org/bots/api#chat
type Chat struct {
	ID                    int              `json:"id"`
	Type                  string           `json:"type"`
	Title                 string           `json:"title,omitempty"`
	Username              string           `json:"username,omitempty"`
	FirstName             string           `json:"first_name,omitempty"`
	LastName              string           `json:"last_name,omitempty"`
	Photo                 *ChatPhoto       `json:"photo,omitempty"`
	Bio                   string           `json:"bio"`
	HasPrivateForwards    bool             `json:"has_private_forwards,omitempty"`
	Description           string           `json:"description,omitempty"`
	InviteLink            string           `json:"invite_link,omitempty"`
	PinnedMessage         *Message         `json:"pinned_message,omitempty"`
	Permissions           *ChatPermissions `json:"permissions,omitempty"`
	SlowModeDelay         int              `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime int              `json:"message_auto_delete_time,omitempty"`
	HasProtectedContent   bool             `json:"has_protected_content,omitempty"`
	StickerSetName        string           `json:"sticker_set_name,omitempty"`
	CanSetStickerSet      bool             `json:"can_set_sticker_set,omitempty"`
	LinkedChatID          int              `json:"linked_chat_id,omitempty"`
	Location              *ChatLocation    `json:"location,omitempty"`
}
