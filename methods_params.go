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

type SendMessageParams struct {
	ChatID                   any                    `json:"chat_id"`
	MessageThreadID          int                    `json:"message_thread_id,omitempty"`
	Text                     string                 `json:"text"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	Entities                 []models.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    bool                   `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type ForwardMessageParams struct {
	ChatID              any    `json:"chat_id"`
	MessageThreadID     int    `json:"message_thread_id,omitempty"`
	FromChatID          string `json:"from_chat_id"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ProtectContent      bool   `json:"protect_content,omitempty"`
	MessageID           int    `json:"message_id"`
}

type CopyMessageParams struct {
	ChatID                   any                    `json:"chat_id"`
	MessageThreadID          int                    `json:"message_thread_id,omitempty"`
	FromChatID               string                 `json:"from_chat_id"`
	MessageID                int                    `json:"message_id"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendPhotoParams struct {
	ChatID                   any                    `json:"chat_id"`
	MessageThreadID          int                    `json:"message_thread_id,omitempty"`
	Photo                    models.InputFile       `json:"photo"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler               bool                   `json:"has_spoiler,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendAudioParams struct {
	ChatID                   any                    `json:"chat_id"`
	MessageThreadID          int                    `json:"message_thread_id,omitempty"`
	Audio                    models.InputFile       `json:"audio"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	Duration                 int                    `json:"duration,omitempty"`
	Performer                string                 `json:"performer,omitempty"`
	Title                    string                 `json:"title,omitempty"`
	Thumbnail                models.InputFile       `json:"thumbnail,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendDocumentParams struct {
	ChatID                      any                    `json:"chat_id"`
	MessageThreadID             int                    `json:"message_thread_id,omitempty"`
	Document                    models.InputFile       `json:"document"`
	Thumbnail                   models.InputFile       `json:"thumbnail,omitempty"`
	Caption                     string                 `json:"caption,omitempty"`
	ParseMode                   models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities             []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                   `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool                   `json:"disable_notification,omitempty"`
	ProtectContent              bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID            int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply    bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup                 models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendVideoParams struct {
	ChatID                   any                    `json:"chat_id"`
	MessageThreadID          int                    `json:"message_thread_id,omitempty"`
	Video                    models.InputFile       `json:"video"`
	Duration                 int                    `json:"duration,omitempty"`
	Width                    int                    `json:"width,omitempty"`
	Height                   int                    `json:"height,omitempty"`
	Thumbnail                models.InputFile       `json:"thumbnail,omitempty"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler               bool                   `json:"has_spoiler,omitempty"`
	SupportsStreaming        bool                   `json:"supports_streaming,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendAnimationParams struct {
	ChatID                   any                    `json:"chat_id"`
	MessageThreadID          int                    `json:"message_thread_id,omitempty"`
	Animation                models.InputFile       `json:"animation"`
	Duration                 int                    `json:"duration,omitempty"`
	Width                    int                    `json:"width,omitempty"`
	Height                   int                    `json:"height,omitempty"`
	Thumbnail                models.InputFile       `json:"thumbnail,omitempty"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler               bool                   `json:"has_spoiler,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendVoiceParams struct {
	ChatID                   any                    `json:"chat_id"`
	MessageThreadID          int                    `json:"message_thread_id,omitempty"`
	Voice                    models.InputFile       `json:"voice"`
	Caption                  string                 `json:"caption,omitempty"`
	ParseMode                models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities          []models.MessageEntity `json:"caption_entities,omitempty"`
	Duration                 int                    `json:"duration,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendVideoNoteParams struct {
	ChatID                   any                `json:"chat_id"`
	MessageThreadID          int                `json:"message_thread_id,omitempty"`
	VideoNote                models.InputFile   `json:"video_note"`
	Duration                 int                `json:"duration,omitempty"`
	Length                   int                `json:"length,omitempty"`
	Thumbnail                models.InputFile   `json:"thumbnail,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendMediaGroupParams struct {
	ChatID                   any                 `json:"chat_id"`
	MessageThreadID          int                 `json:"message_thread_id,omitempty"`
	Media                    []models.InputMedia `json:"media"`
	DisableNotification      bool                `json:"disable_notification,omitempty"`
	ProtectContent           bool                `json:"protect_content,omitempty"`
	ReplyToMessageID         int                 `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                `json:"allow_sending_without_reply,omitempty"`
}

type SendLocationParams struct {
	ChatID                   any                `json:"chat_id"`
	MessageThreadID          int                `json:"message_thread_id,omitempty"`
	Latitude                 float64            `json:"latitude"`
	Longitude                float64            `json:"longitude"`
	HorizontalAccuracy       float64            `json:"horizontal_accuracy,omitempty"`
	LivePeriod               int                `json:"live_period,omitempty"`
	Heading                  int                `json:"heading,omitempty"`
	ProximityAlertRadius     int                `json:"proximity_alert_radius,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type EditMessageLiveLocationParams struct {
	ChatID               any                `json:"chat_id"`
	MessageID            int                `json:"message_id,omitempty"`
	InlineMessageID      string             `json:"inline_message_id,omitempty"`
	Latitude             float64            `json:"latitude"`
	Longitude            float64            `json:"longitude"`
	HorizontalAccuracy   float64            `json:"horizontal_accuracy,omitempty"`
	Heading              int                `json:"heading,omitempty"`
	ProximityAlertRadius int                `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type StopMessageLiveLocationParams struct {
	ChatID          any                `json:"chat_id,omitempty"`
	MessageID       int                `json:"message_id,omitempty"`
	InlineMessageID string             `json:"inline_message_id,omitempty"`
	ReplyMarkup     models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendVenueParams struct {
	ChatID                   any                `json:"chat_id"`
	MessageThreadID          int                `json:"message_thread_id,omitempty"`
	Latitude                 float64            `json:"latitude"`
	Longitude                float64            `json:"longitude"`
	Title                    string             `json:"title"`
	Address                  string             `json:"address"`
	FoursquareID             string             `json:"foursquare_id,omitempty"`
	FoursquareType           string             `json:"foursquare_type,omitempty"`
	GooglePlaceID            string             `json:"google_place_id,omitempty"`
	GooglePlaceType          string             `json:"google_place_type,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendContactParams struct {
	ChatID                   any                `json:"chat_id"`
	MessageThreadID          int                `json:"message_thread_id,omitempty"`
	PhoneNumber              string             `json:"phone_number"`
	FirstName                string             `json:"first_name"`
	LastName                 string             `json:"last_name,omitempty"`
	VCard                    string             `json:"vcard,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendPollParams struct {
	ChatID                   any                    `json:"chat_id"`
	MessageThreadID          int                    `json:"message_thread_id,omitempty"`
	Question                 string                 `json:"question"`
	Options                  []string               `json:"options"`
	IsAnonymous              *bool                  `json:"is_anonymous,omitempty"`
	Type                     string                 `json:"type,omitempty"`
	AllowsMultipleAnswers    bool                   `json:"allows_multiple_answers,omitempty"`
	CorrectOptionID          int                    `json:"correct_option_id"`
	Explanation              string                 `json:"explanation,omitempty"`
	ExplanationParseMode     string                 `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities      []models.MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod               int                    `json:"open_period,omitempty"`
	CloseDate                int                    `json:"close_date,omitempty"`
	IsClosed                 bool                   `json:"is_closed,omitempty"`
	DisableNotification      bool                   `json:"disable_notification,omitempty"`
	ProtectContent           bool                   `json:"protect_content,omitempty"`
	ReplyToMessageID         int                    `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool                   `json:"allow_sending_without_reply"`
	ReplyMarkup              models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type SendDiceParams struct {
	ChatID                   any                `json:"chat_id"`
	MessageThreadID          int                `json:"message_thread_id,omitempty"`
	Emoji                    string             `json:"emoji,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type SendChatActionParams struct {
	ChatID          any               `json:"chat_id"`
	MessageThreadID int               `json:"message_thread_id,omitempty"`
	Action          models.ChatAction `json:"action"`
}

type GetUserProfilePhotosParams struct {
	UserID int64 `json:"user_id"`
	Offset int   `json:"offset,omitempty"`
	Limit  int   `json:"limit,omitempty"`
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
	Description string `json:"title"`
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
	ChatID     any                    `json:"chat_id"`
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

type EditMessageTextParams struct {
	ChatID                any                    `json:"chat_id"`
	MessageID             int                    `json:"message_id,omitempty"`
	InlineMessageID       string                 `json:"inline_message_id,omitempty"`
	Text                  string                 `json:"text"`
	ParseMode             models.ParseMode       `json:"parse_mode,omitempty"`
	Entities              []models.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview bool                   `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type EditMessageCaptionParams struct {
	ChatID                any                    `json:"chat_id"`
	MessageID             int                    `json:"message_id,omitempty"`
	InlineMessageID       string                 `json:"inline_message_id,omitempty"`
	Caption               string                 `json:"caption,omitempty"`
	ParseMode             models.ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []models.MessageEntity `json:"caption_entities,omitempty"`
	DisableWebPagePreview bool                   `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           models.ReplyMarkup     `json:"reply_markup,omitempty"`
}

type EditMessageMediaParams struct {
	ChatID          any                `json:"chat_id"`
	MessageID       int                `json:"message_id,omitempty"`
	InlineMessageID string             `json:"inline_message_id,omitempty"`
	Media           models.InputMedia  `json:"media"`
	ReplyMarkup     models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type EditMessageReplyMarkupParams struct {
	ChatID          any                `json:"chat_id"`
	MessageID       int                `json:"message_id,omitempty"`
	InlineMessageID string             `json:"inline_message_id,omitempty"`
	ReplyMarkup     models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type StopPollParams struct {
	ChatID      any                `json:"chat_id"`
	MessageID   int                `json:"message_id"`
	ReplyMarkup models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type DeleteMessageParams struct {
	ChatID    any `json:"chat_id"`
	MessageID int `json:"message_id"`
}

type SendStickerParams struct {
	ChatID                   any                `json:"chat_id"`
	MessageThreadID          int                `json:"message_thread_id,omitempty"`
	Sticker                  models.InputFile   `json:"sticker"`
	Emoji                    string             `json:"emoji,omitempty"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type GetStickerSetParams struct {
	Name string `json:"name"`
}

type GetCustomEmojiStickersParams struct {
	CustomEmojiIDs []string `json:"custom_emoji_ids"`
}

type UploadStickerFileParams struct {
	UserID     int64            `json:"user_id"`
	PngSticker models.InputFile `json:"png_sticker"`
}

type CreateNewStickerSetParams struct {
	UserID          int64            `json:"user_id"`
	Name            string           `json:"name"`
	Title           string           `json:"title"`
	Sticker         models.InputFile `json:"sticker"`
	StickerFormat   string           `json:"sticker_format"`
	StickerType     string           `json:"sticker_type,omitempty"`
	NeedsRepainting bool             `json:"needs_repainting,omitempty"`
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
	Thumbnail models.InputFile `json:"thumbnail"`
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

type SendInvoiceParams struct {
	ChatID                    any                   `json:"chat_id"`
	MessageThreadID           int                   `json:"message_thread_id,omitempty"`
	Title                     string                `json:"title"`
	Description               string                `json:"description"`
	Payload                   string                `json:"payload"`
	ProviderToken             string                `json:"provider_token"`
	Currency                  string                `json:"currency"`
	Prices                    []models.LabeledPrice `json:"prices"`
	MaxTipAmount              int                   `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int                 `json:"suggested_tip_amounts,omitempty"`
	StartParameter            string                `json:"start_parameter,omitempty"`
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
	DisableNotification       bool                  `json:"disable_notification,omitempty"`
	ProtectContent            bool                  `json:"protect_content,omitempty"`
	ReplyToMessageID          int                   `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply  bool                  `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup               models.ReplyMarkup    `json:"reply_markup,omitempty"`
}

type CreateInvoiceLinkParams struct {
	Title                     string                `json:"title"`
	Description               string                `json:"description"`
	Payload                   string                `json:"payload"`
	ProviderToken             string                `json:"provider_token"`
	Currency                  string                `json:"currency"`
	Prices                    []models.LabeledPrice `json:"prices"`
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

type SetPassportDataErrorsParams struct {
	UserID int64                         `json:"user_id"`
	Errors []models.PassportElementError `json:"errors"`
}

type SendGameParams struct {
	ChatID                   any                `json:"chat_id"`
	MessageThreadID          int                `json:"message_thread_id,omitempty"`
	GameShorName             string             `json:"game_short_name"`
	DisableNotification      bool               `json:"disable_notification,omitempty"`
	ProtectContent           bool               `json:"protect_content,omitempty"`
	ReplyToMessageID         int                `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool               `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              models.ReplyMarkup `json:"reply_markup,omitempty"`
}

type SetGameScoreParams struct {
	UserID             int64 `json:"user_id"`
	Score              int   `json:"score"`
	Force              bool  `json:"force,omitempty"`
	DisableEditMessage bool  `json:"disable_edit_message,omitempty"`
	ChatID             int   `json:"chat_id,omitempty"`
	MessageID          int   `json:"message_id,omitempty"`
	InlineMessageID    int   `json:"inline_message_id,omitempty"`
}

type GetGameHighScoresParams struct {
	UserID          int64 `json:"user_id"`
	ChatID          int   `json:"chat_id,omitempty"`
	MessageID       int   `json:"message_id,omitempty"`
	InlineMessageID int   `json:"inline_message_id,omitempty"`
}
