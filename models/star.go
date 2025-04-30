package models

import (
	"encoding/json"
	"fmt"
)

type TransactionPartnerType string

const (
	TransactionPartnerTypeUser             TransactionPartnerType = "user"
	TransactionPartnerTypeChat             TransactionPartnerType = "chat"
	TransactionPartnerTypeAffiliateProgram TransactionPartnerType = "affiliate_program"
	TransactionPartnerTypeFragment         TransactionPartnerType = "fragment"
	TransactionPartnerTypeTelegramAds      TransactionPartnerType = "telegram_ads"
	TransactionPartnerTypeTelegramApi      TransactionPartnerType = "telegram_api"
	TransactionPartnerTypeOther            TransactionPartnerType = "other"
)

// TransactionPartner https://core.telegram.org/bots/api#transactionpartner
type TransactionPartner struct {
	Type TransactionPartnerType

	User             *TransactionPartnerUser             `json:"user,omitempty"`
	Chat             *TransactionPartnerChat             `json:"chat,omitempty"`
	AffiliateProgram *TransactionPartnerAffiliateProgram `json:"affiliate_program,omitempty"`
	Fragment         *TransactionPartnerFragment         `json:"fragment,omitempty"`
	TelegramAds      *TransactionPartnerTelegramAds      `json:"telegram_ads,omitempty"`
	TelegramApi      *TransactionPartnerTelegramApi      `json:"telegram_api,omitempty"`
	Other            *TransactionPartnerOther            `json:"other,omitempty"`
}

