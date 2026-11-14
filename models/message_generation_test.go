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

func TestMessageGenerationStopped_DraftIDQuoted(t *testing.T) {
	var m MessageGenerationStopped
	if err := json.Unmarshal([]byte(`{"chat":{"id":1,"type":"private"},"message_thread_id":7,"draft_id":"42"}`), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if m.DraftID != 42 || m.Chat.ID != 1 || m.MessageThreadID != 7 {
		t.Fatalf("bad value: %+v", m)
	}
}

func TestMessageGenerationStopped_DraftIDNumber(t *testing.T) {
	var m MessageGenerationStopped
	if err := json.Unmarshal([]byte(`{"chat":{"id":1,"type":"private"},"draft_id":42}`), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if m.DraftID != 42 || m.Chat.ID != 1 {
		t.Fatalf("bad value: %+v", m)
	}
}

func TestMessageGenerationStopped_DraftIDInvalid(t *testing.T) {
	for _, src := range []string{
		`{"chat":{"id":1,"type":"private"},"draft_id":"abc"}`,
		`{"chat":{"id":1,"type":"private"},"draft_id":"4.2"}`,
	} {
		var m MessageGenerationStopped
		if err := json.Unmarshal([]byte(src), &m); err == nil {
			t.Fatalf("expected error for %s, got %+v", src, m)
		}
	}
}

func TestMessageGenerationStopped_MarshalNumber(t *testing.T) {
	out, err := json.Marshal(MessageGenerationStopped{Chat: Chat{ID: 1, Type: "private"}, DraftID: 42})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(out, &raw); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if string(raw["draft_id"]) != "42" {
		t.Fatalf("draft_id encoded as %s", raw["draft_id"])
	}
}

func TestUpdate_StoppedMessageGenerationQuotedDraftID(t *testing.T) {
	src := `{"update_id":1,"stopped_message_generation":{"chat":{"id":5,"type":"private"},"message_thread_id":3,"draft_id":"42"}}`
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
