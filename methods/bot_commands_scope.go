package methods

// BotCommandScope https://core.telegram.org/bots/api#botcommandscope
type BotCommandScope interface {
	botCommandScopeTag()
}

// BotCommandScopeDefault https://core.telegram.org/bots/api#botcommandscopedefault
type BotCommandScopeDefault struct {
	Type string `json:"type"`
}

func (BotCommandScopeDefault) botCommandScopeTag() {}

// BotCommandScopeAllPrivateChats https://core.telegram.org/bots/api#botcommandscopeallprivatechats
type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type"`
}

func (BotCommandScopeAllPrivateChats) botCommandScopeTag() {}

// BotCommandScopeAllGroupChats https://core.telegram.org/bots/api#botcommandscopeallgroupchats
type BotCommandScopeAllGroupChats struct {
	Type string `json:"type"` // todo
}

func (BotCommandScopeAllGroupChats) botCommandScopeTag() {}

// BotCommandScopeAllChatAdministrators https://core.telegram.org/bots/api#botcommandscopeallchatadministrators
type BotCommandScopeAllChatAdministrators struct {
	Type string `json:"type"`
}

func (BotCommandScopeAllChatAdministrators) botCommandScopeTag() {}

// BotCommandScopeChat https://core.telegram.org/bots/api#botcommandscopechat
type BotCommandScopeChat struct {
	Type string `json:"type"`
}

func (BotCommandScopeChat) botCommandScopeTag() {}

// BotCommandScopeChatAdministrators https://core.telegram.org/bots/api#botcommandscopechatadministrators
type BotCommandScopeChatAdministrators struct {
	Type   string `json:"type"`
	ChatID any    `json:"chat_id"`
}

func (BotCommandScopeChatAdministrators) botCommandScopeTag() {}

// BotCommandScopeChatMember https://core.telegram.org/bots/api#botcommandscopechatmember
type BotCommandScopeChatMember struct {
	Type   string `json:"type"`
	ChatID any    `json:"chat_id"`
	UserID int    `json:"user_id"`
}

func (BotCommandScopeChatMember) botCommandScopeTag() {}
