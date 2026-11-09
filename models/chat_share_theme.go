package models

import (
	"encoding/json"
)

// EmojiKind is the discriminator of EmojiType https://core.telegram.org/bots/api#emojitype
type EmojiKind string

const (
	EmojiKindEmoji       EmojiKind = "emoji"
	EmojiKindCustomEmoji EmojiKind = "custom_emoji"
)

// EmojiType https://core.telegram.org/bots/api#emojitype
//
// EmojiType is a tagged union: Type holds the discriminator and the matching
// variant pointer is populated. An unknown discriminator from a newer Bot API
// release keeps its value in Type and leaves every variant nil.
type EmojiType struct {
	Type EmojiKind

	Emoji       *EmojiTypeEmoji
	CustomEmoji *EmojiTypeCustomEmoji
}

func (e *EmojiType) UnmarshalJSON(data []byte) error {
	v := struct {
		Type EmojiKind `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if v.Type == "" {
		return missingDiscriminator("EmojiType")
	}

	e.Type = v.Type

	switch v.Type {
	case EmojiKindEmoji:
		e.Emoji = &EmojiTypeEmoji{}
		return json.Unmarshal(data, e.Emoji)
	case EmojiKindCustomEmoji:
		e.CustomEmoji = &EmojiTypeCustomEmoji{}
		return json.Unmarshal(data, e.CustomEmoji)
	}

	return nil
}

func (e EmojiType) MarshalJSON() ([]byte, error) {
	switch e.Type {
	case EmojiKindEmoji:
		return marshalVariant("EmojiType", e.Type, e.Emoji, func(v *EmojiTypeEmoji) { v.Type = e.Type })
	case EmojiKindCustomEmoji:
		return marshalVariant("EmojiType", e.Type, e.CustomEmoji, func(v *EmojiTypeCustomEmoji) { v.Type = e.Type })
	}

	return marshalUnknownVariant("EmojiType", "type", e.Type)
}

// EmojiTypeEmoji https://core.telegram.org/bots/api#emojitypeemoji
//
// A regular emoji.
type EmojiTypeEmoji struct {
	Type  EmojiKind `json:"type"`
	Emoji string    `json:"emoji"`
}

// EmojiTypeCustomEmoji https://core.telegram.org/bots/api#emojitypecustomemoji
//
// A custom emoji.
type EmojiTypeCustomEmoji struct {
	Type          EmojiKind `json:"type"`
	CustomEmojiID string    `json:"custom_emoji_id"`
}

// ChatShareThemeType https://core.telegram.org/bots/api#chatsharetheme
type ChatShareThemeType string

const (
	ChatShareThemeTypeEmoji      ChatShareThemeType = "emoji"
	ChatShareThemeTypeUniqueGift ChatShareThemeType = "unique_gift"
)

// ChatShareTheme https://core.telegram.org/bots/api#chatsharetheme
//
// Describes the theme of the screen for sharing a link to a user or a chat.
// ChatShareTheme is a tagged union: Type holds the discriminator and the matching
// variant pointer is populated. An unknown discriminator from a newer Bot API
// release keeps its value in Type and leaves every variant nil.
type ChatShareTheme struct {
	Type ChatShareThemeType

	Emoji      *ChatShareThemeEmoji
	UniqueGift *ChatShareThemeUniqueGift
}

func (c *ChatShareTheme) UnmarshalJSON(data []byte) error {
	v := struct {
		Type ChatShareThemeType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if v.Type == "" {
		return missingDiscriminator("ChatShareTheme")
	}

	c.Type = v.Type

	switch v.Type {
	case ChatShareThemeTypeEmoji:
		c.Emoji = &ChatShareThemeEmoji{}
		return json.Unmarshal(data, c.Emoji)
	case ChatShareThemeTypeUniqueGift:
		c.UniqueGift = &ChatShareThemeUniqueGift{}
		return json.Unmarshal(data, c.UniqueGift)
	}

	return nil
}

func (c ChatShareTheme) MarshalJSON() ([]byte, error) {
	switch c.Type {
	case ChatShareThemeTypeEmoji:
		return marshalVariant("ChatShareTheme", c.Type, c.Emoji, func(v *ChatShareThemeEmoji) { v.Type = c.Type })
	case ChatShareThemeTypeUniqueGift:
		return marshalVariant("ChatShareTheme", c.Type, c.UniqueGift, func(v *ChatShareThemeUniqueGift) { v.Type = c.Type })
	}

	return marshalUnknownVariant("ChatShareTheme", "type", c.Type)
}

// ChatShareThemeEmoji https://core.telegram.org/bots/api#chatsharethemeemoji
//
// The theme with emoji placed on a background.
type ChatShareThemeEmoji struct {
	Type                 ChatShareThemeType `json:"type"`
	Emojis               []EmojiType        `json:"emojis"`
	BackgroundBackdropID int                `json:"background_backdrop_id,omitempty"`
}

// ChatShareThemeUniqueGift https://core.telegram.org/bots/api#chatsharethemeuniquegift
//
// The theme based on a unique gift owned by the user or the chat.
type ChatShareThemeUniqueGift struct {
	Type ChatShareThemeType `json:"type"`
	Gift UniqueGift         `json:"gift"`
}

// ChatShareThemeBackdrop https://core.telegram.org/bots/api#chatsharethemebackdrop
//
// Describes a backdrop, which can be used as the background of a theme of the
// screen for sharing a link to a user or a chat.
type ChatShareThemeBackdrop struct {
	ID     int                      `json:"id"`
	Name   string                   `json:"name"`
	Colors UniqueGiftBackdropColors `json:"colors"`
}

// ChatShareThemeBackdrops https://core.telegram.org/bots/api#chatsharethemebackdrops
type ChatShareThemeBackdrops struct {
	Backdrops []ChatShareThemeBackdrop `json:"backdrops"`
}
