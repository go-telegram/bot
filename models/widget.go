package models

import "encoding/json"

// Widget https://core.telegram.org/bots/api#widget
type Widget struct {
	ID         string `json:"id,omitempty"`
	Width      int    `json:"width,omitempty"`
	Height     int    `json:"height,omitempty"`
	HTML       string `json:"html"`
	PublicData string `json:"public_data,omitempty"`
	ChatData   string `json:"chat_data,omitempty"`
}

// WidgetQueryChatType is the chat a widget in an inline message was sent to.
type WidgetQueryChatType string

const (
	WidgetQueryChatTypeSender     WidgetQueryChatType = "sender"
	WidgetQueryChatTypePrivate    WidgetQueryChatType = "private"
	WidgetQueryChatTypeGroup      WidgetQueryChatType = "group"
	WidgetQueryChatTypeSupergroup WidgetQueryChatType = "supergroup"
	WidgetQueryChatTypeChannel    WidgetQueryChatType = "channel"
)

// WidgetQuery https://core.telegram.org/bots/api#widgetquery
type WidgetQuery struct {
	ID              string              `json:"id"`
	From            User                `json:"from"`
	Chat            *Chat               `json:"chat,omitempty"`
	MessageID       int                 `json:"message_id,omitempty"`
	InlineMessageID string              `json:"inline_message_id,omitempty"`
	ChatInstance    string              `json:"chat_instance"`
	ChatType        WidgetQueryChatType `json:"chat_type,omitempty"`
	Type            WidgetQueryType     `json:"type"`
	Data            string              `json:"data"`
	WidgetID        string              `json:"widget_id,omitempty"`
}

// SharedWidgetQuery https://core.telegram.org/bots/api#sharedwidgetquery
type SharedWidgetQuery struct {
	ID       string          `json:"id"`
	From     User            `json:"from"`
	Type     WidgetQueryType `json:"type"`
	Data     string          `json:"data"`
	WidgetID string          `json:"widget_id"`
}

// WidgetQueryTypeType https://core.telegram.org/bots/api#widgetquerytype
type WidgetQueryTypeType string

const (
	WidgetQueryTypeTypeData            WidgetQueryTypeType = "data"
	WidgetQueryTypeTypeSignData        WidgetQueryTypeType = "sign_data"
	WidgetQueryTypeTypeSentTransaction WidgetQueryTypeType = "sent_transaction"
)

// WidgetQueryType https://core.telegram.org/bots/api#widgetquerytype
type WidgetQueryType struct {
	Type WidgetQueryTypeType

	WidgetQueryTypeData            *WidgetQueryTypeData
	WidgetQueryTypeSignData        *WidgetQueryTypeSignData
	WidgetQueryTypeSentTransaction *WidgetQueryTypeSentTransaction
}

func (q *WidgetQueryType) UnmarshalJSON(data []byte) error {
	v := struct {
		Type WidgetQueryTypeType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if v.Type == "" {
		return missingDiscriminator("WidgetQueryType")
	}

	q.Type = v.Type

	switch v.Type {
	case WidgetQueryTypeTypeData:
		q.WidgetQueryTypeData = &WidgetQueryTypeData{}
		return json.Unmarshal(data, q.WidgetQueryTypeData)
	case WidgetQueryTypeTypeSignData:
		q.WidgetQueryTypeSignData = &WidgetQueryTypeSignData{}
		return json.Unmarshal(data, q.WidgetQueryTypeSignData)
	case WidgetQueryTypeTypeSentTransaction:
		q.WidgetQueryTypeSentTransaction = &WidgetQueryTypeSentTransaction{}
		return json.Unmarshal(data, q.WidgetQueryTypeSentTransaction)
	}

	return nil
}

func (q WidgetQueryType) MarshalJSON() ([]byte, error) {
	switch q.Type {
	case WidgetQueryTypeTypeData:
		return marshalVariant("WidgetQueryType", q.Type, q.WidgetQueryTypeData, func(v *WidgetQueryTypeData) { v.Type = q.Type })
	case WidgetQueryTypeTypeSignData:
		return marshalVariant("WidgetQueryType", q.Type, q.WidgetQueryTypeSignData, func(v *WidgetQueryTypeSignData) { v.Type = q.Type })
	case WidgetQueryTypeTypeSentTransaction:
		return marshalVariant("WidgetQueryType", q.Type, q.WidgetQueryTypeSentTransaction, func(v *WidgetQueryTypeSentTransaction) { v.Type = q.Type })
	}

	return marshalUnknownVariant("WidgetQueryType", "type", q.Type)
}

// WidgetQueryTypeData https://core.telegram.org/bots/api#widgetquerytypedata
type WidgetQueryTypeData struct {
	Type       WidgetQueryTypeType `json:"type"`
	WalletCard *TonWalletCard      `json:"wallet_card,omitempty"`
}

// WidgetQueryTypeSignData https://core.telegram.org/bots/api#widgetquerytypesigndata
type WidgetQueryTypeSignData struct {
	Type     WidgetQueryTypeType `json:"type"`
	Wallet   WidgetQueryWallet   `json:"wallet"`
	SignData string              `json:"sign_data"`
}

// WidgetQueryTypeSentTransaction https://core.telegram.org/bots/api#widgetquerytypesenttransaction
type WidgetQueryTypeSentTransaction struct {
	Type    WidgetQueryTypeType `json:"type"`
	Wallet  WidgetQueryWallet   `json:"wallet"`
	MsgHash string              `json:"msg_hash"`
}

// WidgetQueryWallet https://core.telegram.org/bots/api#widgetquerywallet
type WidgetQueryWallet struct {
	Address   string `json:"address"`
	Amount    string `json:"amount"`
	PublicKey string `json:"public_key,omitempty"`
	StateInit string `json:"state_init,omitempty"`
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
