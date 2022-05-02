package models

import (
	"encoding/json"
)

// BotCommandScope https://core.telegram.org/bots/api#botcommandscope
type BotCommandScope interface {
	botCommandScopeTag()

	MarshalCustom() ([]byte, error)
}

// BotCommandScopeDefault https://core.telegram.org/bots/api#botcommandscopedefault
type BotCommandScopeDefault struct{}

func (BotCommandScopeDefault) botCommandScopeTag() {}

func (m *BotCommandScopeDefault) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*BotCommandScopeDefault
	}{
		Type:                   "default",
		BotCommandScopeDefault: m,
	}

	return json.Marshal(&ret)
}

// BotCommandScopeAllPrivateChats https://core.telegram.org/bots/api#botcommandscopeallprivatechats
type BotCommandScopeAllPrivateChats struct{}

func (BotCommandScopeAllPrivateChats) botCommandScopeTag() {}

func (m *BotCommandScopeAllPrivateChats) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*BotCommandScopeAllPrivateChats
	}{
		Type:                           "all_private_chats",
		BotCommandScopeAllPrivateChats: m,
	}

	return json.Marshal(&ret)
}

// BotCommandScopeAllGroupChats https://core.telegram.org/bots/api#botcommandscopeallgroupchats
type BotCommandScopeAllGroupChats struct{}

func (BotCommandScopeAllGroupChats) botCommandScopeTag() {}

func (m *BotCommandScopeAllGroupChats) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*BotCommandScopeAllGroupChats
	}{
		Type:                         "all_group_chats",
		BotCommandScopeAllGroupChats: m,
	}

	return json.Marshal(&ret)
}

// BotCommandScopeAllChatAdministrators https://core.telegram.org/bots/api#botcommandscopeallchatadministrators
type BotCommandScopeAllChatAdministrators struct{}

func (BotCommandScopeAllChatAdministrators) botCommandScopeTag() {}

func (m *BotCommandScopeAllChatAdministrators) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*BotCommandScopeAllChatAdministrators
	}{
		Type:                                 "all_chat_administrators",
		BotCommandScopeAllChatAdministrators: m,
	}

	return json.Marshal(&ret)
}

// BotCommandScopeChat https://core.telegram.org/bots/api#botcommandscopechat
type BotCommandScopeChat struct{}

func (BotCommandScopeChat) botCommandScopeTag() {}

func (m *BotCommandScopeChat) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*BotCommandScopeChat
	}{
		Type:                "chat",
		BotCommandScopeChat: m,
	}

	return json.Marshal(&ret)
}

// BotCommandScopeChatAdministrators https://core.telegram.org/bots/api#botcommandscopechatadministrators
type BotCommandScopeChatAdministrators struct {
	ChatID any `json:"chat_id"`
}

func (BotCommandScopeChatAdministrators) botCommandScopeTag() {}

// BotCommandScopeChatMember https://core.telegram.org/bots/api#botcommandscopechatmember
type BotCommandScopeChatMember struct {
	ChatID any `json:"chat_id"`
	UserID int `json:"user_id"`
}

func (BotCommandScopeChatMember) botCommandScopeTag() {}

func (m *BotCommandScopeChatMember) MarshalCustom() ([]byte, error) {
	ret := struct {
		Type string `json:"type"`
		*BotCommandScopeChatMember
	}{
		Type:                      "chat_member",
		BotCommandScopeChatMember: m,
	}

	return json.Marshal(&ret)
}
