package models

import (
	"encoding/json"
	"testing"
)

func TestChatShareTheme_EmojiRoundTrip(t *testing.T) {
	src := `{"type":"emoji","emojis":[{"type":"emoji","emoji":"🔥"},{"type":"custom_emoji","custom_emoji_id":"42"}],"background_backdrop_id":7}`
	var theme ChatShareTheme
	if err := json.Unmarshal([]byte(src), &theme); err != nil {
		t.Fatal(err)
	}
	if theme.ChatShareThemeEmoji == nil || len(theme.ChatShareThemeEmoji.Emojis) != 2 {
		t.Fatalf("%+v", theme.ChatShareThemeEmoji)
	}
	if theme.ChatShareThemeEmoji.Emojis[0].EmojiTypeEmoji == nil || theme.ChatShareThemeEmoji.Emojis[0].EmojiTypeEmoji.Emoji != "🔥" {
		t.Fatal("regular emoji not decoded")
	}
	if theme.ChatShareThemeEmoji.Emojis[1].EmojiTypeCustomEmoji == nil || theme.ChatShareThemeEmoji.Emojis[1].EmojiTypeCustomEmoji.CustomEmojiID != "42" {
		t.Fatal("custom emoji not decoded")
	}
	out, err := json.Marshal(theme)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != src {
		t.Fatalf("got %s", out)
	}
}

func TestChatShareTheme_UniqueGift(t *testing.T) {
	var theme ChatShareTheme
	src := `{"type":"unique_gift","gift":{"gift_id":"g","base_name":"Plush Pepe","name":"PlushPepe-1","number":1,"model":{"name":"m","sticker":{"file_id":"f","file_unique_id":"u","type":"regular","width":1,"height":1},"rarity_per_mille":10},"symbol":{"name":"s","sticker":{"file_id":"f","file_unique_id":"u","type":"regular","width":1,"height":1},"rarity_per_mille":10},"backdrop":{"name":"b","colors":{"center_color":1,"edge_color":2,"symbol_color":3,"text_color":4},"rarity_per_mille":10}}}`
	if err := json.Unmarshal([]byte(src), &theme); err != nil {
		t.Fatal(err)
	}
	if theme.ChatShareThemeUniqueGift == nil || theme.ChatShareThemeUniqueGift.Gift.Name != "PlushPepe-1" {
		t.Fatalf("%+v", theme.ChatShareThemeUniqueGift)
	}
}

func TestChatShareThemeBackdrops_Decode(t *testing.T) {
	var list ChatShareThemeBackdrops
	src := `{"backdrops":[{"id":7,"name":"Sunset","colors":{"center_color":1,"edge_color":2,"symbol_color":3,"text_color":4}}]}`
	if err := json.Unmarshal([]byte(src), &list); err != nil {
		t.Fatal(err)
	}
	if len(list.Backdrops) != 1 || list.Backdrops[0].Name != "Sunset" || list.Backdrops[0].Colors.TextColor != 4 {
		t.Fatalf("%+v", list)
	}
}
