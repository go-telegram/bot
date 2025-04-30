package models

import (
	"encoding/json"
	"fmt"
)

// Gifts https://core.telegram.org/bots/api#gifts
type Gifts struct {
	Gifts []Gift `json:"gifts"`
}

// Gift https://core.telegram.org/bots/api#gift
type Gift struct {
	ID               string  `json:"id"`
	Sticker          Sticker `json:"sticker"`
	StarCount        int     `json:"star_count"`
	UpgradeStarCount int     `json:"upgrade_star_count,omitempty"`
	TotalCount       int     `json:"total_count,omitempty"`
	RemainingCount   int     `json:"remaining_count,omitempty"`
}

// AcceptedGiftTypes https://core.telegram.org/bots/api#acceptedgifttypes
type AcceptedGiftTypes struct {
	UnlimitedGifts bool `json:"unlimited_gifts"`
	LimitedGifts   bool `json:"limited_gifts"`
	UniqueGifts    bool `json:"unique_gifts"`
}

// OwnedGift https://core.telegram.org/bots/api#ownedgift
type OwnedGift struct {
	Type OwnedGiftType `json:"type"`

	OwnedGiftRegular *OwnedGiftRegular
	OwnedGiftUnique  *OwnedGiftUnique
}

func (g *OwnedGift) UnmarshalJSON(data []byte) error {
	v := struct {
		Type OwnedGiftType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v.Type {
	case OwnedGiftTypeRegular:
		g.Type = OwnedGiftTypeRegular
		g.OwnedGiftRegular = &OwnedGiftRegular{
			Type: OwnedGiftTypeRegular,
		}
		return json.Unmarshal(data, g.OwnedGiftRegular)
	case OwnedGiftTypeUnique:
		g.Type = OwnedGiftTypeUnique
		g.OwnedGiftUnique = &OwnedGiftUnique{
			Type: OwnedGiftTypeUnique,
		}
		return json.Unmarshal(data, g.OwnedGiftUnique)
	}

	return fmt.Errorf("unsupported OwnedGift type")
}

// OwnedGiftType https://core.telegram.org/bots/api#ownedgift
type OwnedGiftType string

const (
	OwnedGiftTypeRegular OwnedGiftType = "regular"
	OwnedGiftTypeUnique  OwnedGiftType = "unique"
)

// OwnedGiftRegular https://core.telegram.org/bots/api#ownedgiftregular
type OwnedGiftRegular struct {
	Type              OwnedGiftType `json:"type"`
	Gift              Gift          `json:"gift"`
	OwnedGiftID       string        `json:"owned_gift_id,omitempty"`
	SenderUser        *User         `json:"sender_user,omitempty"`
	SendDate          int           `json:"send_date"`
	IsSaved           bool          `json:"is_saved,omitempty"`
	CanBeTransferred  bool          `json:"can_be_transferred,omitempty"`
	TransferStarCount int           `json:"transfer_star_count,omitempty"`
}

// OwnedGiftUnique https://core.telegram.org/bots/api#ownedgiftunique
type OwnedGiftUnique struct {
	Type OwnedGiftType `json:"type"`
	Gift UniqueGift    `json:"gift"`
}

// OwnedGifts https://core.telegram.org/bots/api#ownedgifts
type OwnedGifts struct {
	TotalCount int         `json:"total_count"`
	Gifts      []OwnedGift `json:"gifts"`
	NextOffset string      `json:"next_offset,omitempty"`
}

// UniqueGift https://core.telegram.org/bots/api#uniquegift
type UniqueGift struct {
	BaseName string             `json:"base_name"`
	Name     string             `json:"name"`
	Number   int                `json:"number"`
	Model    UniqueGiftModel    `json:"model"`
	Symbol   UniqueGiftSymbol   `json:"symbol"`
	Backdrop UniqueGiftBackdrop `json:"backdrop"`
}

// UniqueGiftModel https://core.telegram.org/bots/api#uniquegiftmodel
type UniqueGiftModel struct {
	Name           string  `json:"name"`
	Sticker        Sticker `json:"sticker"`
	RarityPerMille int     `json:"rarity_per_mille"`
}

// UniqueGiftSymbol https://core.telegram.org/bots/api#uniquegiftsymbol
type UniqueGiftSymbol struct {
	Name           string  `json:"name"`
	Sticker        Sticker `json:"sticker"`
	RarityPerMille int     `json:"rarity_per_mille"`
}

// UniqueGiftBackdropColors https://core.telegram.org/bots/api#uniquegiftbackdropcolors
type UniqueGiftBackdropColors struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	SymbolColor int `json:"symbol_color"`
	TextColor   int `json:"text_color"`
}

// UniqueGiftBackdrop https://core.telegram.org/bots/api#uniquegiftbackdrop
type UniqueGiftBackdrop struct {
	Name           string                   `json:"name"`
	Colors         UniqueGiftBackdropColors `json:"colors"`
	RarityPerMille int                      `json:"rarity_per_mille"`
}

// GiftInfo https://core.telegram.org/bots/api#giftinfo
type GiftInfo struct {
	Gift                    Gift            `json:"gift"`
	OwnedGiftID             string          `json:"owned_gift_id,omitempty"`
	ConvertStarCount        int             `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount int             `json:"prepaid_upgrade_star_count,omitempty"`
	CanBeUpgraded           bool            `json:"can_be_upgraded,omitempty"`
	Text                    string          `json:"text,omitempty"`
	Entities                []MessageEntity `json:"entities,omitempty"`
	IsPrivate               bool            `json:"is_private,omitempty"`
}

// UniqueGiftInfo https://core.telegram.org/bots/api#uniquegiftinfo
type UniqueGiftInfo struct {
	Gift              UniqueGift `json:"gift"`
	Origin            string     `json:"origin"`
	OwnedGiftID       string     `json:"owned_gift_id,omitempty"`
	TransferStarCount int        `json:"transfer_star_count,omitempty"`
}