func (m *TransactionPartner) UnmarshalJSON(data []byte) error {
	v := struct {
		Type TransactionPartnerType `json:"type"`
	}{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	switch v.Type {
	case TransactionPartnerTypeUser:
		m.Type = TransactionPartnerTypeUser
		m.User = &TransactionPartnerUser{}
		return json.Unmarshal(data, m.User)
	case TransactionPartnerTypeChat:
		m.Type = TransactionPartnerTypeChat
		m.Chat = &TransactionPartnerChat{}
		return json.Unmarshal(data, m.Chat)
	case TransactionPartnerTypeAffiliateProgram:
		m.Type = TransactionPartnerTypeAffiliateProgram
		m.AffiliateProgram = &TransactionPartnerAffiliateProgram{}
		return json.Unmarshal(data, m.AffiliateProgram)
	case TransactionPartnerTypeFragment:
		m.Type = TransactionPartnerTypeFragment
		m.Fragment = &TransactionPartnerFragment{}
		return json.Unmarshal(data, m.Fragment)
	case TransactionPartnerTypeTelegramAds:
		m.Type = TransactionPartnerTypeTelegramAds
		m.TelegramAds = &TransactionPartnerTelegramAds{}
		return json.Unmarshal(data, m.TelegramAds)
	case TransactionPartnerTypeTelegramApi:
		m.Type = TransactionPartnerTypeTelegramApi
		m.TelegramApi = &TransactionPartnerTelegramApi{}
		return json.Unmarshal(data, m.TelegramApi)
	case TransactionPartnerTypeOther:
		m.Type = TransactionPartnerTypeOther
		m.Other = &TransactionPartnerOther{}
		return json.Unmarshal(data, m.Other)
	}

	return fmt.Errorf("unsupported TransactionPartner type")
}

// AffiliateInfo https://core.telegram.org/bots/api#affiliateinfo
type AffiliateInfo struct {
	AffiliateUser      *User `json:"affiliate_user,omitempty"`
	AffiliateChat      *Chat `json:"affiliate_chat,omitempty"`
	CommissionPerMille int   `json:"commission_per_mille"`
	Amount             int   `json:"amount"`
	NanostarAmount     int   `json:"nanostar_amount"`
}

// TransactionPartnerUser https://core.telegram.org/bots/api#transactionpartneruser
type TransactionPartnerUser struct {
	Type                        TransactionPartnerType `json:"type"`
	TransactionType             string                 `json:"transaction_type"`
	User                        User                   `json:"user"`
	Affiliate                   *AffiliateInfo         `json:"affiliate,omitempty"`
	InvoicePayload              string                 `json:"invoice_payload,omitempty"`
	SubscriptionPeriod          int                    `json:"subscription_period,omitempty"`
	PaidMedia                   []*PaidMedia           `json:"paid_media,omitempty"`
	PaidMediaPayload            string                 `json:"paid_media_payload,omitempty"`
	Gift                        string                 `json:"gift,omitempty"`
	PremiumSubscriptionDuration int                    `json:"premium_subscription_duration,omitempty"`
}

// TransactionPartnerChat https://core.telegram.org/bots/api#transactionpartnerchat
type TransactionPartnerChat struct {
	Type TransactionPartnerType `json:"type"`
	Chat Chat                   `json:"chat"`
	Gift *Gift                  `json:"gift,omitempty"`
}

// TransactionPartnerAffiliateProgram https://core.telegram.org/bots/api#transactionpartneraffiliateprogram
type TransactionPartnerAffiliateProgram struct {
	Type               TransactionPartnerType `json:"type"`
	SponsorUser        *User                  `json:"sponsor_user,omitempty"`
	CommissionPerMille int                    `json:"commission_per_mille"`
}

// TransactionPartnerFragment https://core.telegram.org/bots/api#transactionpartnerfragment
type TransactionPartnerFragment struct {
	Type            TransactionPartnerType  `json:"type"`
	WithdrawalState *RevenueWithdrawalState `json:"withdrawal_state,omitempty"`
}

// TransactionPartnerTelegramAds https://core.telegram.org/bots/api#transactionpartnertelegramads
type TransactionPartnerTelegramAds struct {
	Type TransactionPartnerType `json:"type"`
}

// TransactionPartnerTelegramApi https://core.telegram.org/bots/api#transactionpartnertelegramapi
type TransactionPartnerTelegramApi struct {
	Type         TransactionPartnerType `json:"type"`
	RequestCount int                    `json:"request_count"`
}

// TransactionPartnerOther https://core.telegram.org/bots/api#transactionpartnerother
type TransactionPartnerOther struct {
	Type TransactionPartnerType `json:"type"`
}

type RevenueWithdrawalStateType string

const (
	RevenueWithdrawalStateTypePending   RevenueWithdrawalStateType = "pending"
	RevenueWithdrawalStateTypeSucceeded RevenueWithdrawalStateType = "succeeded"
	RevenueWithdrawalStateTypeFailed    RevenueWithdrawalStateType = "failed"
)

// RevenueWithdrawalState https://core.telegram.org/bots/api#revenuewithdrawalstate
type RevenueWithdrawalState struct {
	Type RevenueWithdrawalStateType `json:"type"`

	Pending   *RevenueWithdrawalStatePending   `json:"pending,omitempty"`
	Succeeded *RevenueWithdrawalStateSucceeded `json:"succeeded,omitempty"`
	Failed    *RevenueWithdrawalStateFailed    `json:"failed,omitempty"`
}

func (m *RevenueWithdrawalState) UnmarshalJSON(data []byte) error {
	v := struct {
		Type RevenueWithdrawalStateType `json:"type"`
	}{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	switch v.Type {
	case RevenueWithdrawalStateTypePending:
		m.Type = RevenueWithdrawalStateTypePending
		m.Pending = &RevenueWithdrawalStatePending{}
		return json.Unmarshal(data, m.Pending)
	case RevenueWithdrawalStateTypeSucceeded:
		m.Type = RevenueWithdrawalStateTypeSucceeded
		m.Succeeded = &RevenueWithdrawalStateSucceeded{}
		return json.Unmarshal(data, m.Succeeded)
	case RevenueWithdrawalStateTypeFailed:
		m.Type = RevenueWithdrawalStateTypeFailed
		m.Failed = &RevenueWithdrawalStateFailed{}
		return json.Unmarshal(data, m.Failed)
	}

	return fmt.Errorf("unsupported RevenueWithdrawalState type")
}

// RevenueWithdrawalStatePending https://core.telegram.org/bots/api#revenuewithdrawalstatepending
type RevenueWithdrawalStatePending struct {
	Type RevenueWithdrawalStateType `json:"type"`
}

// RevenueWithdrawalStateSucceeded https://core.telegram.org/bots/api#revenuewithdrawalstatesucceeded
type RevenueWithdrawalStateSucceeded struct {
	Type RevenueWithdrawalStateType `json:"type"`
	Date int                        `json:"date"`
	URL  string                     `json:"url"`
}

// RevenueWithdrawalStateFailed https://core.telegram.org/bots/api#revenuewithdrawalstatefailed
type RevenueWithdrawalStateFailed struct {
	Type RevenueWithdrawalStateType `json:"type"`
}

// StarTransaction https://core.telegram.org/bots/api#startransaction
type StarTransaction struct {
	ID             string              `json:"id"`
	Amount         int                 `json:"amount"`
	NanostarAmount int                 `json:"nanostar_amount,omitempty"`
	Date           int                 `json:"date"`
	Source         *TransactionPartner `json:"source,omitempty"`
	Receiver       *TransactionPartner `json:"receiver,omitempty"`
}

// StarTransactions https://core.telegram.org/bots/api#startransactions
type StarTransactions struct {
	Transactions []StarTransaction `json:"transactions"`
}

// StarAmount https://core.telegram.org/bots/api#staramount
type StarAmount struct {
	Amount         int `json:"amount"`
	NanostarAmount int `json:"nanostar_amount,omitempty"`
}
