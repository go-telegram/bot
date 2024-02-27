package models

// CallbackQuery https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	ID              string                   `json:"id"`
	From            User                     `json:"from,omitempty"`
	Message         MaybeInaccessibleMessage `json:"message,omitempty"`
	InlineMessageID string                   `json:"inline_message_id,omitempty"`
	ChatInstance    string                   `json:"chat_instance,omitempty"`
	Data            string                   `json:"data,omitempty"`
	GameShortName   string                   `json:"game_short_name,omitempty"`
}
