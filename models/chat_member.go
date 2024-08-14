package models

import (
	"encoding/json"
	"fmt"
)

// ChatMemberUpdated https://core.telegram.org/bots/api#chatmemberupdated
type ChatMemberUpdated struct {
	Chat                    Chat            `json:"chat"`
	From                    User            `json:"from"`
	Date                    int             `json:"date"`
	OldChatMember           ChatMember      `json:"old_chat_member"`
	NewChatMember           ChatMember      `json:"new_chat_member"`
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`
	ViaJoinRequest          bool            `json:"via_join_request,omitempty"`
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link,omitempty"`
}

type ChatMemberType string

const (
	ChatMemberTypeOwner         ChatMemberType = "creator"
	ChatMemberTypeAdministrator ChatMemberType = "administrator"
	ChatMemberTypeMember        ChatMemberType = "member"
	ChatMemberTypeRestricted    ChatMemberType = "restricted"
	ChatMemberTypeLeft          ChatMemberType = "left"
	ChatMemberTypeBanned        ChatMemberType = "kicked"
)

// ChatMember https://core.telegram.org/bots/api#chatmember
type ChatMember struct {
	Type ChatMemberType

	Owner         *ChatMemberOwner
	Administrator *ChatMemberAdministrator
	Member        *ChatMemberMember
	Restricted    *ChatMemberRestricted
	Left          *ChatMemberLeft
	Banned        *ChatMemberBanned
}

func (c *ChatMember) UnmarshalJSON(data []byte) error {
	v := struct {
		Status ChatMemberType `json:"status"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v.Status {
	case ChatMemberTypeOwner:
		c.Type = ChatMemberTypeOwner
		c.Owner = &ChatMemberOwner{}
		return json.Unmarshal(data, c.Owner)
	case ChatMemberTypeAdministrator:
		c.Type = ChatMemberTypeAdministrator
		c.Administrator = &ChatMemberAdministrator{}
		return json.Unmarshal(data, c.Administrator)
	case ChatMemberTypeMember:
		c.Type = ChatMemberTypeMember
		c.Member = &ChatMemberMember{}
		return json.Unmarshal(data, c.Member)
	case ChatMemberTypeRestricted:
		c.Type = ChatMemberTypeRestricted
		c.Restricted = &ChatMemberRestricted{}
		return json.Unmarshal(data, c.Restricted)
	case ChatMemberTypeLeft:
		c.Type = ChatMemberTypeLeft
		c.Left = &ChatMemberLeft{}
		return json.Unmarshal(data, c.Left)
	case ChatMemberTypeBanned:
		c.Type = ChatMemberTypeBanned
		c.Banned = &ChatMemberBanned{}
		return json.Unmarshal(data, c.Banned)
	}

	return fmt.Errorf("unsupported ChatMember type")
}

func (c *ChatMember) MarshalJSON() ([]byte, error) {
	switch c.Type {
	case ChatMemberTypeOwner:
		c.Owner.Status = ChatMemberTypeOwner
		return json.Marshal(c.Owner)
	case ChatMemberTypeAdministrator:
		c.Administrator.Status = ChatMemberTypeAdministrator
		return json.Marshal(c.Administrator)
	case ChatMemberTypeMember:
		c.Member.Status = ChatMemberTypeMember
		return json.Marshal(c.Member)
	case ChatMemberTypeRestricted:
		c.Restricted.Status = ChatMemberTypeRestricted
		return json.Marshal(c.Restricted)
	case ChatMemberTypeLeft:
		c.Left.Status = ChatMemberTypeLeft
		return json.Marshal(c.Left)
	case ChatMemberTypeBanned:
		c.Banned.Status = ChatMemberTypeBanned
		return json.Marshal(c.Banned)
	}

	return nil, fmt.Errorf("unsupported ChatMember type")
}

// ChatMemberOwner https://core.telegram.org/bots/api#chatmemberowner
type ChatMemberOwner struct {
	Status      ChatMemberType `json:"status"` // The member's status in the chat, always “creator”
	User        *User          `json:"user"`
	IsAnonymous bool           `json:"is_anonymous"`
	CustomTitle string         `json:"custom_title,omitempty"`
}

// ChatMemberAdministrator https://core.telegram.org/bots/api#chatmemberadministrator
type ChatMemberAdministrator struct {
	Status              ChatMemberType `json:"status"` // The member's status in the chat, always “administrator”
	User                User           `json:"user"`
	CanBeEdited         bool           `json:"can_be_edited"`
	IsAnonymous         bool           `json:"is_anonymous"`
	CanManageChat       bool           `json:"can_manage_chat"`
	CanDeleteMessages   bool           `json:"can_delete_messages"`
	CanManageVideoChats bool           `json:"can_manage_video_chats"`
	CanRestrictMembers  bool           `json:"can_restrict_members"`
	CanPromoteMembers   bool           `json:"can_promote_members"`
	CanChangeInfo       bool           `json:"can_change_info"`
	CanInviteUsers      bool           `json:"can_invite_users"`
	CanPostMessages     bool           `json:"can_post_messages,omitempty"`
	CanEditMessages     bool           `json:"can_edit_messages,omitempty"`
	CanPinMessages      bool           `json:"can_pin_messages,omitempty"`
	CanPostStories      bool           `json:"can_post_stories,omitempty"`
	CanEditStories      bool           `json:"can_edit_stories,omitempty"`
	CanDeleteStories    bool           `json:"can_delete_stories,omitempty"`
	CanManageTopics     bool           `json:"can_manage_topics,omitempty"`
	CustomTitle         string         `json:"custom_title,omitempty"`
}

// ChatMemberMember https://core.telegram.org/bots/api#chatmembermember
type ChatMemberMember struct {
	Status    ChatMemberType `json:"status"` // The member's status in the chat, always “member”
	User      *User          `json:"user"`
	UntilDate int            `json:"until_date,omitempty"`
}

// ChatMemberRestricted https://core.telegram.org/bots/api#chatmemberrestricted
type ChatMemberRestricted struct {
	Status                ChatMemberType `json:"status"` // The member's status in the chat, always “restricted”
	User                  *User          `json:"user"`
	IsMember              bool           `json:"is_member"`
	CanSendMessages       bool           `json:"can_send_messages"`
	CanSendAudios         bool           `json:"can_send_audios"`
	CanSendDocuments      bool           `json:"can_send_documents"`
	CanSendPhotos         bool           `json:"can_send_photos"`
	CanSendVideos         bool           `json:"can_send_videos"`
	CanSendVideoNotes     bool           `json:"can_send_video_notes"`
	CanSendVoiceNotes     bool           `json:"can_send_voice_notes"`
	CanSendPolls          bool           `json:"can_send_polls"`
	CanSendOtherMessages  bool           `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool           `json:"can_add_web_page_previews"`
	CanChangeInfo         bool           `json:"can_change_info"`
	CanInviteUsers        bool           `json:"can_invite_users"`
	CanPinMessages        bool           `json:"can_pin_messages"`
	CanManageTopics       bool           `json:"can_manage_topics,omitempty"`
	UntilDate             int            `json:"until_date"`
}

// ChatMemberLeft https://core.telegram.org/bots/api#chatmemberleft
type ChatMemberLeft struct {
	Status ChatMemberType `json:"status"` // The member's status in the chat, always “left”
	User   *User          `json:"user"`
}

// ChatMemberBanned https://core.telegram.org/bots/api#chatmemberbanned
type ChatMemberBanned struct {
	Status    ChatMemberType `json:"status"` // The member's status in the chat, always “kicked”
	User      *User          `json:"user"`
	UntilDate int            `json:"until_date"`
}
