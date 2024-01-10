package models

import (
	"encoding/json"
	"fmt"
)

// ChatBoostUpdated https://core.telegram.org/bots/api#chatboostupdated
type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`
	Boost ChatBoost `json:"boost"`
}

// ChatBoostRemoved https://core.telegram.org/bots/api#chatboostremoved
type ChatBoostRemoved struct {
	Chat       Chat            `json:"chat"`
	BoostID    string          `json:"boost_id"`
	RemoveDate int             `json:"remove_date"`
	Source     ChatBoostSource `json:"source"`
}

// UserChatBoosts https://core.telegram.org/bots/api#userchatboosts
type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}

// ChatBoost https://core.telegram.org/bots/api#chatboost
type ChatBoost struct {
	BoostID        string          `json:"boost_id"`
	AddDate        int             `json:"add_date"`
	ExpirationDate int             `json:"expiration_date"`
	Source         ChatBoostSource `json:"source"`
}

// ChatBoostSource https://core.telegram.org/bots/api#chatboostsource
type ChatBoostSource struct {
	Source ChatBoostSourceType

	ChatBoostSourcePremium  *ChatBoostSourcePremium
	ChatBoostSourceGiftCode *ChatBoostSourceGiftCode
	ChatBoostSourceGiveaway *ChatBoostSourceGiveaway
}

func (cbs *ChatBoostSource) UnmarshalJSON(data []byte) error {
	v := struct {
		Source string `json:"source"`
	}{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	switch v.Source {
	case "premium":
		cbs.Source = ChatBoostSourceTypePremium
		cbs.ChatBoostSourcePremium = &ChatBoostSourcePremium{}
		return json.Unmarshal(data, cbs.ChatBoostSourcePremium)
	case "gift_code":
		cbs.Source = ChatBoostSourceTypeGiftCode
		cbs.ChatBoostSourceGiftCode = &ChatBoostSourceGiftCode{}
		return json.Unmarshal(data, cbs.ChatBoostSourceGiftCode)
	case "giveaway":
		cbs.Source = ChatBoostSourceTypeGiveaway
		cbs.ChatBoostSourceGiveaway = &ChatBoostSourceGiveaway{}
		return json.Unmarshal(data, cbs.ChatBoostSourceGiveaway)
	}

	return fmt.Errorf("unsupported ChatBoostSource type")
}

// ChatBoostSourceType https://core.telegram.org/bots/api#chatboostsource
type ChatBoostSourceType int

const (
	ChatBoostSourceTypePremium ChatBoostSourceType = iota
	ChatBoostSourceTypeGiftCode
	ChatBoostSourceTypeGiveaway
)

// ChatBoostSourcePremium https://core.telegram.org/bots/api#chatboostsourcepremium
type ChatBoostSourcePremium struct {
	Source string `json:"source"` // always “premium”
	User   User   `json:"user"`
}

// ChatBoostSourceGiftCode https://core.telegram.org/bots/api#chatboostsourcegiftcode
type ChatBoostSourceGiftCode struct {
	Source string `json:"source"` // always “gift_code”
	User   User   `json:"user"`
}

// ChatBoostSourceGiveaway https://core.telegram.org/bots/api#chatboostsourcegiveaway
type ChatBoostSourceGiveaway struct {
	Source string `json:"source"` // always “giveaway”
	User   User   `json:"user"`
}
