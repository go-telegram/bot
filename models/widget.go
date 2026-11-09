package models

import (
	"encoding/json"
)

// Widget https://core.telegram.org/bots/api#widget
//
// Describes an HTML widget in a message.
type Widget struct {
	ID         string `json:"id,omitempty"`
	Width      int    `json:"width,omitempty"`
	Height     int    `json:"height,omitempty"`
	HTML       string `json:"html"`
	PublicData string `json:"public_data,omitempty"`
	ChatData   string `json:"chat_data,omitempty"`
}

// WidgetQueryKind is the discriminator of WidgetQueryType https://core.telegram.org/bots/api#widgetquerytype
type WidgetQueryKind string

const (
	WidgetQueryKindData            WidgetQueryKind = "data"
	WidgetQueryKindSignData        WidgetQueryKind = "sign_data"
	WidgetQueryKindSentTransaction WidgetQueryKind = "sent_transaction"
)

// WidgetQueryType https://core.telegram.org/bots/api#widgetquerytype
//
// WidgetQueryType is a tagged union: Type holds the discriminator and the matching
// variant pointer is populated. An unknown discriminator from a newer Bot API
// release keeps its value in Type and leaves every variant nil.
type WidgetQueryType struct {
	Type WidgetQueryKind

	Data            *WidgetQueryTypeData
	SignData        *WidgetQueryTypeSignData
	SentTransaction *WidgetQueryTypeSentTransaction
}

func (w *WidgetQueryType) UnmarshalJSON(data []byte) error {
	v := struct {
		Type WidgetQueryKind `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if v.Type == "" {
		return missingDiscriminator("WidgetQueryType")
	}

	w.Type = v.Type

	switch v.Type {
	case WidgetQueryKindData:
		w.Data = &WidgetQueryTypeData{}
		return json.Unmarshal(data, w.Data)
	case WidgetQueryKindSignData:
		w.SignData = &WidgetQueryTypeSignData{}
		return json.Unmarshal(data, w.SignData)
	case WidgetQueryKindSentTransaction:
		w.SentTransaction = &WidgetQueryTypeSentTransaction{}
		return json.Unmarshal(data, w.SentTransaction)
	}

	return nil
}

func (w WidgetQueryType) MarshalJSON() ([]byte, error) {
	switch w.Type {
	case WidgetQueryKindData:
		return marshalVariant("WidgetQueryType", w.Type, w.Data, func(v *WidgetQueryTypeData) { v.Type = w.Type })
	case WidgetQueryKindSignData:
		return marshalVariant("WidgetQueryType", w.Type, w.SignData, func(v *WidgetQueryTypeSignData) { v.Type = w.Type })
	case WidgetQueryKindSentTransaction:
		return marshalVariant("WidgetQueryType", w.Type, w.SentTransaction, func(v *WidgetQueryTypeSentTransaction) { v.Type = w.Type })
	}

	return marshalUnknownVariant("WidgetQueryType", "type", w.Type)
}

// WidgetQueryTypeData https://core.telegram.org/bots/api#widgetquerytypedata
//
// The query contains data for the bot.
type WidgetQueryTypeData struct {
	Type       WidgetQueryKind `json:"type"`
	WalletCard *TonWalletCard  `json:"wallet_card,omitempty"`
}

// WidgetQueryTypeSignData https://core.telegram.org/bots/api#widgetquerytypesigndata
//
// The query contains data signed by a TON wallet of the user who sent the query.
type WidgetQueryTypeSignData struct {
	Type     WidgetQueryKind   `json:"type"`
	Wallet   WidgetQueryWallet `json:"wallet"`
	SignData string            `json:"sign_data"`
}

// WidgetQueryTypeSentTransaction https://core.telegram.org/bots/api#widgetquerytypesenttransaction
//
// The query notifies the bot about a transaction sent from a TON wallet of the user.
type WidgetQueryTypeSentTransaction struct {
	Type    WidgetQueryKind   `json:"type"`
	Wallet  WidgetQueryWallet `json:"wallet"`
	MsgHash string            `json:"msg_hash"`
}

// WidgetQueryWallet https://core.telegram.org/bots/api#widgetquerywallet
//
// Describes a TON wallet of the user, which is used in a query from a widget.
type WidgetQueryWallet struct {
	Address   string `json:"address"`
	Amount    string `json:"amount"`
	PublicKey string `json:"public_key,omitempty"`
	StateInit string `json:"state_init,omitempty"`
}

// WidgetQuery https://core.telegram.org/bots/api#widgetquery
//
// An incoming query from a widget in a message sent by the bot or via the bot.
type WidgetQuery struct {
	ID              string          `json:"id"`
	From            User            `json:"from"`
	Chat            *Chat           `json:"chat,omitempty"`
	MessageID       int             `json:"message_id,omitempty"`
	InlineMessageID string          `json:"inline_message_id,omitempty"`
	ChatInstance    string          `json:"chat_instance"`
	ChatType        string          `json:"chat_type,omitempty"`
	Type            WidgetQueryType `json:"type"`
	Data            string          `json:"data"`
	WidgetID        string          `json:"widget_id,omitempty"`
}

// SharedWidgetQuery https://core.telegram.org/bots/api#sharedwidgetquery
//
// An incoming query from a forwarded message with a shared widget created by the bot.
type SharedWidgetQuery struct {
	ID       string          `json:"id"`
	From     User            `json:"from"`
	Type     WidgetQueryType `json:"type"`
	Data     string          `json:"data"`
	WidgetID string          `json:"widget_id"`
}

// InputWidgetMessageContent https://core.telegram.org/bots/api#inputwidgetmessagecontent
type InputWidgetMessageContent struct {
	HTML                  string          `json:"html"`
	Width                 int             `json:"width,omitempty"`
	Height                int             `json:"height,omitempty"`
	PublicData            string          `json:"public_data,omitempty"`
	ChatData              string          `json:"chat_data,omitempty"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	AllowQueries          bool            `json:"allow_queries,omitempty"`
	IsShared              bool            `json:"is_shared,omitempty"`
}

func (InputWidgetMessageContent) inputMessageContentTag() {}

// InputSharedWidgetMessageContent https://core.telegram.org/bots/api#inputsharedwidgetmessagecontent
type InputSharedWidgetMessageContent struct {
	WidgetID              string          `json:"widget_id"`
	ChatData              string          `json:"chat_data,omitempty"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
}

func (InputSharedWidgetMessageContent) inputMessageContentTag() {}
