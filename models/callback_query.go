package models

// CallbackQuery https://core.telegram.org/bots/api#callbackquery
//
// For a button on a message sent by the bot, Message is set unless
// callback_query_mode of setMySettings uses inline_message_id instead.
// InlineMessageID is set for a message sent via the bot in inline mode,
// and can also be set for the bot's own messages in that mode.
type CallbackQuery struct {
	ID              string                   `json:"id"`
	From            User                     `json:"from,omitempty"`
	Message         MaybeInaccessibleMessage `json:"message,omitempty"`
	InlineMessageID string                   `json:"inline_message_id,omitempty"`
	ChatInstance    string                   `json:"chat_instance,omitempty"`
	Data            string                   `json:"data,omitempty"`
	GameShortName   string                   `json:"game_short_name,omitempty"`
}
