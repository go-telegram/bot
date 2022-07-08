package models

// MaskPosition https://core.telegram.org/bots/api#maskposition
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

// Sticker https://core.telegram.org/bots/api#sticker
type Sticker struct {
	FileID           string       `json:"file_id"`
	FileUniqueID     string       `json:"file_unique_id"`
	Width            int          `json:"width"`
	Height           int          `json:"height"`
	IsAnimated       bool         `json:"is_animated"`
	IsVideo          bool         `json:"is_video"`
	Thumb            *PhotoSize   `json:"thumb,omitempty"`
	Emoji            string       `json:"emoji,omitempty"`
	SetName          string       `json:"set_name,omitempty"`
	PremiumAnimation *File        `json:"premium_animation,omitempty"`
	MaskPosition     MaskPosition `json:"mask_position,omitempty"`
	FileSize         int          `json:"file_size,omitempty"`
}
