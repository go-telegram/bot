package models

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestMessage_RichMessage verifies a received Message carrying a rich_message
// (the single inbound integration point of Bot API 10.1) decodes into the
// RichMessage field.
func TestMessage_RichMessage(t *testing.T) {
	src := `{"message_id":42,"date":1,"chat":{"id":1,"type":"private"},"rich_message":{"blocks":[{"type":"thinking","text":"deciding"},{"type":"paragraph","text":"done"}]}}`

	var m Message
	if err := json.Unmarshal([]byte(src), &m); err != nil {
		t.Fatalf("unmarshal message: %v", err)
	}
	if m.RichMessage == nil {
		t.Fatal("rich_message not decoded")
	}
	if len(m.RichMessage.Blocks) != 2 {
		t.Fatalf("expected 2 blocks, got %d", len(m.RichMessage.Blocks))
	}
	if m.RichMessage.Blocks[0].Type != RichBlockTypeThinking {
		t.Fatal("first block should be thinking")
	}
}

// TestInputRichMessage_ThinkingTag verifies the send-side shape: a thinking
// block is carried as the custom <tg-thinking> HTML tag inside HTML.
func TestInputRichMessage_ThinkingTag(t *testing.T) {
	in := InputRichMessage{HTML: "<tg-thinking>reasoning</tg-thinking>final answer"}

	out, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	// json.Marshal HTML-escapes angle brackets; assert the decoded value instead.
	var back InputRichMessage
	if err := json.Unmarshal(out, &back); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if back.HTML != in.HTML {
		t.Fatalf("HTML round-trip mismatch: %q", back.HTML)
	}
	if !strings.Contains(back.HTML, "<tg-thinking>") {
		t.Fatal("expected <tg-thinking> tag preserved")
	}
	// Empty optional fields must be omitted so we never send conflicting html+markdown.
	if strings.Contains(string(out), "markdown") {
		t.Fatalf("empty markdown should be omitted, got %s", out)
	}
}
