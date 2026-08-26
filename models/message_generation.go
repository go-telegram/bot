package models

// MessageGenerationStopped https://core.telegram.org/bots/api#messagegenerationstopped
//
// Describes an update about a user stopping message generation.
type MessageGenerationStopped struct {
	Chat            Chat `json:"chat"`
	MessageThreadID int  `json:"message_thread_id,omitempty"`
	DraftID         int  `json:"draft_id"`
}
