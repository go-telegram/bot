package models

// Video https://core.telegram.org/bots/api#video
type Video struct {
	FileID         string      `json:"file_id"`
	FileUniqueID   string      `json:"file_unique_id"`
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	Duration       int         `json:"duration"`
	Thumbnail      *PhotoSize  `json:"thumbnail,omitempty"`
	Cover          []PhotoSize `json:"cover,omitempty"`
	StartTimestamp int         `json:"start_timestamp,omitempty"`
	FileName       string      `json:"file_name,omitempty"`
	MimeType       string      `json:"mime_type,omitempty"`
	FileSize       int64       `json:"file_size,omitempty"`
}
