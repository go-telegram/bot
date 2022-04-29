package models

// BotCommand https://core.telegram.org/bots/api#botcommand
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// BotCommandScope https://core.telegram.org/bots/api#botcommandscope
type BotCommandScope struct {
}

// BotCommandScopeDefault https://core.telegram.org/bots/api#botcommandscopedefault
type BotCommandScopeDefault struct {
	Type string `json:"type"`
}

// BotCommandScopeAllPrivateChats https://core.telegram.org/bots/api#botcommandscopeallprivatechats
type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type"`
}

// todo: implement it
