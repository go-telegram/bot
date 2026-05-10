package models

// SentGuestMessage https://core.telegram.org/bots/api#sentguestmessage
type SentGuestMessage struct {
	InlineMessageID string `json:"inline_message_id"`
}

// LivePhoto https://core.telegram.org/bots/api#livephoto
type LivePhoto struct {
	Photo        []PhotoSize `json:"photo,omitempty"`
	FileID       string      `json:"file_id"`
	FileUniqueID string      `json:"file_unique_id"`
	Width        int         `json:"width"`
	Height       int         `json:"height"`
	Duration     int         `json:"duration"`
	MimeType     string      `json:"mime_type,omitempty"`
	FileSize     int         `json:"file_size,omitempty"`
}

// PaidMediaLivePhoto https://core.telegram.org/bots/api#paidmedialivephoto
type PaidMediaLivePhoto struct {
	Type      string    `json:"type"`
	LivePhoto LivePhoto `json:"live_photo"`
}

// InputPaidMediaLivePhoto https://core.telegram.org/bots/api#inputpaidmedialivephoto
type InputPaidMediaLivePhoto struct {
	Type  string `json:"type"`
	Media string `json:"media"`
	Photo string `json:"photo"`
}

// BotAccessSettings https://core.telegram.org/bots/api#botaccesssettings
type BotAccessSettings struct {
	IsAccessRestricted bool   `json:"is_access_restricted"`
	AddedUsers         []User `json:"added_users,omitempty"`
}
