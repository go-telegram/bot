package models

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestChatPermissions_JSON_encodesFalseCanSendMessages(t *testing.T) {
	p := ChatPermissions{
		CanSendMessages:      false,
		CanSendPhotos:        true,
		CanSendVideos:        true,
		CanSendOtherMessages: true,
	}
	b, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	if !strings.Contains(s, `"can_send_messages":false`) {
		t.Fatalf("expected explicit false for can_send_messages, got %s", s)
	}
}
