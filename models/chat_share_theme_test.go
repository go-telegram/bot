package models

import (
	"encoding/json"
	"testing"
)

func roundTrip(t *testing.T, src string, v interface{}) {
	t.Helper()
	if err := json.Unmarshal([]byte(src), v); err != nil {
		t.Fatalf("unmarshal %s: %v", src, err)
	}
	out, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal %s: %v", src, err)
	}
	if string(out) != src {
		t.Fatalf("round-trip mismatch:\n got %s\nwant %s", out, src)
	}
}

func TestEmojiType_Variants(t *testing.T) {
	var e EmojiType
	roundTrip(t, `{"type":"emoji","emoji":"🔥"}`, &e)
	if e.Type != EmojiKindEmoji || e.Emoji == nil || e.Emoji.Emoji != "🔥" {
		t.Fatalf("emoji not decoded: %+v", e)
	}

	var c EmojiType
	roundTrip(t, `{"type":"custom_emoji","custom_emoji_id":"5368324170671202286"}`, &c)
	if c.Type != EmojiKindCustomEmoji || c.CustomEmoji == nil || c.CustomEmoji.CustomEmojiID != "5368324170671202286" {
		t.Fatalf("custom emoji not decoded: %+v", c)
	}

	var u EmojiType
	roundTrip(t, `{"type":"future"}`, &u)
	if u.Type != "future" || u.Emoji != nil || u.CustomEmoji != nil {
		t.Fatalf("unknown variant not tolerated: %+v", u)
	}

	if err := json.Unmarshal([]byte(`{}`), &u); err == nil {
		t.Fatal("expected an error for an emoji without a discriminator")
	}
}

func TestChatShareTheme_Emoji(t *testing.T) {
	var th ChatShareTheme
	roundTrip(t, `{"type":"emoji","emojis":[{"type":"emoji","emoji":"🔥"},{"type":"custom_emoji","custom_emoji_id":"1"}],"background_backdrop_id":3}`, &th)
	if th.Type != ChatShareThemeTypeEmoji || th.Emoji == nil {
		t.Fatalf("emoji theme not decoded: %+v", th)
	}
	if len(th.Emoji.Emojis) != 2 || th.Emoji.Emojis[1].CustomEmoji == nil || th.Emoji.BackgroundBackdropID != 3 {
		t.Fatalf("emoji theme fields: %+v", th.Emoji)
	}
}

func TestChatShareTheme_UniqueGift(t *testing.T) {
	src := `{"type":"unique_gift","gift":{"gift_id":"g","base_name":"Cap","name":"Cap-1","number":1,"model":{"name":"m","sticker":{"file_id":"f","file_unique_id":"u","type":"regular","width":1,"height":1,"is_animated":false,"is_video":false},"rarity_per_mille":1},"symbol":{"name":"s","sticker":{"file_id":"f","file_unique_id":"u","type":"regular","width":1,"height":1,"is_animated":false,"is_video":false},"rarity_per_mille":1},"backdrop":{"name":"b","colors":{"center_color":1,"edge_color":2,"symbol_color":3,"text_color":4},"rarity_per_mille":1}}}`
	var th ChatShareTheme
	if err := json.Unmarshal([]byte(src), &th); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if th.Type != ChatShareThemeTypeUniqueGift || th.UniqueGift == nil || th.UniqueGift.Gift.Name != "Cap-1" {
		t.Fatalf("unique gift theme not decoded: %+v", th)
	}
	if _, err := json.Marshal(th); err != nil {
		t.Fatalf("marshal: %v", err)
	}
}

func TestChatShareTheme_Unknown(t *testing.T) {
	var th ChatShareTheme
	roundTrip(t, `{"type":"future"}`, &th)
	if th.Type != "future" || th.Emoji != nil || th.UniqueGift != nil {
		t.Fatalf("unknown variant not tolerated: %+v", th)
	}
	if err := json.Unmarshal([]byte(`{}`), &th); err == nil {
		t.Fatal("expected an error for a theme without a discriminator")
	}
}

func TestChatFullInfo_ShareTheme(t *testing.T) {
	src := `{"id":1,"type":"private","max_reaction_count":0,"bio":"","join_to_send_messages":false,"join_by_request":false,"accepted_gift_types":{"unlimited_gifts":false,"limited_gifts":false,"unique_gifts":false,"premium_subscription":false},"share_theme":{"type":"emoji","emojis":[{"type":"emoji","emoji":"🎉"}]}}`
	var c ChatFullInfo
	if err := json.Unmarshal([]byte(src), &c); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if c.ShareTheme == nil || c.ShareTheme.Emoji == nil || len(c.ShareTheme.Emoji.Emojis) != 1 {
		t.Fatalf("share_theme not decoded: %+v", c.ShareTheme)
	}
}

func TestChatShareThemeBackdrops(t *testing.T) {
	src := `{"backdrops":[{"id":1,"name":"Sunset","colors":{"center_color":1,"edge_color":2,"symbol_color":3,"text_color":4}}]}`
	var b ChatShareThemeBackdrops
	if err := json.Unmarshal([]byte(src), &b); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(b.Backdrops) != 1 || b.Backdrops[0].ID != 1 || b.Backdrops[0].Name != "Sunset" || b.Backdrops[0].Colors.EdgeColor != 2 {
		t.Fatalf("backdrops not decoded: %+v", b)
	}
}
