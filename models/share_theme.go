package models

import "encoding/json"

// ChatShareThemeType https://core.telegram.org/bots/api#chatsharetheme
type ChatShareThemeType string

const (
	ChatShareThemeTypeEmoji      ChatShareThemeType = "emoji"
	ChatShareThemeTypeUniqueGift ChatShareThemeType = "unique_gift"
)

// ChatShareTheme https://core.telegram.org/bots/api#chatsharetheme
type ChatShareTheme struct {
	Type ChatShareThemeType

	ChatShareThemeEmoji      *ChatShareThemeEmoji
	ChatShareThemeUniqueGift *ChatShareThemeUniqueGift
}

func (t *ChatShareTheme) UnmarshalJSON(data []byte) error {
	v := struct {
		Type ChatShareThemeType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if v.Type == "" {
		return missingDiscriminator("ChatShareTheme")
	}

	t.Type = v.Type

	switch v.Type {
	case ChatShareThemeTypeEmoji:
		t.ChatShareThemeEmoji = &ChatShareThemeEmoji{}
		return json.Unmarshal(data, t.ChatShareThemeEmoji)
	case ChatShareThemeTypeUniqueGift:
		t.ChatShareThemeUniqueGift = &ChatShareThemeUniqueGift{}
		return json.Unmarshal(data, t.ChatShareThemeUniqueGift)
	}

	return nil
}

func (t ChatShareTheme) MarshalJSON() ([]byte, error) {
	switch t.Type {
	case ChatShareThemeTypeEmoji:
		return marshalVariant("ChatShareTheme", t.Type, t.ChatShareThemeEmoji, func(v *ChatShareThemeEmoji) { v.Type = t.Type })
	case ChatShareThemeTypeUniqueGift:
		return marshalVariant("ChatShareTheme", t.Type, t.ChatShareThemeUniqueGift, func(v *ChatShareThemeUniqueGift) { v.Type = t.Type })
	}

	return marshalUnknownVariant("ChatShareTheme", "type", t.Type)
}

// ChatShareThemeEmoji https://core.telegram.org/bots/api#chatsharethemeemoji
type ChatShareThemeEmoji struct {
	Type                 ChatShareThemeType `json:"type"`
	Emojis               []EmojiType        `json:"emojis"`
	BackgroundBackdropID int                `json:"background_backdrop_id,omitempty"`
}

// ChatShareThemeUniqueGift https://core.telegram.org/bots/api#chatsharethemeuniquegift
type ChatShareThemeUniqueGift struct {
	Type ChatShareThemeType `json:"type"`
	Gift UniqueGift         `json:"gift"`
}

// EmojiTypeType https://core.telegram.org/bots/api#emojitype
type EmojiTypeType string

const (
	EmojiTypeTypeEmoji       EmojiTypeType = "emoji"
	EmojiTypeTypeCustomEmoji EmojiTypeType = "custom_emoji"
)

// EmojiType https://core.telegram.org/bots/api#emojitype
type EmojiType struct {
	Type EmojiTypeType

	EmojiTypeEmoji       *EmojiTypeEmoji
	EmojiTypeCustomEmoji *EmojiTypeCustomEmoji
}

func (e *EmojiType) UnmarshalJSON(data []byte) error {
	v := struct {
		Type EmojiTypeType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if v.Type == "" {
		return missingDiscriminator("EmojiType")
	}

	e.Type = v.Type

	switch v.Type {
	case EmojiTypeTypeEmoji:
		e.EmojiTypeEmoji = &EmojiTypeEmoji{}
		return json.Unmarshal(data, e.EmojiTypeEmoji)
	case EmojiTypeTypeCustomEmoji:
		e.EmojiTypeCustomEmoji = &EmojiTypeCustomEmoji{}
		return json.Unmarshal(data, e.EmojiTypeCustomEmoji)
	}

	return nil
}

func (e EmojiType) MarshalJSON() ([]byte, error) {
	switch e.Type {
	case EmojiTypeTypeEmoji:
		return marshalVariant("EmojiType", e.Type, e.EmojiTypeEmoji, func(v *EmojiTypeEmoji) { v.Type = e.Type })
	case EmojiTypeTypeCustomEmoji:
		return marshalVariant("EmojiType", e.Type, e.EmojiTypeCustomEmoji, func(v *EmojiTypeCustomEmoji) { v.Type = e.Type })
	}

	return marshalUnknownVariant("EmojiType", "type", e.Type)
}

// EmojiTypeEmoji https://core.telegram.org/bots/api#emojitypeemoji
type EmojiTypeEmoji struct {
	Type  EmojiTypeType `json:"type"`
	Emoji string        `json:"emoji"`
}

// EmojiTypeCustomEmoji https://core.telegram.org/bots/api#emojitypecustomemoji
type EmojiTypeCustomEmoji struct {
	Type          EmojiTypeType `json:"type"`
	CustomEmojiID string        `json:"custom_emoji_id"`
}

// ChatShareThemeBackdrop https://core.telegram.org/bots/api#chatsharethemebackdrop
type ChatShareThemeBackdrop struct {
	ID     int                      `json:"id"`
	Name   string                   `json:"name"`
	Colors UniqueGiftBackdropColors `json:"colors"`
}

// ChatShareThemeBackdrops https://core.telegram.org/bots/api#chatsharethemebackdrops
type ChatShareThemeBackdrops struct {
	Backdrops []ChatShareThemeBackdrop `json:"backdrops"`
}
