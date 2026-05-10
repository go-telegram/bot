package models

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestInputPollOptionMarshalJSON_embedsInputMediaType(t *testing.T) {
	opt := InputPollOption{
		Text: "opt",
		Media: &InputMediaPhoto{
			Media: "AgACAgIAAxkBAAIC",
		},
	}
	b, err := json.Marshal(opt)
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	if !strings.Contains(s, `"type":"photo"`) {
		t.Fatalf("expected type photo in JSON: %s", s)
	}
}
