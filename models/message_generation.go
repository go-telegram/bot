package models

import (
	"encoding/json"
	"strconv"
)

// MessageGenerationStopped https://core.telegram.org/bots/api#messagegenerationstopped
//
// Describes an update about a user stopping message generation.
type MessageGenerationStopped struct {
	Chat            Chat `json:"chat"`
	MessageThreadID int  `json:"message_thread_id,omitempty"`
	DraftID         int  `json:"draft_id"`
}

func (m *MessageGenerationStopped) UnmarshalJSON(data []byte) error {
	type alias MessageGenerationStopped
	v := struct {
		*alias
		DraftID json.Number `json:"draft_id"`
	}{alias: (*alias)(m)}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if v.DraftID == "" {
		return nil
	}
	id, err := strconv.Atoi(v.DraftID.String())
	if err != nil {
		return err
	}
	m.DraftID = id
	return nil
}
