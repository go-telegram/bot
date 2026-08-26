package models

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestChatMemberAdministrator_CanSendWelcomeMessages verifies the 10.3 right decodes.
func TestChatMemberAdministrator_CanSendWelcomeMessages(t *testing.T) {
	src := `{"status":"administrator","user":{"id":1,"is_bot":false,"first_name":"A"},"can_be_edited":true,"can_send_welcome_messages":true}`
	var cm ChatMember
	if err := json.Unmarshal([]byte(src), &cm); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if cm.Administrator == nil || !cm.Administrator.CanSendWelcomeMessages {
		t.Fatalf("can_send_welcome_messages not decoded: %+v", cm.Administrator)
	}
}

// TestChatAdministratorRights_CanSendWelcomeMessages verifies the 10.3 right marshals.
func TestChatAdministratorRights_CanSendWelcomeMessages(t *testing.T) {
	out, err := json.Marshal(ChatAdministratorRights{CanSendWelcomeMessages: true})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !strings.Contains(string(out), `"can_send_welcome_messages":true`) {
		t.Fatalf("missing can_send_welcome_messages in %s", out)
	}
}
