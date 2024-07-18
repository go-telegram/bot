package models

import (
	"encoding/json"
	"fmt"
)

// TextQuote https://core.telegram.org/bots/api#textquote
type TextQuote struct {
	Text     string          `json:"text"`
	Entities []MessageEntity `json:"entities,omitempty"`
	Position int             `json:"position"`
	IsManual bool            `json:"is_manual,omitempty"`
}

// ExternalReplyInfo https://core.telegram.org/bots/api#externalreplyinfo
type ExternalReplyInfo struct {
	Origin             MessageOrigin       `json:"origin"`
	Chat               *Chat               `json:"chat,omitempty"`
	MessageID          int                 `json:"message_id,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	Animation          *Animation          `json:"animation,omitempty"`
	Audio              *Audio              `json:"audio,omitempty"`
	Document           *Document           `json:"document,omitempty"`
	PaidMedia          *PaidMediaInfo      `json:"paid_media,omitempty"`
	Photo              []PhotoSize         `json:"photo,omitempty"`
	Sticker            *Sticker            `json:"sticker,omitempty"`
	Story              *Story              `json:"story,omitempty"`
	Video              *Video              `json:"video,omitempty"`
	VideoNote          *VideoNote          `json:"video_note,omitempty"`
	Voice              *Voice              `json:"voice,omitempty"`
	HasMediaSpoiler    bool                `json:"has_media_spoiler,omitempty"`
	Contact            *Contact            `json:"contact,omitempty"`
	Dice               *Dice               `json:"dice,omitempty"`
	Game               *Game               `json:"game,omitempty"`
	Giveaway           *Giveaway           `json:"giveaway,omitempty"`
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`
	Invoice            *Invoice            `json:"invoice,omitempty"`
	Location           *Location           `json:"location,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	Venue              *Venue              `json:"venue,omitempty"`
}

// ReplyParameters https://core.telegram.org/bots/api#replyparameters
type ReplyParameters struct {
	MessageID                int             `json:"message_id"`
	ChatID                   any             `json:"chat_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	Quote                    string          `json:"quote,omitempty"`
	QuoteParseMode           ParseMode       `json:"quote_parse_mode,omitempty"`
	QuoteEntities            []MessageEntity `json:"quote_entities,omitempty"`
	QuotePosition            int             `json:"quote_position,omitempty"`
}

// MessageOriginType https://core.telegram.org/bots/api#messageorigin
type MessageOriginType string

const (
	MessageOriginTypeUser       MessageOriginType = "user"
	MessageOriginTypeHiddenUser MessageOriginType = "hidden_user"
	MessageOriginTypeChat       MessageOriginType = "chat"
	MessageOriginTypeChannel    MessageOriginType = "channel"
)

// MessageOrigin https://core.telegram.org/bots/api#messageorigin
type MessageOrigin struct {
	Type MessageOriginType

	MessageOriginUser       *MessageOriginUser
	MessageOriginHiddenUser *MessageOriginHiddenUser
	MessageOriginChat       *MessageOriginChat
	MessageOriginChannel    *MessageOriginChannel
}

func (mo *MessageOrigin) UnmarshalJSON(data []byte) error {
	v := struct {
		Type MessageOriginType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v.Type {
	case MessageOriginTypeUser:
		mo.Type = MessageOriginTypeUser
		mo.MessageOriginUser = &MessageOriginUser{}
		return json.Unmarshal(data, mo.MessageOriginUser)
	case MessageOriginTypeHiddenUser:
		mo.Type = MessageOriginTypeHiddenUser
		mo.MessageOriginHiddenUser = &MessageOriginHiddenUser{}
		return json.Unmarshal(data, mo.MessageOriginHiddenUser)
	case MessageOriginTypeChat:
		mo.Type = MessageOriginTypeChat
		mo.MessageOriginChat = &MessageOriginChat{}
		return json.Unmarshal(data, mo.MessageOriginChat)
	case MessageOriginTypeChannel:
		mo.Type = MessageOriginTypeChannel
		mo.MessageOriginChannel = &MessageOriginChannel{}
		return json.Unmarshal(data, mo.MessageOriginChannel)
	}

	return fmt.Errorf("unsupported MessageOrigin type")
}

func (mo *MessageOrigin) MarshalJSON() ([]byte, error) {
	switch mo.Type {
	case MessageOriginTypeUser:
		mo.MessageOriginUser.Type = MessageOriginTypeUser
		return json.Marshal(mo.MessageOriginUser)
	case MessageOriginTypeHiddenUser:
		mo.MessageOriginHiddenUser.Type = MessageOriginTypeHiddenUser
		return json.Marshal(mo.MessageOriginHiddenUser)
	case MessageOriginTypeChat:
		mo.MessageOriginChat.Type = MessageOriginTypeChat
		return json.Marshal(mo.MessageOriginChat)
	case MessageOriginTypeChannel:
		mo.MessageOriginChannel.Type = MessageOriginTypeChannel
		return json.Marshal(mo.MessageOriginChannel)
	}

	return nil, fmt.Errorf("unsupported MessageOrigin type")
}

// MessageOriginUser https://core.telegram.org/bots/api#messageoriginuser
type MessageOriginUser struct {
	Type       MessageOriginType `json:"type"` // always “user”
	Date       int               `json:"date"`
	SenderUser User              `json:"sender_user"`
}

// MessageOriginHiddenUser https://core.telegram.org/bots/api#messageoriginhiddenuser
type MessageOriginHiddenUser struct {
	Type           MessageOriginType `json:"type"` // always “hidden_user”
	Date           int               `json:"date"`
	SenderUserName string            `json:"sender_user_name"`
}

// MessageOriginChat https://core.telegram.org/bots/api#messageoriginchat
type MessageOriginChat struct {
	Type            MessageOriginType `json:"type"` // always “chat”
	Date            int               `json:"date"`
	SenderChat      Chat              `json:"sender_chat"`
	AuthorSignature *string           `json:"author_signature,omitempty"`
}

// MessageOriginChannel https://core.telegram.org/bots/api#messageoriginchannel
type MessageOriginChannel struct {
	Type            MessageOriginType `json:"type"` // always “channel”
	Date            int               `json:"date"`
	Chat            Chat              `json:"chat"`
	MessageID       int               `json:"message_id"`
	AuthorSignature *string           `json:"author_signature,omitempty"`
}
