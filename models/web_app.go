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
