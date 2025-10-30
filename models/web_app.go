package models

// See Telegram API docs: https://core.telegram.org/bots/api#webappdata
type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#webappinfo
type WebAppInfo struct {
	URL string `json:"url"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#sentwebappmessage
type SentWebAppMessage struct {
	InlineMessageID string `json:"inline_message_id"`
}
