package models

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestMessage_Community verifies the 10.2 community service messages decode.
func TestMessage_Community(t *testing.T) {
	src := `{"message_id":1,"date":1,"chat":{"id":1,"type":"supergroup"},"community_chat_added":{"community":{"id":555,"name":"Gophers"}}}`
	var m Message
	if err := json.Unmarshal([]byte(src), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if m.CommunityChatAdded == nil {
		t.Fatal("community_chat_added not decoded")
	}
	if m.CommunityChatAdded.Community.ID != 555 || m.CommunityChatAdded.Community.Name != "Gophers" {
		t.Fatalf("bad community: %+v", m.CommunityChatAdded.Community)
	}
}

// TestMessage_CommunityChatRemoved verifies the empty service message decodes.
func TestMessage_CommunityChatRemoved(t *testing.T) {
	src := `{"message_id":1,"date":1,"chat":{"id":1,"type":"supergroup"},"community_chat_removed":{}}`
	var m Message
	if err := json.Unmarshal([]byte(src), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if m.CommunityChatRemoved == nil {
		t.Fatal("community_chat_removed not decoded")
	}
}

// TestMessage_Ephemeral verifies the ephemeral fields on Message.
func TestMessage_Ephemeral(t *testing.T) {
	src := `{"message_id":1,"date":1,"chat":{"id":1,"type":"supergroup"},"receiver_user":{"id":9,"is_bot":false,"first_name":"A"},"ephemeral_message_id":77}`
	var m Message
	if err := json.Unmarshal([]byte(src), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if m.ReceiverUser == nil || m.ReceiverUser.ID != 9 {
		t.Fatalf("receiver_user not decoded: %+v", m.ReceiverUser)
	}
	if m.EphemeralMessageID != 77 {
		t.Fatalf("ephemeral_message_id = %d", m.EphemeralMessageID)
	}
}

// TestBotSubscriptionUpdated verifies the new Update field.
func TestBotSubscriptionUpdated(t *testing.T) {
	src := `{"update_id":1,"subscription":{"user":{"id":5,"is_bot":false,"first_name":"B"},"invoice_payload":"payload","state":"active"}}`
	var u Update
	if err := json.Unmarshal([]byte(src), &u); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if u.Subscription == nil || u.Subscription.State != "active" {
		t.Fatalf("subscription not decoded: %+v", u.Subscription)
	}
}

// TestBotCommand_IsEphemeral verifies the new BotCommand field marshals.
func TestBotCommand_IsEphemeral(t *testing.T) {
	out, err := json.Marshal(BotCommand{Command: "secret", Description: "d", IsEphemeral: true})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var back map[string]any
	_ = json.Unmarshal(out, &back)
	if back["is_ephemeral"] != true {
		t.Fatalf("is_ephemeral not set: %s", out)
	}
}

// TestReplyParameters_Ephemeral verifies ephemeral reply support and that
// message_id is now omitted when zero.
func TestReplyParameters_Ephemeral(t *testing.T) {
	out, err := json.Marshal(ReplyParameters{EphemeralMessageID: 42})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	s := string(out)
	if want := `"ephemeral_message_id":42`; !strings.Contains(s, want) {
		t.Fatalf("missing %s in %s", want, s)
	}
	if strings.Contains(s, `"message_id"`) {
		t.Fatalf("message_id should be omitted when zero: %s", s)
	}
}

// TestMessage_CommunityChatJoined verifies the 10.3 community join service message decodes.
func TestMessage_CommunityChatJoined(t *testing.T) {
	src := `{"message_id":1,"date":1,"chat":{"id":1,"type":"supergroup"},"community_chat_joined":{"community":{"id":555,"name":"Gophers"}}}`
	var m Message
	if err := json.Unmarshal([]byte(src), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if m.CommunityChatJoined == nil {
		t.Fatal("community_chat_joined not decoded")
	}
	if m.CommunityChatJoined.Community.ID != 555 || m.CommunityChatJoined.Community.Name != "Gophers" {
		t.Fatalf("bad community: %+v", m.CommunityChatJoined.Community)
	}
}
