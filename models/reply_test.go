package models

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalMessageOrigin_user(t *testing.T) {
	src := `{"type":"user","date":123,"sender_user":{"id":123}}`

	var mo MessageOrigin
	err := json.Unmarshal([]byte(src), &mo)
	if err != nil {
		t.Fatal(err)
	}

	if mo.Type != MessageOriginTypeUser {
		t.Fatal("wrong type")
	}

	if mo.MessageOriginUser == nil {
		t.Fatal("MessageOriginUser is nil")
	}

	if mo.MessageOriginUser.Date != 123 {
		t.Fatal("wrong date")
	}

	if mo.MessageOriginUser.SenderUser.ID != 123 {
		t.Fatal("wrong sender user id")
	}
}

func TestUnmarshalMessageOrigin_hidden_user(t *testing.T) {
	src := `{"type":"hidden_user","date":123,"sender_user_name":"bar"}`

	var mo MessageOrigin
	err := json.Unmarshal([]byte(src), &mo)
	if err != nil {
		t.Fatal(err)
	}

	if mo.Type != MessageOriginTypeHiddenUser {
		t.Fatal("wrong type")
	}

	if mo.MessageOriginHiddenUser == nil {
		t.Fatal("MessageOriginHiddenUser is nil")
	}

	if mo.MessageOriginHiddenUser.Date != 123 {
		t.Fatal("wrong date")
	}

	if mo.MessageOriginHiddenUser.SenderUserName != "bar" {
		t.Fatal("wrong sender user name")
	}
}

func TestUnmarshalMessageOrigin_chat(t *testing.T) {
	src := `{"type":"chat","date":123,"sender_chat":{"id":123}}`

	var mo MessageOrigin
	err := json.Unmarshal([]byte(src), &mo)
	if err != nil {
		t.Fatal(err)
	}

	if mo.Type != MessageOriginTypeChat {
		t.Fatal("wrong type")
	}

	if mo.MessageOriginChat == nil {
		t.Fatal("MessageOriginChat is nil")
	}

	if mo.MessageOriginChat.Date != 123 {
		t.Fatal("wrong date")
	}

	if mo.MessageOriginChat.SenderChat.ID != 123 {
		t.Fatal("wrong chat id")
	}
}

func TestUnmarshalMessageOrigin_channel(t *testing.T) {
	src := `{"type":"channel","date":123,"chat":{"id":123}}`

	var mo MessageOrigin
	err := json.Unmarshal([]byte(src), &mo)
	if err != nil {
		t.Fatal(err)
	}

	if mo.Type != MessageOriginTypeChannel {
		t.Fatal("wrong type")
	}

	if mo.MessageOriginChannel == nil {
		t.Fatal("MessageOriginChannel is nil")
	}

	if mo.MessageOriginChannel.Date != 123 {
		t.Fatal("wrong date")
	}

	if mo.MessageOriginChannel.Chat.ID != 123 {
		t.Fatal("wrong chat id")
	}
}

// TestTextQuote_RichTexts verifies a quote chosen from a rich message decodes with
// its rich_texts and may carry an empty text.
func TestTextQuote_RichTexts(t *testing.T) {
	src := `{"text":"","rich_texts":[{"type":"bold","text":"hi"}," there"],"position":3}`
	var q TextQuote
	if err := json.Unmarshal([]byte(src), &q); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(q.RichTexts) != 2 || q.RichTexts[0].RichTextBold == nil || q.RichTexts[1].PlainText != " there" || q.Position != 3 {
		t.Fatalf("rich_texts not decoded: %+v", q)
	}
}

// TestReplyParameters_QuoteRichTexts verifies quote_rich_texts marshals as a RichText
// list and is omitted when empty.
func TestReplyParameters_QuoteRichTexts(t *testing.T) {
	out, err := json.Marshal(ReplyParameters{MessageID: 1, QuoteRichTexts: []RichText{{Type: RichTextTypeBold, RichTextBold: &RichTextBold{Text: RichText{PlainText: "hi"}}}, {PlainText: " there"}}})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if want := `{"message_id":1,"quote_rich_texts":[{"type":"bold","text":"hi"}," there"]}`; string(out) != want {
		t.Fatalf("got %s, want %s", out, want)
	}

	out, err = json.Marshal(ReplyParameters{MessageID: 1})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if want := `{"message_id":1}`; string(out) != want {
		t.Fatalf("got %s, want %s", out, want)
	}
}
