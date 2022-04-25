package models

type MessageEntityType string

const (
	MessageEntityTypeMention       MessageEntityType = "mention"
	MessageEntityTypeHashtag       MessageEntityType = "hashtag"
	MessageEntityTypeCashtag       MessageEntityType = "cashtag"
	MessageEntityTypeBotCommand    MessageEntityType = "bot_command"
	MessageEntityTypeURL           MessageEntityType = "url"
	MessageEntityTypeEmail         MessageEntityType = "email"
	MessageEntityTypePhoneNumber   MessageEntityType = "phone_number"
	MessageEntityTypeBold          MessageEntityType = "bold"
	MessageEntityTypeItalic        MessageEntityType = "italic"
	MessageEntityTypeUnderline     MessageEntityType = "underline"
	MessageEntityTypeStrikethrough MessageEntityType = "strikethrough"
	MessageEntityTypeCode          MessageEntityType = "code"
	MessageEntityTypePre           MessageEntityType = "pre"
	MessageEntityTypeTextLink      MessageEntityType = "text_link"
	MessageEntityTypeTextMention   MessageEntityType = "text_mention"
)

type MessageEntity struct {
	Type     MessageEntityType `json:"type"`
	Offset   int               `json:"offset"`
	Length   int               `json:"length"`
	URL      string            `json:"url,omitempty"`
	User     *User             `json:"user,omitempty"`
	Language string            `json:"language,omitempty"`
}

type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size"`
}

type MessageID struct {
	ID int `json:"message_id"`
}

type Message struct {
	ID                   int             `json:"message_id"`
	From                 *User           `json:"from,omitempty"`
	SenderChat           *Chat           `json:"sender_chat,omitempty"`
	Date                 int             `json:"date"`
	Chat                 Chat            `json:"chat"`
	ForwardFrom          *User           `json:"forward_from,omitempty"`
	ForwardFromChat      *Chat           `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID *int            `json:"forward_from_message_id,omitempty"`
	ForwardSignature     *string         `json:"forward_signature,omitempty"`
	ForwardSenderName    *string         `json:"forward_sender_name,omitempty"`
	ForwardDate          *int            `json:"forward_date,omitempty"`
	IsAutomaticForward   *bool           `json:"is_automatic_forward,omitempty"`
	ReplyToMessage       *Message        `json:"reply_to_message,omitempty"`
	ViaBot               *User           `json:"via_bot,omitempty"`
	EditDate             *int            `json:"edit_date,omitempty"`
	HasProtectedContent  *bool           `json:"has_protected_content,omitempty"`
	MediaGroupID         *string         `json:"media_group_id,omitempty"`
	AuthorSignature      *string         `json:"author_signature,omitempty"`
	Text                 string          `json:"text,omitempty"`
	Entities             []MessageEntity `json:"entities,omitempty"`
	Photo                []PhotoSize     `json:"photo,omitempty"`
	//Animation            *Animation      `json:"animation,omitempty"`
	//Audio *Audio `json:"audio,omitempty"`
	Document *Document `json:"document,omitempty"`
	//Sticker *Sticker `json:"sticker,omitempty"`
	//Video *Video `json:"video,omitempty"`
	//VideoNote *VideoNote `json:"video_note,omitempty"`
	//Voice *Voice `json:"voice,omitempty"`
	Caption         *string         `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	//Contact         *Contact        `json:"contact,omitempty"`
	//Dice *Dice `json:"dice,omitempty"`
	//Game *Game `json:"game,omitempty"`
	//Poll *Poll `json:"poll,omitempty"`
	//Venue *Venue `json:"venue,omitempty"`
	//Location *Location `json:"location,omitempty"`
	NewChatMembers     []User      `json:"new_chat_members,omitempty"`
	LeftChatMember     *User       `json:"left_chat_member,omitempty"`
	NewChatTitle       *string     `json:"new_chat_title,omitempty"`
	NewChatPhoto       []PhotoSize `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto    *bool       `json:"delete_chat_photo,omitempty"`
	GroupChatCreated   *bool       `json:"group_chat_created,omitempty"`
	ChannelChatCreated *bool       `json:"channel_chat_created,omitempty"`
	//MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatID   *int     `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID *int     `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage     *Message `json:"pinned_message,omitempty"`
	//Invoice           *Invoice `json:"invoice,omitempty"`
	//SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`
	ConnectedWebsite *string `json:"connected_website,omitempty"`
	//PassportData     *PassportData `json:"passport_data,omitempty"`
	//ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`
	//VoiceChatScheduled *VoiceChatScheduled `json:"voice_chat_scheduled,omitempty"`
	//VoiceChatStarted *VoiceChatStarted `json:"voice_chat_started,omitempty"`
	//VoiceChatEnded *VoiceChatEnded `json:"voice_chat_ended,omitempty"`
	//VoiceChatParticipantsInvited *VoiceChatParticipantsInvited `json:"voice_chat_participants_invited,omitempty"`
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}
