package models

// WebAppData https://core.telegram.org/bots/api#webappdata
type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

// WebAppInfo https://core.telegram.org/bots/api#webappinfo
type WebAppInfo struct {
	URL string `json:"url"`
}

// SentWebAppMessage https://core.telegram.org/bots/api#sentwebappmessage
type SentWebAppMessage struct {
	InlineMessageID string `json:"inline_message_id"`
}

// BottomButton https://core.telegram.org/bots/webapps#bottombutton
// Describes the main or secondary button at the bottom of the Mini App (Telegram.WebApp.MainButton / SecondaryButton).
type BottomButton struct {
	Type              string `json:"type,omitempty"`
	IconCustomEmojiID string `json:"iconCustomEmojiId,omitempty"`
	Text              string `json:"text,omitempty"`
	Color             string `json:"color,omitempty"`
	TextColor         string `json:"textColor,omitempty"`
	IsVisible         bool   `json:"isVisible,omitempty"`
	IsActive          bool   `json:"isActive,omitempty"`
	HasShineEffect    bool   `json:"hasShineEffect,omitempty"`
	Position          string `json:"position,omitempty"`
	IsProgressVisible bool   `json:"isProgressVisible,omitempty"`
}
