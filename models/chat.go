package models

type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSupergroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`
	From       User            `json:"from"`
	Date       int             `json:"date"`
	Bio        string          `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

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

type Chat struct {
	ID                 int        `json:"id"`
	Type               ChatType   `json:"type"`
	Title              string     `json:"title,omitempty"`
	Username           string     `json:"username,omitempty"`
	FirstName          string     `json:"first_name,omitempty"`
	LastName           string     `json:"last_name,omitempty"`
	Photo              *ChatPhoto `json:"photo,omitempty"`
	Bio                string     `json:"bio"`
	HasPrivateForwards bool       `json:"has_private_forwards,omitempty"`
	Description        string     `json:"description,omitempty"`
	InviteLink         string     `json:"invite_link,omitempty"`
	PinnedMessage      *Message   `json:"pinned_message,omitempty"`
	//Permissions           *ChatPermissions `json:"permissions,omitempty"`
	SlowModeDelay         int    `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime int    `json:"message_auto_delete_time,omitempty"`
	HasProtectedContent   bool   `json:"has_protected_content,omitempty"`
	StickerSetName        string `json:"sticker_set_name,omitempty"`
	CanSetStickerSet      bool   `json:"can_set_sticker_set,omitempty"`
	LinkedChatID          int    `json:"linked_chat_id,omitempty"`
	//Location              *ChatLocation    `json:"location,omitempty"`
}
