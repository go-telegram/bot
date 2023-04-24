package models

// BotCommand https://core.telegram.org/bots/api#botcommand
type BotCommand struct {
	Command     string `json:"command" rules:"min:1,max:32"`
	Description string `json:"description" rules:"min:1,max:256"`
}

// BotName https://core.telegram.org/bots/api#botname
type BotName struct {
	Name string `json:"name"`
}

// BotDescription https://core.telegram.org/bots/api#botdescription
type BotDescription struct {
	Description string `json:"description"`
}

// BotShortDescription https://core.telegram.org/bots/api#botshortdescription
type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}
