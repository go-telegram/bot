package models

import "testing"

func TestUnmarshalReactionType_emoji(t *testing.T) {
	src := `{"type":"emoji","emoji":"foo"}`

	var rt ReactionType
	err := rt.UnmarshalJSON([]byte(src))
	if err != nil {
		t.Fatal(err)
	}

	if rt.Type != ReactionTypeTypeEmoji {
		t.Fatal("wrong type")
	}

	if rt.ReactionTypeEmoji == nil {
		t.Fatal("ReactionTypeEmoji is nil")
	}

	if rt.ReactionTypeEmoji.Emoji != "foo" {
		t.Fatal("wrong emoji")
	}
}

func TestUnmarshalReactionType_custom_emoji(t *testing.T) {
	src := `{"type":"custom_emoji","custom_emoji_id":"bar"}`

	var rt ReactionType
	err := rt.UnmarshalJSON([]byte(src))
	if err != nil {
		t.Fatal(err)
	}

	if rt.Type != ReactionTypeTypeCustomEmoji {
		t.Fatal("wrong type")
	}

	if rt.ReactionTypeCustomEmoji == nil {
		t.Fatal("ReactionTypeCustomEmoji is nil")
	}

	if rt.ReactionTypeCustomEmoji.CustomEmojiID != "bar" {
		t.Fatal("wrong custom emoji id")
	}
}

func TestUnmarshalReactionType_paid(t *testing.T) {
	src := `{"type":"paid"}`

	var rt ReactionType
	err := rt.UnmarshalJSON([]byte(src))
	if err != nil {
		t.Fatal(err)
	}

	if rt.Type != ReactionTypeTypePaid {
		t.Fatal("wrong type")
	}

	if rt.ReactionTypePaid == nil {
		t.Fatal("ReactionTypePaid is nil")
	}
}
