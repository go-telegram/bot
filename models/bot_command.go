package models

// BotCommand https://core.telegram.org/bots/api#botcommand
type BotCommand struct {
	Command     string `json:"command" rules:"min:1,max:32"`
	Description string `json:"description" rules:"min:1,max:256"`
}
