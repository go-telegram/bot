package models

import (
	"encoding/json"
	"testing"
)

// TestUpdate_StoppedMessageGeneration verifies the 10.3 Update field decodes.
func TestUpdate_StoppedMessageGeneration(t *testing.T) {
	src := `{"update_id":1,"stopped_message_generation":{"chat":{"id":5,"type":"private"},"message_thread_id":3,"draft_id":42}}`
	var u Update
	if err := json.Unmarshal([]byte(src), &u); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if u.StoppedMessageGeneration == nil {
		t.Fatal("stopped_message_generation not decoded")
	}
	if u.StoppedMessageGeneration.Chat.ID != 5 || u.StoppedMessageGeneration.MessageThreadID != 3 || u.StoppedMessageGeneration.DraftID != 42 {
		t.Fatalf("bad update: %+v", u.StoppedMessageGeneration)
	}
}
