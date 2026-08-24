package models

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestInlineKeyboardButton_Disabled verifies a disabled button is encoded as an
// empty object, which is how Telegram represents DisabledButton.
func TestInlineKeyboardButton_Disabled(t *testing.T) {
	out, err := json.Marshal(InlineKeyboardButton{Text: "Off", Disabled: &DisabledButton{}})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if want := `{"text":"Off","disabled":{}}`; string(out) != want {
		t.Fatalf("got %s, want %s", out, want)
	}
}

// TestInlineKeyboardMarkup_ForceReply verifies the 10.3 force_reply flag is sent
// only when set.
func TestInlineKeyboardMarkup_ForceReply(t *testing.T) {
	out, err := json.Marshal(InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{}, ForceReply: true})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !strings.Contains(string(out), `"force_reply":true`) {
		t.Fatalf("missing force_reply in %s", out)
	}

	out, _ = json.Marshal(InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{}})
	if strings.Contains(string(out), `"force_reply"`) {
		t.Fatalf("force_reply should be omitted when false: %s", out)
	}
}

// TestReplyKeyboardMarkup_ForceReply verifies the 10.3 force_reply flag is sent
// only when set.
func TestReplyKeyboardMarkup_ForceReply(t *testing.T) {
	out, err := json.Marshal(ReplyKeyboardMarkup{Keyboard: [][]KeyboardButton{}, ForceReply: true})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !strings.Contains(string(out), `"force_reply":true`) {
		t.Fatalf("missing force_reply in %s", out)
	}

	out, _ = json.Marshal(ReplyKeyboardMarkup{Keyboard: [][]KeyboardButton{}})
	if strings.Contains(string(out), `"force_reply"`) {
		t.Fatalf("force_reply should be omitted when false: %s", out)
	}
}
