package models

type MessageEntityType string

const (
	MessageEntityTypeMention              MessageEntityType = "mention"
	MessageEntityTypeHashtag              MessageEntityType = "hashtag"
	MessageEntityTypeCashtag              MessageEntityType = "cashtag"
	MessageEntityTypeBotCommand           MessageEntityType = "bot_command"
	MessageEntityTypeURL                  MessageEntityType = "url"
	MessageEntityTypeEmail                MessageEntityType = "email"
	MessageEntityTypePhoneNumber          MessageEntityType = "phone_number"
	MessageEntityTypeBold                 MessageEntityType = "bold"
	MessageEntityTypeItalic               MessageEntityType = "italic"
	MessageEntityTypeUnderline            MessageEntityType = "underline"
	MessageEntityTypeStrikethrough        MessageEntityType = "strikethrough"
	MessageEntityTypeSpoiler              MessageEntityType = "spoiler"
	MessageEntityTypeBlockquote           MessageEntityType = "blockquote"
	MessageEntityTypeExpandableBlockquote MessageEntityType = "expandable_blockquote"
	MessageEntityTypeCode                 MessageEntityType = "code"
	MessageEntityTypePre                  MessageEntityType = "pre"
	MessageEntityTypeTextLink             MessageEntityType = "text_link"
	MessageEntityTypeTextMention          MessageEntityType = "text_mention"
	MessageEntityTypeCustomEmoji          MessageEntityType = "custom_emoji"
)

// MessageEntity https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type          MessageEntityType `json:"type"`
	Offset        int               `json:"offset"`
	Length        int               `json:"length"`
	URL           string            `json:"url,omitempty"`
	User          *User             `json:"user,omitempty"`
	Language      string            `json:"language,omitempty"`
	CustomEmojiID string            `json:"custom_emoji_id"`
}
