package models

import "testing"

func TestUnmarshalMaybeInaccessibleMessage_inaccessible(t *testing.T) {
	src := `{"date":0,"chat":{"id":123},"message_id":987}`

	mim := MaybeInaccessibleMessage{}
	err := mim.UnmarshalJSON([]byte(src))
	if err != nil {
		t.Fatal(err)
	}

	if mim.Type != MaybeInaccessibleMessageTypeInaccessibleMessage {
		t.Fatal("wrong type")
	}

	if mim.InaccessibleMessage == nil {
		t.Fatal("InaccessibleMessage is nil")
	}

	if mim.InaccessibleMessage.Chat.ID != 123 {
		t.Fatal("wrong chat id")
	}

	if mim.InaccessibleMessage.MessageID != 987 {
		t.Fatal("wrong message id")
	}
}

func TestUnmarshalMaybeInaccessibleMessage_message(t *testing.T) {
	src := `{"date":42,"chat":{"id":123},"message_id":987}`

	mim := MaybeInaccessibleMessage{}
	err := mim.UnmarshalJSON([]byte(src))
	if err != nil {
		t.Fatal(err)
	}

	if mim.Type != MaybeInaccessibleMessageTypeMessage {
		t.Fatal("wrong type")
	}

	if mim.Message == nil {
		t.Fatal("Message is nil")
	}

	if mim.Message.Chat.ID != 123 {
		t.Fatal("wrong chat id")
	}

	if mim.Message.ID != 987 {
		t.Fatal("wrong message id")
	}
}
