package models

// BotSubscriptionState https://core.telegram.org/bots/api#botsubscriptionupdated
type BotSubscriptionState string

const (
	BotSubscriptionStateCanceled BotSubscriptionState = "canceled"
	BotSubscriptionStateActive   BotSubscriptionState = "active"
	BotSubscriptionStateFailed   BotSubscriptionState = "failed"
)

// BotSubscriptionUpdated https://core.telegram.org/bots/api#botsubscriptionupdated
//
// Contains information about changes to a user payment subscription toward the
// current bot.
type BotSubscriptionUpdated struct {
	User           User                 `json:"user"`
	InvoicePayload string               `json:"invoice_payload"`
	State          BotSubscriptionState `json:"state"`
}
