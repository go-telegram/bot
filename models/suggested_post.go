package models

// See Telegram API docs: https://core.telegram.org/bots/api#suggestedpostparameters
type SuggestedPostParameters struct {
	Price    *SuggestedPostPrice `json:"price,omitempty"`
	SendDate int                 `json:"send_date,omitempty"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#suggestedpostprice
type SuggestedPostPrice struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#suggestedpostinfo
type SuggestedPostInfo struct {
	State    string              `json:"state"`
	Price    *SuggestedPostPrice `json:"price,omitempty"`
	SendDate int                 `json:"send_date,omitempty"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#suggestedpostapproved
type SuggestedPostApproved struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message,omitempty"`
	Price                *SuggestedPostPrice `json:"price,omitempty"`
	SendDate             int                 `json:"send_date,omitempty"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#suggestedpostapprovalfailed
type SuggestedPostApprovalFailed struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message,omitempty"`
	Price                *SuggestedPostPrice `json:"price,omitempty"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#suggestedpostdeclined
type SuggestedPostDeclined struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	Comment              string   `json:"comment,omitempty"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#suggestedpostpaid
type SuggestedPostPaid struct {
	SuggestedPostMessage *Message    `json:"suggested_post_message,omitempty"`
	Currency             string      `json:"currency"`
	Amount               int         `json:"amount,omitempty"`
	StarAmount           *StarAmount `json:"star_amount,omitempty"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#suggestedpostrefunded
type SuggestedPostRefunded struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	Reason               string   `json:"reason"`
}
