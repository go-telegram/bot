package models

// StickerSet https://core.telegram.org/bots/api#stickerset
type StickerSet struct {
	Name          string     `json:"name"`
	Title         string     `json:"title"`
	IsAnimated    bool       `json:"is_animated"`
	IsVideo       bool       `json:"is_video"`
	ContainsMasks bool       `json:"contains_masks"`
	Stickers      []Sticker  `json:"stickers"`
	Thumb         *PhotoSize `json:"thumb,omitempty"`
}
