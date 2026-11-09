package models

// KeyboardButtonRequestManagedBot https://core.telegram.org/bots/api#keyboardbuttonrequestmanagedbot
type KeyboardButtonRequestManagedBot struct {
	RequestID     int32  `json:"request_id"`
	SuggestedName string `json:"suggested_name,omitempty"`
	// SuggestedUsername must end with "bot". A Premium user may instead suggest a secondary username ending with "agent", "ai", "app", "game", or "gay".
	SuggestedUsername string `json:"suggested_username,omitempty"`
}

// ManagedBotCreated https://core.telegram.org/bots/api#managedbotcreated
type ManagedBotCreated struct {
	Bot *User `json:"bot,omitempty"`
}

// ManagedBotUpdated https://core.telegram.org/bots/api#managedbotupdated
type ManagedBotUpdated struct {
	User *User `json:"user,omitempty"`
	Bot  *User `json:"bot,omitempty"`
}

// PreparedKeyboardButton https://core.telegram.org/bots/api#preparedkeyboardbutton
type PreparedKeyboardButton struct {
	ID string `json:"id"`
}
