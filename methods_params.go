package bot

import (
	"github.com/go-telegram/bot/models"
)

type SetWebhookParams struct {
	URL                string           `json:"url"`
	Certificate        models.InputFile `json:"certificate,omitempty"`
	IPAddress          string           `json:"ip_address,omitempty"`
	MaxConnections     int              `json:"max_connections,omitempty"`
	AllowedUpdates     []string         `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool             `json:"drop_pending_updates,omitempty"`
	SecretToken        string           `json:"secret_token,omitempty"`
}

type DeleteWebhookParams struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

// SendMessageParams https://core.telegram.org/bots/api#sendmessage
type SendMessageParams struct {
	BusinessConnectionID string                     `json:"business_connection_id,omitempty"`
	ChatID               any                        `json:"chat_id"`
	MessageThreadID      int                        `json:"message_thread_id,omitempty"`
	Text                 string                     `json:"text"`
	ParseMode            models.ParseMode           `json:"parse_mode,omitempty"`
	Entities             []models.MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions   *models.LinkPreviewOptions `json:"link_preview_options,omitempty"`
	DisableNotification  bool                       `json:"disable_notification,omitempty"`
	ProtectContent       bool                       `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                       `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string                     `json:"message_effect_id,omitempty"`
	ReplyParameters      *models.ReplyParameters    `json:"reply_parameters,omitempty"`
	ReplyMarkup          models.ReplyMarkup         `json:"reply_markup,omitempty"`
}

// ForwardMessageParams https://core.telegram.org/bots/api#forwardmessage
type ForwardMessageParams struct {
	ChatID              any  `json:"chat_id"`
	MessageThreadID     int  `json:"message_thread_id,omitempty"`
	FromChatID          any  `json:"from_chat_id"`
	VideoStartTimestamp int  `json:"video_start_timestamp,omitempty"`
	DisableNotification bool `json:"disable_notification,omitempty"`
	ProtectContent      bool `json:"protect_content,omitempty"`
	MessageID           int  `json:"message_id"`
}

// ForwardMessagesParams https://core.telegram.org/bots/api#forwardmessages
type ForwardMessagesParams struct {
	ChatID              any   `json:"chat_id"`
	MessageThreadID     int   `json:"message_thread_id,omitempty"`
	FromChatID          any   `json:"from_chat_id"`
	MessageIDs          []int `json:"message_ids"`
	DisableNotification bool  `json:"disable_notification,omitempty"`
	ProtectContent      bool  `json:"protect_content,omitempty"`
}

// CopyMessageParams https://core.telegram.org/bots/api#copymessage
type CopyMessageParams struct {
	ChatID                any                     `json:"chat_id"`
	MessageThreadID       int                     `json:"message_thread_id,omitempty"`
	FromChatID            any                     `json:"from_chat_id"`
	MessageID             int                     `json:"message_id"`
	VideoStartTimestamp   int                     `json:"video_start_timestamp,omitempty"`
	Caption               string                  `json:"caption,omitempty"`
	ParseMode             models.ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities       []models.MessageEntity  `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                    `json:"show_caption_above_media,omitempty"`
	DisableNotification   bool                    `json:"disable_notification,omitempty"`
	ProtectContent        bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast    bool                    `json:"allow_paid_broadcast,omitempty"`
	ReplyParameters       *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup           models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

// CopyMessagesParams https://core.telegram.org/bots/api#copymessages
type CopyMessagesParams struct {
	ChatID              any   `json:"chat_id"`
	MessageThreadID     int   `json:"message_thread_id,omitempty"`
	FromChatID          any   `json:"from_chat_id"`
	MessageIDs          []int `json:"message_ids"`
	DisableNotification bool  `json:"disable_notification,omitempty"`
	ProtectContent      bool  `json:"protect_content,omitempty"`
	RemoveCaption       bool  `json:"remove_caption,omitempty"`
}

// SendPhotoParams https://core.telegram.org/bots/api#sendphoto
type SendPhotoParams struct {
	BusinessConnectionID  string                  `json:"business_connection_id,omitempty"`
	ChatID                any                     `json:"chat_id"`
	MessageThreadID       int                     `json:"message_thread_id,omitempty"`
	Photo                 models.InputFile        `json:"photo"`
	Caption               string                  `json:"caption,omitempty"`
	ParseMode             models.ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities       []models.MessageEntity  `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                    `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool                    `json:"has_spoiler,omitempty"`
	DisableNotification   bool                    `json:"disable_notification,omitempty"`
	ProtectContent        bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast    bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID       string                  `json:"message_effect_id,omitempty"`
	ReplyParameters       *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup           models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

// SendAudioParams https://core.telegram.org/bots/api#sendaudio
type SendAudioParams struct {
	BusinessConnectionID string                  `json:"business_connection_id,omitempty"`
	ChatID               any                     `json:"chat_id"`
	MessageThreadID      int                     `json:"message_thread_id,omitempty"`
	Audio                models.InputFile        `json:"audio"`
	Caption              string                  `json:"caption,omitempty"`
	ParseMode            models.ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities      []models.MessageEntity  `json:"caption_entities,omitempty"`
	Duration             int                     `json:"duration,omitempty"`
	Performer            string                  `json:"performer,omitempty"`
	Title                string                  `json:"title,omitempty"`
	Thumbnail            models.InputFile        `json:"thumbnail,omitempty"`
	DisableNotification  bool                    `json:"disable_notification,omitempty"`
	ProtectContent       bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string                  `json:"message_effect_id,omitempty"`
	ReplyParameters      *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

// SendDocumentParams https://core.telegram.org/bots/api#senddocument
type SendDocumentParams struct {
	BusinessConnectionID        string                  `json:"business_connection_id,omitempty"`
	ChatID                      any                     `json:"chat_id"`
	MessageThreadID             int                     `json:"message_thread_id,omitempty"`
	Document                    models.InputFile        `json:"document"`
	Thumbnail                   models.InputFile        `json:"thumbnail,omitempty"`
	Caption                     string                  `json:"caption,omitempty"`
	ParseMode                   models.ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities             []models.MessageEntity  `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                    `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool                    `json:"disable_notification,omitempty"`
	ProtectContent              bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast          bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID             string                  `json:"message_effect_id,omitempty"`
	ReplyParameters             *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup                 models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

// SendVideoParams https://core.telegram.org/bots/api#sendvideo
type SendVideoParams struct {
	BusinessConnectionID  string                  `json:"business_connection_id,omitempty"`
	ChatID                any                     `json:"chat_id"`
	MessageThreadID       int                     `json:"message_thread_id,omitempty"`
	Video                 models.InputFile        `json:"video"`
	Duration              int                     `json:"duration,omitempty"`
	Width                 int                     `json:"width,omitempty"`
	Height                int                     `json:"height,omitempty"`
	Thumbnail             models.InputFile        `json:"thumbnail,omitempty"`
	Cover                 models.InputFile        `json:"cover,omitempty"`
	StartTimestamp        int                     `json:"start_timestamp,omitempty"`
	Caption               string                  `json:"caption,omitempty"`
	ParseMode             models.ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities       []models.MessageEntity  `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                    `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool                    `json:"has_spoiler,omitempty"`
	SupportsStreaming     bool                    `json:"supports_streaming,omitempty"`
	DisableNotification   bool                    `json:"disable_notification,omitempty"`
	ProtectContent        bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast    bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID       string                  `json:"message_effect_id,omitempty"`
	ReplyParameters       *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup           models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

// SendAnimationParams https://core.telegram.org/bots/api#sendanimation
type SendAnimationParams struct {
	BusinessConnectionID  string                  `json:"business_connection_id,omitempty"`
	ChatID                any                     `json:"chat_id"`
	MessageThreadID       int                     `json:"message_thread_id,omitempty"`
	Animation             models.InputFile        `json:"animation"`
	Duration              int                     `json:"duration,omitempty"`
	Width                 int                     `json:"width,omitempty"`
	Height                int                     `json:"height,omitempty"`
	Thumbnail             models.InputFile        `json:"thumbnail,omitempty"`
	Caption               string                  `json:"caption,omitempty"`
	ParseMode             models.ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities       []models.MessageEntity  `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                    `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool                    `json:"has_spoiler,omitempty"`
	DisableNotification   bool                    `json:"disable_notification,omitempty"`
	ProtectContent        bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast    bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID       string                  `json:"message_effect_id,omitempty"`
	ReplyParameters       *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup           models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

// SendVoiceParams https://core.telegram.org/bots/api#sendvoice
type SendVoiceParams struct {
	BusinessConnectionID string                  `json:"business_connection_id,omitempty"`
	ChatID               any                     `json:"chat_id"`
	MessageThreadID      int                     `json:"message_thread_id,omitempty"`
	Voice                models.InputFile        `json:"voice"`
	Caption              string                  `json:"caption,omitempty"`
	ParseMode            models.ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities      []models.MessageEntity  `json:"caption_entities,omitempty"`
	Duration             int                     `json:"duration,omitempty"`
	DisableNotification  bool                    `json:"disable_notification,omitempty"`
	ProtectContent       bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string                  `json:"message_effect_id,omitempty"`
	ReplyParameters      *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

// SendVideoNoteParams https://core.telegram.org/bots/api#sendvideonote
type SendVideoNoteParams struct {
	BusinessConnectionID string                  `json:"business_connection_id,omitempty"`
	ChatID               any                     `json:"chat_id"`
	MessageThreadID      int                     `json:"message_thread_id,omitempty"`
	VideoNote            models.InputFile        `json:"video_note"`
	Duration             int                     `json:"duration,omitempty"`
	Length               int                     `json:"length,omitempty"`
	Thumbnail            models.InputFile        `json:"thumbnail,omitempty"`
	DisableNotification  bool                    `json:"disable_notification,omitempty"`
	ProtectContent       bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string                  `json:"message_effect_id,omitempty"`
	ReplyParameters      *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

// SendPaidMediaParams https://core.telegram.org/bots/api#sendpaidmedia
type SendPaidMediaParams struct {
	BusinessConnectionID  string                  `json:"business_connection_id,omitempty"`
	ChatID                any                     `json:"chat_id"`
	StarCount             int                     `json:"star_count"`
	Media                 []models.InputPaidMedia `json:"media"`
	Payload               string                  `json:"payload,omitempty"`
	Caption               string                  `json:"caption,omitempty"`
	ParseMode             models.ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities       []models.MessageEntity  `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                    `json:"show_caption_above_media,omitempty"`
	DisableNotification   bool                    `json:"disable_notification,omitempty"`
	ProtectContent        bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast    bool                    `json:"allow_paid_broadcast,omitempty"`
	ReplyParameters       *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup           models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

// SendMediaGroupParams https://core.telegram.org/bots/api#sendmediagroup
type SendMediaGroupParams struct {
	BusinessConnectionID string                  `json:"business_connection_id,omitempty"`
	ChatID               any                     `json:"chat_id"`
	MessageThreadID      int                     `json:"message_thread_id,omitempty"`
	Media                []models.InputMedia     `json:"media"`
	DisableNotification  bool                    `json:"disable_notification,omitempty"`
	ProtectContent       bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string                  `json:"message_effect_id,omitempty"`
	ReplyParameters      *models.ReplyParameters `json:"reply_parameters,omitempty"`
}

// SendLocationParams https://core.telegram.org/bots/api#sendlocation
type SendLocationParams struct {
	BusinessConnectionID string                  `json:"business_connection_id,omitempty"`
	ChatID               any                     `json:"chat_id"`
	MessageThreadID      int                     `json:"message_thread_id,omitempty"`
	Latitude             float64                 `json:"latitude"`
	Longitude            float64                 `json:"longitude"`
	HorizontalAccuracy   float64                 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int                     `json:"live_period,omitempty"`
	Heading              int                     `json:"heading,omitempty"`
	ProximityAlertRadius int                     `json:"proximity_alert_radius,omitempty"`
	DisableNotification  bool                    `json:"disable_notification,omitempty"`
	ProtectContent       bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string                  `json:"message_effect_id,omitempty"`
	ReplyParameters      *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

type EditMessageLiveLocationParams struct {
	BusinessConnectionID string             `json:"business_connection_id,omitempty"`
	ChatID               any                `json:"chat_id,omitempty"`
	MessageID            int                `json:"message_id,omitempty"`
	InlineMessageID      string             `json:"inline_message_id,omitempty"`
	Latitude             float64            `json:"latitude"`
	Longitude            float64            `json:"longitude"`
	LivePeriod           int                `json:"live_period,omitempty"`
	HorizontalAccuracy   float64            `json:"horizontal_accuracy,omitempty"`
	Heading              int                `json:"heading,omitempty"`
	ProximityAlertRadius int                `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type StopMessageLiveLocationParams struct {
	BusinessConnectionID string             `json:"business_connection_id,omitempty"`
	ChatID               any                `json:"chat_id,omitempty"`
	MessageID            int                `json:"message_id,omitempty"`
	InlineMessageID      string             `json:"inline_message_id,omitempty"`
	ReplyMarkup          models.ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendVenueParams https://core.telegram.org/bots/api#sendvenue
type SendVenueParams struct {
	BusinessConnectionID string                  `json:"business_connection_id,omitempty"`
	ChatID               any                     `json:"chat_id"`
	MessageThreadID      int                     `json:"message_thread_id,omitempty"`
	Latitude             float64                 `json:"latitude"`
	Longitude            float64                 `json:"longitude"`
	Title                string                  `json:"title"`
	Address              string                  `json:"address"`
	FoursquareID         string                  `json:"foursquare_id,omitempty"`
	FoursquareType       string                  `json:"foursquare_type,omitempty"`
	GooglePlaceID        string                  `json:"google_place_id,omitempty"`
	GooglePlaceType      string                  `json:"google_place_type,omitempty"`
	DisableNotification  bool                    `json:"disable_notification,omitempty"`
	ProtectContent       bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string                  `json:"message_effect_id,omitempty"`
	ReplyParameters      *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

// SendContactParams https://core.telegram.org/bots/api#sendcontact
type SendContactParams struct {
	BusinessConnectionID string                  `json:"business_connection_id,omitempty"`
	ChatID               any                     `json:"chat_id"`
	MessageThreadID      int                     `json:"message_thread_id,omitempty"`
	PhoneNumber          string                  `json:"phone_number"`
	FirstName            string                  `json:"first_name"`
	LastName             string                  `json:"last_name,omitempty"`
	VCard                string                  `json:"vcard,omitempty"`
	DisableNotification  bool                    `json:"disable_notification,omitempty"`
	ProtectContent       bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string                  `json:"message_effect_id,omitempty"`
	ReplyParameters      *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

// SendPollParams https://core.telegram.org/bots/api#sendpoll
type SendPollParams struct {
	BusinessConnectionID  string                   `json:"business_connection_id,omitempty"`
	ChatID                any                      `json:"chat_id"`
	MessageThreadID       int                      `json:"message_thread_id,omitempty"`
	Question              string                   `json:"question"`
	QuestionParseMode     models.ParseMode         `json:"question_parse_mode,omitempty"`
	QuestionEntities      []models.MessageEntity   `json:"question_entities,omitempty"`
	Options               []models.InputPollOption `json:"options"`
	IsAnonymous           *bool                    `json:"is_anonymous,omitempty"`
	Type                  string                   `json:"type,omitempty"`
	AllowsMultipleAnswers bool                     `json:"allows_multiple_answers,omitempty"`
	CorrectOptionID       int                      `json:"correct_option_id"`
	Explanation           string                   `json:"explanation,omitempty"`
	ExplanationParseMode  string                   `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities   []models.MessageEntity   `json:"explanation_entities,omitempty"`
	OpenPeriod            int                      `json:"open_period,omitempty"`
	CloseDate             int                      `json:"close_date,omitempty"`
	IsClosed              bool                     `json:"is_closed,omitempty"`
	DisableNotification   bool                     `json:"disable_notification,omitempty"`
	ProtectContent        bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast    bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID       string                   `json:"message_effect_id,omitempty"`
	ReplyParameters       *models.ReplyParameters  `json:"reply_parameters,omitempty"`
	ReplyMarkup           models.ReplyMarkup       `json:"reply_markup,omitempty"`
}

// SendDiceParams https://core.telegram.org/bots/api#senddice
type SendDiceParams struct {
	BusinessConnectionID string                  `json:"business_connection_id,omitempty"`
	ChatID               any                     `json:"chat_id"`
	MessageThreadID      int                     `json:"message_thread_id,omitempty"`
	Emoji                string                  `json:"emoji,omitempty"`
	DisableNotification  bool                    `json:"disable_notification,omitempty"`
	ProtectContent       bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string                  `json:"message_effect_id,omitempty"`
	ReplyParameters      *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

type SendChatActionParams struct {
	BusinessConnectionID string            `json:"business_connection_id,omitempty"`
	ChatID               any               `json:"chat_id"`
	MessageThreadID      int               `json:"message_thread_id,omitempty"`
	Action               models.ChatAction `json:"action"`
}

// SetMessageReactionParams https://core.telegram.org/bots/api#setmessagereaction
type SetMessageReactionParams struct {
	ChatID    any                   `json:"chat_id"`
	MessageID int                   `json:"message_id"`
	Reaction  []models.ReactionType `json:"reaction,omitempty"`
	IsBig     *bool                 `json:"is_big,omitempty"`
}

type GetUserProfilePhotosParams struct {
	UserID int64 `json:"user_id"`
	Offset int   `json:"offset,omitempty"`
	Limit  int   `json:"limit,omitempty"`
}

// SetUserEmojiStatusParams https://core.telegram.org/bots/api#setuseremojistatus
type SetUserEmojiStatusParams struct {
	UserID                    int64  `json:"user_id"`
	EmojiStatusCustomEmojiID  string `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate int    `json:"emoji_status_expiration_date,omitempty"`
}

type GetFileParams struct {
	FileID string `json:"file_id"`
}

type BanChatMemberParams struct {
	ChatID         any   `json:"chat_id"`
	UserID         int64 `json:"user_id"`
	UntilDate      int   `json:"until_date,omitempty"`
	RevokeMessages bool  `json:"revoke_messages,omitempty"`
}

type UnbanChatMemberParams struct {
	ChatID       any   `json:"chat_id"`
	UserID       int64 `json:"user_id"`
	OnlyIfBanned bool  `json:"only_if_banned,omitempty"`
}

type RestrictChatMemberParams struct {
	ChatID                        any                     `json:"chat_id"`
	UserID                        int64                   `json:"user_id"`
	Permissions                   *models.ChatPermissions `json:"permissions,omitempty"`
	UseIndependentChatPermissions bool                    `json:"use_independent_chat_permissions,omitempty"`
	UntilDate                     int                     `json:"until_date,omitempty"`
}

type PromoteChatMemberParams struct {
	ChatID              any   `json:"chat_id" rules:"required,chat_id"`
	UserID              int64 `json:"user_id" rules:"required"`
	IsAnonymous         bool  `json:"is_anonymous,omitempty"`
	CanManageChat       bool  `json:"can_manage_chat,omitempty"`
	CanDeleteMessages   bool  `json:"can_delete_messages,omitempty"`
	CanManageVideoChats bool  `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers  bool  `json:"can_restrict_members,omitempty"`
	CanPromoteMembers   bool  `json:"can_promote_members,omitempty"`
	CanChangeInfo       bool  `json:"can_change_info,omitempty"`
	CanInviteUsers      bool  `json:"can_invite_users,omitempty"`
	CanPostMessages     bool  `json:"can_post_messages,omitempty"`
	CanEditMessages     bool  `json:"can_edit_messages,omitempty"`
	CanPinMessages      bool  `json:"can_pin_messages,omitempty"`
	CanPostStories      bool  `json:"can_post_stories,omitempty"`
	CanEditStories      bool  `json:"can_edit_stories,omitempty"`
	CanDeleteStories    bool  `json:"can_delete_stories,omitempty"`
	CanManageTopics     bool  `json:"can_manage_topics,omitempty"`
}

type SetChatAdministratorCustomTitleParams struct {
	ChatID      any    `json:"chat_id"`
	UserID      int64  `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

type BanChatSenderChatParams struct {
	ChatID       any `json:"chat_id"`
	SenderChatID int `json:"sender_chat_id"`
}

type UnbanChatSenderChatParams struct {
	ChatID       any `json:"chat_id"`
	SenderChatID int `json:"sender_chat_id"`
}

type SetChatPermissionsParams struct {
	ChatID                        any                    `json:"chat_id"`
	Permissions                   models.ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool                   `json:"use_independent_chat_permissions,omitempty"`
}

type ExportChatInviteLinkParams struct {
	ChatID any `json:"chat_id"`
}

type CreateChatInviteLinkParams struct {
	ChatID             any    `json:"chat_id"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int    `json:"expire_date,omitempty"`
	MemberLimit        int    `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

type EditChatInviteLinkParams struct {
	ChatID             any    `json:"chat_id"`
	InviteLink         string `json:"invite_link"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int    `json:"expire_date,omitempty"`
	MemberLimit        int    `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

type CreateChatSubscriptionInviteLinkParams struct {
	ChatID             any    `json:"chat_id"`
	Name               string `json:"name,omitempty"`
	SubscriptionPeriod int    `json:"subscription_period"`
	SubscriptionPrice  int    `json:"subscription_price"`
}

type EditChatSubscriptionInviteLinkParams struct {
	ChatID     any    `json:"chat_id"`
	InviteLink string `json:"invite_link"`
	Name       string `json:"name,omitempty"`
}

type RevokeChatInviteLinkParams struct {
	ChatID     any    `json:"chat_id"`
	InviteLink string `json:"invite_link"`
}

type ApproveChatJoinRequestParams struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

type DeclineChatJoinRequestParams struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

type SetChatPhotoParams struct {
	ChatID any              `json:"chat_id"`
	Photo  models.InputFile `json:"photo"`
}

type DeleteChatPhotoParams struct {
	ChatID any `json:"chat_id"`
}

type SetChatTitleParams struct {
	ChatID any    `json:"chat_id"`
	Title  string `json:"title"`
}

type SetChatDescriptionParams struct {
	ChatID      any    `json:"chat_id"`
	Description string `json:"description"`
}

type PinChatMessageParams struct {
	ChatID              any  `json:"chat_id"`
	MessageID           int  `json:"message_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
}

type UnpinChatMessageParams struct {
	ChatID    any `json:"chat_id"`
	MessageID int `json:"message_id,omitempty"`
}

type UnpinAllChatMessagesParams struct {
	ChatID any `json:"chat_id"`
}

type LeaveChatParams struct {
	ChatID any `json:"chat_id"`
}

type GetChatParams struct {
	ChatID any `json:"chat_id"`
}

type GetChatAdministratorsParams struct {
	ChatID any `json:"chat_id"`
}

type GetChatMemberCountParams struct {
	ChatID any `json:"chat_id"`
}

type GetChatMemberParams struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

type SetChatStickerSetParams struct {
	ChatID         any    `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

type CreateForumTopicParams struct {
	ChatID            any    `json:"chat_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

type EditForumTopicParams struct {
	ChatID            any    `json:"chat_id"`
	MessageThreadID   int    `json:"message_thread_id"`
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

type CloseForumTopicParams struct {
	ChatID          any `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

type ReopenForumTopicParams struct {
	ChatID          any `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

type DeleteForumTopicParams struct {
	ChatID          any `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

type UnpinAllForumTopicMessagesParams struct {
	ChatID          any `json:"chat_id"`
	MessageThreadID int `json:"message_thread_id"`
}

type EditGeneralForumTopicParams struct {
	ChatID any    `json:"chat_id"`
	Name   string `json:"name"`
}

type CloseGeneralForumTopicParams struct {
	ChatID any `json:"chat_id"`
}

type ReopenGeneralForumTopicParams struct {
	ChatID any `json:"chat_id"`
}

type HideGeneralForumTopicParams struct {
	ChatID any `json:"chat_id"`
}

type UnhideGeneralForumTopicParams struct {
	ChatID any `json:"chat_id"`
}

type UnpinAllGeneralForumTopicMessagesParams struct {
	ChatID any `json:"chat_id"`
}

type DeleteChatStickerSetParams struct {
	ChatID         any    `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

type AnswerCallbackQueryParams struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	URL             string `json:"url,omitempty"`
	CacheTime       int    `json:"cache_time,omitempty"`
}

// GetUserChatBoostsParams https://core.telegram.org/bots/api#getuserchatboosts
type GetUserChatBoostsParams struct {
	ChatID any `json:"chat_id"`
	UserID int `json:"user_id"`
}

// GetBusinessConnectionParams https://core.telegram.org/bots/api#getbusinessconnection
type GetBusinessConnectionParams struct {
	BusinessConnectionID any `json:"business_connection_id"`
}

type SetMyCommandsParams struct {
	Commands     []models.BotCommand    `json:"commands"`
	Scope        models.BotCommandScope `json:"scope,omitempty"`
	LanguageCode string                 `json:"language_code,omitempty"`
}

type DeleteMyCommandsParams struct {
	Scope        models.BotCommandScope `json:"scope,omitempty"`
	LanguageCode string                 `json:"language_code,omitempty"`
}

type GetMyCommandsParams struct {
	Scope        models.BotCommandScope `json:"scope,omitempty"`
	LanguageCode string                 `json:"language_code,omitempty"`
}

type SetMyNameParams struct {
	Name         string `json:"name,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

type GetMyNameParams struct {
	LanguageCode string `json:"language_code,omitempty"`
}

type SetMyDescriptionParams struct {
	Description  string `json:"description,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

type GetMyDescriptionParams struct {
	LanguageCode string `json:"language_code,omitempty"`
}

type SetMyShortDescriptionParams struct {
	ShortDescription string `json:"short_description,omitempty"`
	LanguageCode     string `json:"language_code,omitempty"`
}

type GetMyShortDescriptionParams struct {
	LanguageCode string `json:"language_code,omitempty"`
}

type SetChatMenuButtonParams struct {
	ChatID     any                    `json:"chat_id,omitempty"`
	MenuButton models.InputMenuButton `json:"menu_button"`
}

type GetChatMenuButtonParams struct {
	ChatID any `json:"chat_id"`
}

type SetMyDefaultAdministratorRightsParams struct {
	Rights      *models.ChatAdministratorRights `json:"rights,omitempty"`
	ForChannels bool                            `json:"for_channels,omitempty"`
}

type GetMyDefaultAdministratorRightsParams struct {
	ForChannels bool `json:"for_channels,omitempty"`
}

// EditMessageTextParams https://core.telegram.org/bots/api#editmessagetext
type EditMessageTextParams struct {
	BusinessConnectionID string                     `json:"business_connection_id,omitempty"`
	ChatID               any                        `json:"chat_id,omitempty"`
	MessageID            int                        `json:"message_id,omitempty"`
	InlineMessageID      string                     `json:"inline_message_id,omitempty"`
	Text                 string                     `json:"text"`
	ParseMode            models.ParseMode           `json:"parse_mode,omitempty"`
	Entities             []models.MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions   *models.LinkPreviewOptions `json:"link_preview_options,omitempty"`
	ReplyMarkup          models.ReplyMarkup         `json:"reply_markup,omitempty"`
}

type EditMessageCaptionParams struct {
	BusinessConnectionID  string                 `json:"business_connection_id,omitempty"`
	ChatID                any                    `json:"chat_id,omitempty"`
	MessageID             int                    `json:"message_id,omitempty"`
	InlineMessageID       string                 `json:"inline_message_id,omitempty"`
	Caption               string                 `json:"caption,omitempty"`
	ParseMode             models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []models.MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                   `json:"k,omitempty"`
	DisableWebPagePreview bool                   `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type EditMessageMediaParams struct {
	BusinessConnectionID string             `json:"business_connection_id,omitempty"`
	ChatID               any                `json:"chat_id,omitempty"`
	MessageID            int                `json:"message_id,omitempty"`
	InlineMessageID      string             `json:"inline_message_id,omitempty"`
	Media                models.InputMedia  `json:"media"`
	ReplyMarkup          models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type EditMessageReplyMarkupParams struct {
	BusinessConnectionID string             `json:"business_connection_id,omitempty"`
	ChatID               any                `json:"chat_id,omitempty"`
	MessageID            int                `json:"message_id,omitempty"`
	InlineMessageID      string             `json:"inline_message_id,omitempty"`
	ReplyMarkup          models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type StopPollParams struct {
	BusinessConnectionID string             `json:"business_connection_id,omitempty"`
	ChatID               any                `json:"chat_id"`
	MessageID            int                `json:"message_id"`
	ReplyMarkup          models.ReplyMarkup `json:"reply_markup,omitempty"`
}

// DeleteMessageParams https://core.telegram.org/bots/api#deletemessage
type DeleteMessageParams struct {
	ChatID    any `json:"chat_id"`
	MessageID int `json:"message_id"`
}

// DeleteMessagesParams https://core.telegram.org/bots/api#deletemessages
type DeleteMessagesParams struct {
	ChatID     any   `json:"chat_id"`
	MessageIDs []int `json:"message_ids"`
}

// SendStickerParams https://core.telegram.org/bots/api#sendsticker
type SendStickerParams struct {
	BusinessConnectionID string                  `json:"business_connection_id,omitempty"`
	ChatID               any                     `json:"chat_id"`
	MessageThreadID      int                     `json:"message_thread_id,omitempty"`
	Sticker              models.InputFile        `json:"sticker"`
	Emoji                string                  `json:"emoji,omitempty"`
	DisableNotification  bool                    `json:"disable_notification,omitempty"`
	ProtectContent       bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string                  `json:"message_effect_id,omitempty"`
	ReplyParameters      *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

type GetStickerSetParams struct {
	Name string `json:"name"`
}

type GetCustomEmojiStickersParams struct {
	CustomEmojiIDs []string `json:"custom_emoji_ids"`
}

type UploadStickerFileParams struct {
	UserID        int64            `json:"user_id"`
	Sticker       models.InputFile `json:"sticker"`
	StickerFormat string           `json:"sticker_format"`
}

type CreateNewStickerSetParams struct {
	UserID          int64                 `json:"user_id"`
	Name            string                `json:"name"`
	Title           string                `json:"title"`
	Stickers        []models.InputSticker `json:"stickers"`
	StickerType     string                `json:"sticker_type,omitempty"`
	NeedsRepainting bool                  `json:"needs_repainting,omitempty"`
}

type AddStickerToSetParams struct {
	UserID  int64               `json:"user_id"`
	Name    string              `json:"name"`
	Sticker models.InputSticker `json:"sticker"`
}

type SetStickerPositionInSetParams struct {
	Sticker  string `json:"sticker"`
	Position int    `json:"position"`
}

type DeleteStickerFromSetParams struct {
	Sticker string `json:"sticker"`
}

type ReplaceStickerInSetParams struct {
	UserID     int64               `json:"user_id"`
	Name       string              `json:"name"`
	OldSticker string              `json:"old_sticker"`
	Sticker    models.InputSticker `json:"sticker"`
}

type SetStickerEmojiListParams struct {
	Sticker   string   `json:"sticker"`
	EmojiList []string `json:"emoji_list"`
}

type SetStickerKeywordsParams struct {
	Sticker  string   `json:"sticker"`
	Keywords []string `json:"keywords"`
}

type SetStickerMaskPositionParams struct {
	Sticker      string              `json:"sticker"`
	MaskPosition models.MaskPosition `json:"mask_position"`
}

type SetStickerSetTitleParams struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type SetStickerSetThumbnailParams struct {
	Name      string           `json:"name"`
	UserID    int64            `json:"user_id"`
	Thumbnail models.InputFile `json:"thumbnail,omitempty"`
	Format    string           `json:"format"`
}

type SetCustomEmojiStickerSetThumbnailParams struct {
	Name          string `json:"name"`
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

type DeleteStickerSetParams struct {
	Name string `json:"name"`
}

type AnswerInlineQueryParams struct {
	InlineQueryID string                           `json:"inline_query_id"`
	Results       []models.InlineQueryResult       `json:"results"`
	CacheTime     int                              `json:"cache_time,omitempty"`
	IsPersonal    bool                             `json:"is_personal,omitempty"`
	NextOffset    string                           `json:"next_offset,omitempty"`
	Button        *models.InlineQueryResultsButton `json:"button,omitempty"`
}

type AnswerWebAppQueryParams struct {
	WebAppQueryID string                   `json:"web_app_query_id"`
	Result        models.InlineQueryResult `json:"result"`
}

// SavePreparedInlineMessageParams https://core.telegram.org/bots/api#savepreparedinlinemessage
type SavePreparedInlineMessageParams struct {
	UserID            int64                    `json:"user_id"`
	Result            models.InlineQueryResult `json:"result"`
	AllowUserChats    bool                     `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool                     `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool                     `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool                     `json:"allow_channel_chats,omitempty"`
}

// SendInvoiceParams https://core.telegram.org/bots/api#sendinvoice
type SendInvoiceParams struct {
	ChatID                    any                     `json:"chat_id"`
	MessageThreadID           int                     `json:"message_thread_id,omitempty"`
	Title                     string                  `json:"title"`
	Description               string                  `json:"description"`
	Payload                   string                  `json:"payload"`
	ProviderToken             string                  `json:"provider_token,omitempty"`
	Currency                  string                  `json:"currency"`
	Prices                    []models.LabeledPrice   `json:"prices"`
	MaxTipAmount              int                     `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int                   `json:"suggested_tip_amounts,omitempty"`
	StartParameter            string                  `json:"start_parameter,omitempty"`
	ProviderData              string                  `json:"provider_data,omitempty"`
	PhotoURL                  string                  `json:"photo_url,omitempty"`
	PhotoSize                 int                     `json:"photo_size,omitempty"`
	PhotoWidth                int                     `json:"photo_width,omitempty"`
	PhotoHeight               int                     `json:"photo_height,omitempty"`
	NeedName                  bool                    `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                    `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                    `json:"need_email,omitempty"`
	NeedShippingAddress       bool                    `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                    `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                    `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool                    `json:"is_flexible,omitempty"`
	DisableNotification       bool                    `json:"disable_notification,omitempty"`
	ProtectContent            bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast        bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID           string                  `json:"message_effect_id,omitempty"`
	ReplyParameters           *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup               models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

type CreateInvoiceLinkParams struct {
	BusinessConnectionID      string                `json:"business_connection_id,omitempty"`
	Title                     string                `json:"title"`
	Description               string                `json:"description"`
	Payload                   string                `json:"payload"`
	ProviderToken             string                `json:"provider_token,omitempty"`
	Currency                  string                `json:"currency"`
	Prices                    []models.LabeledPrice `json:"prices"`
	SubscriptionPeriod        int                   `json:"subscription_period,omitempty"`
	MaxTipAmount              int                   `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int                 `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string                `json:"provider_data,omitempty"`
	PhotoURL                  string                `json:"photo_url,omitempty"`
	PhotoSize                 int                   `json:"photo_size,omitempty"`
	PhotoWidth                int                   `json:"photo_width,omitempty"`
	PhotoHeight               int                   `json:"photo_height,omitempty"`
	NeedName                  bool                  `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                  `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                  `json:"need_email,omitempty"`
	NeedShippingAddress       bool                  `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                  `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                  `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool                  `json:"is_flexible,omitempty"`
}

type AnswerShippingQueryParams struct {
	ShippingQueryID string                  `json:"shipping_query_id"`
	OK              bool                    `json:"ok"`
	ShippingOptions []models.ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    string                  `json:"error_message,omitempty"`
}

type AnswerPreCheckoutQueryParams struct {
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`
	OK                 bool   `json:"ok"`
	ErrorMessage       string `json:"error_message,omitempty"`
}

type GetStarTransactionsParams struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

type RefundStarPaymentParams struct {
	UserID                  int64  `json:"user_id"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
}

// EditUserStarSubscriptionParams https://core.telegram.org/bots/api#edituserstarsubscription
type EditUserStarSubscriptionParams struct {
	UserID                  int64  `json:"user_id"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	IsCanceled              bool   `json:"is_canceled"`
}

type SetPassportDataErrorsParams struct {
	UserID int64                         `json:"user_id"`
	Errors []models.PassportElementError `json:"errors"`
}

// SendGameParams https://core.telegram.org/bots/api#sendgame
type SendGameParams struct {
	BusinessConnectionID string                  `json:"business_connection_id,omitempty"`
	ChatID               any                     `json:"chat_id"`
	MessageThreadID      int                     `json:"message_thread_id,omitempty"`
	GameShorName         string                  `json:"game_short_name"`
	DisableNotification  bool                    `json:"disable_notification,omitempty"`
	ProtectContent       bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      string                  `json:"message_effect_id,omitempty"`
	ReplyParameters      *models.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          models.ReplyMarkup      `json:"reply_markup,omitempty"`
}

type SetGameScoreParams struct {
	UserID             int64 `json:"user_id"`
	Score              int   `json:"score"`
	Force              bool  `json:"force,omitempty"`
	DisableEditMessage bool  `json:"disable_edit_message,omitempty"`
	ChatID             any   `json:"chat_id,omitempty"`
	MessageID          int   `json:"message_id,omitempty"`
	InlineMessageID    int   `json:"inline_message_id,omitempty"`
}

type GetGameHighScoresParams struct {
	UserID          int64 `json:"user_id"`
	ChatID          any   `json:"chat_id,omitempty"`
	MessageID       int   `json:"message_id,omitempty"`
	InlineMessageID int   `json:"inline_message_id,omitempty"`
}

// SendGiftParams https://core.telegram.org/bots/api#sendgift
type SendGiftParams struct {
	UserID        int64                  `json:"user_id"`
	ChatID        any                    `json:"chat_id,omitempty"`
	GiftID        string                 `json:"gift_id"`
	PayForUpgrade bool                   `json:"pay_for_upgrade,omitempty"`
	Text          string                 `json:"text,omitempty"`
	TextParseMode models.ParseMode       `json:"text_parse_mode,omitempty"`
	TextEntities  []models.MessageEntity `json:"text_entities,omitempty"`
}

// VerifyUserParams https://core.telegram.org/bots/api#verifyuser
type VerifyUserParams struct {
	UserID            int64  `json:"user_id"`
	CustomDescription string `json:"custom_description,omitempty"`
}

// VerifyChatParams https://core.telegram.org/bots/api#verifychat
type VerifyChatParams struct {
	ChatID            any    `json:"chat_id"`
	CustomDescription string `json:"custom_description,omitempty"`
}

// RemoveUserVerificationParams https://core.telegram.org/bots/api#removeuserverification
type RemoveUserVerificationParams struct {
	UserID int64 `json:"user_id"`
}

// RemoveChatVerificationParams https://core.telegram.org/bots/api#removechatverification
type RemoveChatVerificationParams struct {
	ChatID any `json:"chat_id"`
}

// ReadBusinessMessageParams https://core.telegram.org/bots/api#readbusinessmessage
type ReadBusinessMessageParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	ChatID               int    `json:"chat_id"`
	MessageID            int    `json:"message_id"`
}

// DeleteBusinessMessagesParams https://core.telegram.org/bots/api#deletebusinessmessages
type DeleteBusinessMessagesParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	MessageIDs           []int  `json:"message_ids"`
}

// SetBusinessAccountNameParams https://core.telegram.org/bots/api#setbusinessaccountname
type SetBusinessAccountNameParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name,omitempty"`
}

// SetBusinessAccountUsernameParams https://core.telegram.org/bots/api#setbusinessaccountusername
type SetBusinessAccountUsernameParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	Username             string `json:"username,omitempty"`
}

// SetBusinessAccountBioParams https://core.telegram.org/bots/api#setbusinessaccountbio
type SetBusinessAccountBioParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	Bio                  string `json:"bio,omitempty"`
}

// SetBusinessAccountProfilePhotoParams https://core.telegram.org/bots/api#setbusinessaccountprofilephoto
type SetBusinessAccountProfilePhotoParams struct {
	BusinessConnectionID string                   `json:"business_connection_id"`
	Photo                models.InputProfilePhoto `json:"photo"`
	IsPublic             bool                     `json:"is_public,omitempty"`
}

// RemoveBusinessAccountProfilePhotoParams https://core.telegram.org/bots/api#removebusinessaccountprofilephoto
type RemoveBusinessAccountProfilePhotoParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	IsPublic             bool   `json:"is_public,omitempty"`
}

// SetBusinessAccountGiftSettingsParams https://core.telegram.org/bots/api#setbusinessaccountgiftsettings
type SetBusinessAccountGiftSettingsParams struct {
	BusinessConnectionID string                   `json:"business_connection_id"`
	ShowGiftButton       bool                     `json:"show_gift_button"`
	AcceptedGiftTypes    models.AcceptedGiftTypes `json:"accepted_gift_types"`
}

// GetBusinessAccountStarBalanceParams https://core.telegram.org/bots/api#getbusinessaccountstarbalance
type GetBusinessAccountStarBalanceParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
}

// TransferBusinessAccountStarsParams https://core.telegram.org/bots/api#transferbusinessaccountstars
type TransferBusinessAccountStarsParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	StarCount            int    `json:"star_count"`
}

// GetBusinessAccountGiftsParams https://core.telegram.org/bots/api#getbusinessaccountgifts
type GetBusinessAccountGiftsParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	ExcludeUnsaved       bool   `json:"exclude_unsaved,omitempty"`
	ExcludeSaved         bool   `json:"exclude_saved,omitempty"`
	ExcludeUnlimited     bool   `json:"exclude_unlimited,omitempty"`
	ExcludeLimited       bool   `json:"exclude_limited,omitempty"`
	ExcludeUnique        bool   `json:"exclude_unique,omitempty"`
	SortByPrice          bool   `json:"sort_by_price,omitempty"`
	Offset               string `json:"offset,omitempty"`
	Limit                int    `json:"limit,omitempty"`
}

// ConvertGiftToStarsParams https://core.telegram.org/bots/api#convertgifttostars
type ConvertGiftToStarsParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	OwnedGiftID          string `json:"owned_gift_id"`
}

// UpgradeGiftParams https://core.telegram.org/bots/api#upgradegift
type UpgradeGiftParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	OwnedGiftID          string `json:"owned_gift_id"`
	KeepOriginalDetails  bool   `json:"keep_original_details,omitempty"`
	StarCount            int    `json:"star_count,omitempty"`
}

// TransferGiftParams https://core.telegram.org/bots/api#transfergift
type TransferGiftParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	OwnedGiftID          string `json:"owned_gift_id"`
	NewOwnerChatID       int    `json:"new_owner_chat_id"`
	StarCount            int    `json:"star_count"`
}

// PostStoryParams https://core.telegram.org/bots/api#poststory
type PostStoryParams struct {
	BusinessConnectionID string                   `json:"business_connection_id"`
	Content              models.InputStoryContent `json:"content"`
	ActivePeriod         int                      `json:"active_period"`
	Caption              string                   `json:"caption,omitempty"`
	ParseMode            models.ParseMode         `json:"parse_mode,omitempty"`
	CaptionEntities      []models.MessageEntity   `json:"caption_entities,omitempty"`
	Areas                []models.StoryArea       `json:"areas,omitempty"`
	PostToChatPage       bool                     `json:"post_to_chat_page,omitempty"`
	ProtectContent       bool                     `json:"protect_content,omitempty"`
}

// EditStoryParams https://core.telegram.org/bots/api#editstory
type EditStoryParams struct {
	BusinessConnectionID string                   `json:"business_connection_id"`
	StoryID              int                      `json:"story_id"`
	Content              models.InputStoryContent `json:"content"`
	Caption              string                   `json:"caption,omitempty"`
	ParseMode            models.ParseMode         `json:"parse_mode,omitempty"`
	CaptionEntities      []models.MessageEntity   `json:"caption_entities,omitempty"`
	Areas                []models.StoryArea       `json:"areas,omitempty"`
}

// DeleteStoryParams https://core.telegram.org/bots/api#deletestory
type DeleteStoryParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	StoryID              int    `json:"story_id"`
}

// GiftPremiumSubscriptionParams https://core.telegram.org/bots/api#giftpremiumsubscription
type GiftPremiumSubscriptionParams struct {
	UserID        int64                  `json:"user_id"`
	MonthCount    int                    `json:"month_count"`
	StarCount     int                    `json:"star_count"`
	Text          string                 `json:"text,omitempty"`
	TextParseMode string                 `json:"text_parse_mode,omitempty"`
	TextEntities  []models.MessageEntity `json:"text_entities,omitempty"`
}
