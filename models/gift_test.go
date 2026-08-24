package models

import (
	"encoding/json"
	"testing"
)

// TestUniqueGiftInfo_Text verifies the 10.3 text, entities and is_private fields decode.
func TestUniqueGiftInfo_Text(t *testing.T) {
	src := `{"gift":{"gift_id":"g","base_name":"b","name":"n","number":1},"origin":"transfer","text":"hi","entities":[{"type":"bold","offset":0,"length":2}],"is_private":true}`
	var info UniqueGiftInfo
	if err := json.Unmarshal([]byte(src), &info); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if info.Text != "hi" {
		t.Fatalf("text = %q", info.Text)
	}
	if len(info.Entities) != 1 || info.Entities[0].Type != "bold" {
		t.Fatalf("entities not decoded: %+v", info.Entities)
	}
	if !info.IsPrivate {
		t.Fatal("is_private not decoded")
	}
}
