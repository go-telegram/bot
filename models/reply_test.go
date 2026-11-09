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

func TestTextQuote_RichTexts(t *testing.T) {
	src := `{"text":"","entities":null,"rich_texts":[{"type":"bold","text":"hi"}],"position":3}`
	var quote TextQuote
	if err := json.Unmarshal([]byte(`{"text":"","rich_texts":[{"type":"bold","text":"hi"}],"position":3}`), &quote); err != nil {
		t.Fatal(err)
	}
	if quote.Text != "" || len(quote.RichTexts) != 1 || quote.RichTexts[0].RichTextBold == nil {
		t.Fatalf("%+v", quote)
	}
	out, err := json.Marshal(quote)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != `{"text":"","rich_texts":[{"type":"bold","text":"hi"}],"position":3}` {
		t.Fatalf("got %s want preserved empty text, got compared to %s", out, src)
	}
}

func TestReplyParameters_QuoteRichTexts(t *testing.T) {
	out, err := json.Marshal(ReplyParameters{
		MessageID: 4,
		QuoteRichTexts: []RichText{{
			Type:         RichTextTypeBold,
			RichTextBold: &RichTextBold{Text: RichText{PlainText: "hi"}},
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != `{"message_id":4,"quote_rich_texts":[{"type":"bold","text":"hi"}]}` {
		t.Fatalf("got %s", out)
	}
}
