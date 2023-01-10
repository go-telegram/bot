package models

// ForumTopic https://core.telegram.org/bots/api#forumtopic
type ForumTopic struct {
	MessageThreadID   int    `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicCreated https://core.telegram.org/bots/api#forumtopiccreated
type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicClosed https://core.telegram.org/bots/api#forumtopicclosed
type ForumTopicClosed struct {
}

// ForumTopicEdited https://core.telegram.org/bots/api#forumtopicedited
type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicReopened https://core.telegram.org/bots/api#forumtopicreopened
type ForumTopicReopened struct {
}

// GeneralForumTopicHidden https://core.telegram.org/bots/api#generalforumtopichidden
type GeneralForumTopicHidden struct {
}

// GeneralForumTopicUnhidden https://core.telegram.org/bots/api#generalforumtopicunhidden
type GeneralForumTopicUnhidden struct {
}

// WriteAccessAllowed https://core.telegram.org/bots/api#writeaccessallowed
type WriteAccessAllowed struct {
}
