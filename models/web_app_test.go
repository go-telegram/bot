package models

import (
	"encoding/json"
	"testing"
)

func TestBottomButtonJSON(t *testing.T) {
	src := `{"type":"main","iconCustomEmojiId":"123456789","text":"Continue","color":"#aabbcc","textColor":"#112233","isVisible":true,"isActive":true,"hasShineEffect":true,"position":"left","isProgressVisible":false}`

	var b BottomButton
	if err := json.Unmarshal([]byte(src), &b); err != nil {
		t.Fatal(err)
	}

	if b.Type != "main" {
		t.Fatalf("Type: got %q", b.Type)
	}
	if b.IconCustomEmojiID != "123456789" {
		t.Fatalf("IconCustomEmojiID: got %q", b.IconCustomEmojiID)
	}
	if b.Text != "Continue" {
		t.Fatalf("Text: got %q", b.Text)
	}
	if b.Color != "#aabbcc" {
		t.Fatalf("Color: got %q", b.Color)
	}
	if b.TextColor != "#112233" {
		t.Fatalf("TextColor: got %q", b.TextColor)
	}
	if !b.IsVisible || !b.IsActive || !b.HasShineEffect {
		t.Fatalf("bools: IsVisible=%v IsActive=%v HasShineEffect=%v", b.IsVisible, b.IsActive, b.HasShineEffect)
	}
	if b.Position != "left" {
		t.Fatalf("Position: got %q", b.Position)
	}
	if b.IsProgressVisible {
		t.Fatal("IsProgressVisible: want false")
	}

	enc, err := json.Marshal(&b)
	if err != nil {
		t.Fatal(err)
	}

	var round BottomButton
	if err := json.Unmarshal(enc, &round); err != nil {
		t.Fatal(err)
	}
	if round != b {
		t.Fatalf("round-trip mismatch: %+v vs %+v", round, b)
	}
}
