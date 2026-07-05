package models

import (
	"encoding/json"
	"testing"
)

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

func TestUnmarshalMessage_replyToStory(t *testing.T) {
	src := `{"date":0,"chat":{"id":123},"message_id":987,"reply_to_story":{"chat":{"id":42},"id":7}}`

	msg := Message{}
	if err := json.Unmarshal([]byte(src), &msg); err != nil {
		t.Fatal(err)
	}

	if msg.ReplyToStory == nil {
		t.Fatal("ReplyToStory is nil")
	}

	if msg.ReplyToStory.ID != 7 {
		t.Fatal("wrong story id")
	}

	if msg.ReplyToStory.Chat.ID != 42 {
		t.Fatal("wrong story chat id")
	}
}
