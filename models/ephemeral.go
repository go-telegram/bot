package models

// EphemeralMessageParameters https://core.telegram.org/bots/api#ephemeralmessageparameters
//
// Describes the parameters of an ephemeral message to send.
type EphemeralMessageParameters struct {
	ReceiverUserID              int64  `json:"receiver_user_id"`
	CallbackQueryID             string `json:"callback_query_id,omitempty"`
	ReplaceCallbackQueryMessage bool   `json:"replace_callback_query_message,omitempty"`
}
