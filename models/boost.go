package models

import (
	"encoding/json"
)

// ChatBoostAdded https://core.telegram.org/bots/api#chatboostadded
type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
}

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
		Source ChatBoostSourceType `json:"source"`
	}{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	if v.Source == "" {
		return missingDiscriminator("ChatBoostSource")
	}

	cbs.Source = v.Source

	switch v.Source {
	case ChatBoostSourceTypePremium:
		cbs.ChatBoostSourcePremium = &ChatBoostSourcePremium{}
		return json.Unmarshal(data, cbs.ChatBoostSourcePremium)
	case ChatBoostSourceTypeGiftCode:
		cbs.ChatBoostSourceGiftCode = &ChatBoostSourceGiftCode{}
		return json.Unmarshal(data, cbs.ChatBoostSourceGiftCode)
	case ChatBoostSourceTypeGiveaway:
		cbs.ChatBoostSourceGiveaway = &ChatBoostSourceGiveaway{}
		return json.Unmarshal(data, cbs.ChatBoostSourceGiveaway)
	}

	return nil
}

func (cbs *ChatBoostSource) MarshalJSON() ([]byte, error) {
	switch cbs.Source {
	case ChatBoostSourceTypePremium:
		cbs.ChatBoostSourcePremium.Source = ChatBoostSourceTypePremium
		return json.Marshal(cbs.ChatBoostSourcePremium)
	case ChatBoostSourceTypeGiftCode:
		cbs.ChatBoostSourceGiftCode.Source = ChatBoostSourceTypeGiftCode
		return json.Marshal(cbs.ChatBoostSourceGiftCode)
	case ChatBoostSourceTypeGiveaway:
		cbs.ChatBoostSourceGiveaway.Source = ChatBoostSourceTypeGiveaway
		return json.Marshal(cbs.ChatBoostSourceGiveaway)
	}

	return marshalUnknownVariant("ChatBoostSource", "source", cbs.Source)
}

// ChatBoostSourceType https://core.telegram.org/bots/api#chatboostsource
type ChatBoostSourceType string

const (
	ChatBoostSourceTypePremium  ChatBoostSourceType = "premium"
	ChatBoostSourceTypeGiftCode ChatBoostSourceType = "gift_code"
	ChatBoostSourceTypeGiveaway ChatBoostSourceType = "giveaway"
)

// ChatBoostSourcePremium https://core.telegram.org/bots/api#chatboostsourcepremium
type ChatBoostSourcePremium struct {
	Source ChatBoostSourceType `json:"source"` // always “premium”
	User   User                `json:"user"`
}

// ChatBoostSourceGiftCode https://core.telegram.org/bots/api#chatboostsourcegiftcode
type ChatBoostSourceGiftCode struct {
	Source ChatBoostSourceType `json:"source"` // always “gift_code”
	User   User                `json:"user"`
}

// ChatBoostSourceGiveaway https://core.telegram.org/bots/api#chatboostsourcegiveaway
type ChatBoostSourceGiveaway struct {
	Source            ChatBoostSourceType `json:"source"` // always “giveaway”
	GiveawayMessageID int                 `json:"giveaway_message_id"`
	User              User                `json:"user"`
	PrizeStarCount    int                 `json:"prize_star_count,omitempty"`
	IsUnclaimed       bool                `json:"is_unclaimed,omitempty"`
}
