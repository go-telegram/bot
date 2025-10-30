package models

// See Telegram API docs: https://core.telegram.org/bots/api#forumtopic
type ForumTopic struct {
	MessageThreadID   int    `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#forumtopiccreated
type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#forumtopicclosed
type ForumTopicClosed struct {
}

// See Telegram API docs: https://core.telegram.org/bots/api#forumtopicedited
type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#forumtopicreopened
type ForumTopicReopened struct {
}

// See Telegram API docs: https://core.telegram.org/bots/api#generalforumtopichidden
type GeneralForumTopicHidden struct {
}

// See Telegram API docs: https://core.telegram.org/bots/api#generalforumtopicunhidden
type GeneralForumTopicUnhidden struct {
}

// See Telegram API docs: https://core.telegram.org/bots/api#usershared
type UserShared struct {
	RequestID int   `json:"request_id"`
	UserID    int64 `json:"user_id"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#shareduser
type SharedUser struct {
	UserID    int64       `json:"user_id"`
	FirstName string      `json:"first_name,omitempty"`
	LastName  string      `json:"last_name,omitempty"`
	Username  string      `json:"username,omitempty"`
	Photo     []PhotoSize `json:"photo,omitempty"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#usersshared
type UsersShared struct {
	RequestID int          `json:"request_id"`
	Users     []SharedUser `json:"users"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#chatshared
type ChatShared struct {
	RequestID int         `json:"request_id"`
	ChatID    int64       `json:"chat_id"`
	Title     string      `json:"title,omitempty"`
	Username  string      `json:"username,omitempty"`
	Photo     []PhotoSize `json:"photo,omitempty"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#writeaccessallowed
type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request,omitempty"`
	WebAppName         string `json:"web_app_name,omitempty"`
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"`
}
