package models

import (
	"encoding/json"
	"testing"
)

// TestEphemeralMessageParameters verifies the 10.3 send-side parameters marshal
// and that the optional fields are omitted when zero.
func TestEphemeralMessageParameters(t *testing.T) {
	out, err := json.Marshal(EphemeralMessageParameters{ReceiverUserID: 9})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if want := `{"receiver_user_id":9}`; string(out) != want {
		t.Fatalf("got %s, want %s", out, want)
	}

	out, err = json.Marshal(EphemeralMessageParameters{ReceiverUserID: 9, CallbackQueryID: "q", ReplaceCallbackQueryMessage: true})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if want := `{"receiver_user_id":9,"callback_query_id":"q","replace_callback_query_message":true}`; string(out) != want {
		t.Fatalf("got %s, want %s", out, want)
	}
}
