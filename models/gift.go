package models

import (
	"encoding/json"
	"fmt"
)

// Gifts https://core.telegram.org/bots/api#gifts
type Gifts struct {
	Gifts []Gift `json:"gifts"`
}

// GiftBackground https://core.telegram.org/bots/api#giftbackground
type GiftBackground struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	TextColor   int `json:"text_color"`
}

// Gift https://core.telegram.org/bots/api#gift
type Gift struct {
	ID                     string          `json:"id"`
	Sticker                Sticker         `json:"sticker"`
	StarCount              int             `json:"star_count"`
	UpgradeStarCount       int             `json:"upgrade_star_count,omitempty"`
	IsPremium              bool            `json:"is_premium,omitempty"`
	HasColors              bool            `json:"has_colors,omitempty"`
	TotalCount             int             `json:"total_count,omitempty"`
	RemainingCount         int             `json:"remaining_count,omitempty"`
	PersonalTotalCount     int             `json:"personal_total_count,omitempty"`
	PersonalRemainingCount int             `json:"personal_remaining_count,omitempty"`
	Background             *GiftBackground `json:"background,omitempty"`
	UniqueGiftVariantCount int             `json:"unique_gift_variant_count,omitempty"`
	PublisherChat          *Chat           `json:"publisher_chat,omitempty"`
}

// AcceptedGiftTypes https://core.telegram.org/bots/api#acceptedgifttypes
type AcceptedGiftTypes struct {
	UnlimitedGifts    bool `json:"unlimited_gifts"`
	LimitedGifts      bool `json:"limited_gifts"`
	UniqueGifts       bool `json:"unique_gifts"`
	GiftsFromChannels bool `json:"gifts_from_channels,omitempty"`
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
	Type                    OwnedGiftType   `json:"type"`
	Gift                    Gift            `json:"gift"`
	OwnedGiftID             string          `json:"owned_gift_id,omitempty"`
	SenderUser              *User           `json:"sender_user,omitempty"`
	SendDate                int             `json:"send_date"`
	Text                    string          `json:"text,omitempty"`
	Entities                []MessageEntity `json:"entities,omitempty"`
	IsPrivate               bool            `json:"is_private,omitempty"`
	IsSaved                 bool            `json:"is_saved,omitempty"`
	CanBeUpgraded           bool            `json:"can_be_upgraded,omitempty"`
	WasRefunded             bool            `json:"was_refunded,omitempty"`
	ConvertStarCount        int             `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount int             `json:"prepaid_upgrade_star_count,omitempty"`
	IsUpgradeSeparate       bool            `json:"is_upgrade_separate,omitempty"`
	UniqueGiftNumber        int             `json:"unique_gift_number,omitempty"`
	CanBeTransferred        bool            `json:"can_be_transferred,omitempty"`
	TransferStarCount       int             `json:"transfer_star_count,omitempty"`
}

// OwnedGiftUnique https://core.telegram.org/bots/api#ownedgiftunique
type OwnedGiftUnique struct {
	Type              OwnedGiftType `json:"type"`
	Gift              UniqueGift    `json:"gift"`
	OwnedGiftID       string        `json:"owned_gift_id,omitempty"`
	SenderUser        *User         `json:"sender_user,omitempty"`
	SendDate          int           `json:"send_date"`
	IsSaved           bool          `json:"is_saved,omitempty"`
	CanBeTransferred  bool          `json:"can_be_transferred,omitempty"`
	TransferStarCount int           `json:"transfer_star_count,omitempty"`
	NextTransferDate  int           `json:"next_transfer_date,omitempty"`
}

// OwnedGifts https://core.telegram.org/bots/api#ownedgifts
type OwnedGifts struct {
	TotalCount int         `json:"total_count"`
	Gifts      []OwnedGift `json:"gifts"`
	NextOffset string      `json:"next_offset,omitempty"`
}

// UniqueGiftColors https://core.telegram.org/bots/api#uniquegiftcolors
type UniqueGiftColors struct {
	ModelCustomEmojiID   string `json:"model_custom_emoji_id"`
	SymbolCustomEmojiID  string `json:"symbol_custom_emoji_id"`
	LightThemeMainColor  int    `json:"light_theme_main_color"`
	LightThemeOtherColors []int `json:"light_theme_other_colors"`
	DarkThemeMainColor   int    `json:"dark_theme_main_color"`
	DarkThemeOtherColors []int  `json:"dark_theme_other_colors"`
}

// UniqueGift https://core.telegram.org/bots/api#uniquegift
type UniqueGift struct {
	GiftID           string             `json:"gift_id"`
	BaseName         string             `json:"base_name"`
	Name             string             `json:"name"`
	Number           int                `json:"number"`
	Model            UniqueGiftModel    `json:"model"`
	Symbol           UniqueGiftSymbol   `json:"symbol"`
	Backdrop         UniqueGiftBackdrop `json:"backdrop"`
	IsPremium        bool               `json:"is_premium,omitempty"`
	IsFromBlockchain bool               `json:"is_from_blockchain,omitempty"`
	Colors           *UniqueGiftColors  `json:"colors,omitempty"`
	PublisherChat    *Chat              `json:"publisher_chat,omitempty"`
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
	IsUpgradeSeparate       bool            `json:"is_upgrade_separate,omitempty"`
	CanBeUpgraded           bool            `json:"can_be_upgraded,omitempty"`
	Text                    string          `json:"text,omitempty"`
	Entities                []MessageEntity `json:"entities,omitempty"`
	IsPrivate               bool            `json:"is_private,omitempty"`
	UniqueGiftNumber        int             `json:"unique_gift_number,omitempty"`
}

// UniqueGiftInfo https://core.telegram.org/bots/api#uniquegiftinfo
type UniqueGiftInfo struct {
	Gift               UniqueGift `json:"gift"`
	Origin             string     `json:"origin"`
	LastResaleCurrency string     `json:"last_resale_currency,omitempty"`
	LastResaleAmount   int        `json:"last_resale_amount,omitempty"`
	OwnedGiftID        string     `json:"owned_gift_id,omitempty"`
	TransferStarCount  int        `json:"transfer_star_count,omitempty"`
	NextTransferDate   int        `json:"next_transfer_date,omitempty"`
}
