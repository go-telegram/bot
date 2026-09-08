package models

import (
	"encoding/json"
)

// ReactionTypeType https://core.telegram.org/bots/api#reactiontype
type ReactionTypeType string

const (
	ReactionTypeTypeEmoji       ReactionTypeType = "emoji"
	ReactionTypeTypeCustomEmoji ReactionTypeType = "custom_emoji"
	ReactionTypeTypePaid        ReactionTypeType = "paid"
)

// ReactionType https://core.telegram.org/bots/api#reactiontype
type ReactionType struct {
	Type ReactionTypeType

	ReactionTypeEmoji       *ReactionTypeEmoji
	ReactionTypeCustomEmoji *ReactionTypeCustomEmoji
	ReactionTypePaid        *ReactionTypePaid
}

func (rt *ReactionType) MarshalJSON() ([]byte, error) {
	switch rt.Type {
	case ReactionTypeTypeEmoji:
		return marshalVariant("ReactionType", rt.Type, rt.ReactionTypeEmoji, func(v *ReactionTypeEmoji) { v.Type = rt.Type })
	case ReactionTypeTypeCustomEmoji:
		return marshalVariant("ReactionType", rt.Type, rt.ReactionTypeCustomEmoji, func(v *ReactionTypeCustomEmoji) { v.Type = rt.Type })
	case ReactionTypeTypePaid:
		return marshalVariant("ReactionType", rt.Type, rt.ReactionTypePaid, func(v *ReactionTypePaid) { v.Type = string(rt.Type) })
	}

	return marshalUnknownVariant("ReactionType", "type", rt.Type)
}

func (rt *ReactionType) UnmarshalJSON(data []byte) error {
	v := struct {
		Type ReactionTypeType `json:"type"`
	}{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	if v.Type == "" {
		return missingDiscriminator("ReactionType")
	}

	rt.Type = v.Type

	switch v.Type {
	case ReactionTypeTypeEmoji:
		rt.ReactionTypeEmoji = &ReactionTypeEmoji{}
		return json.Unmarshal(data, rt.ReactionTypeEmoji)
	case ReactionTypeTypeCustomEmoji:
		rt.ReactionTypeCustomEmoji = &ReactionTypeCustomEmoji{}
		return json.Unmarshal(data, rt.ReactionTypeCustomEmoji)
	case ReactionTypeTypePaid:
		rt.ReactionTypePaid = &ReactionTypePaid{}
		return json.Unmarshal(data, rt.ReactionTypePaid)
	}

	return nil
}

// ReactionTypeEmoji https://core.telegram.org/bots/api#reactiontypeemoji
type ReactionTypeEmoji struct {
	Type  ReactionTypeType `json:"type"`
	Emoji string           `json:"emoji"`
}

// ReactionTypeCustomEmoji https://core.telegram.org/bots/api#reactiontypecustomemoji
type ReactionTypeCustomEmoji struct {
	Type          ReactionTypeType `json:"type"`
	CustomEmojiID string           `json:"custom_emoji_id"`
}

// ReactionTypePaid https://core.telegram.org/bots/api#reactiontypepaid
type ReactionTypePaid struct {
	Type string `json:"type"`
}

// MessageReactionUpdated https://core.telegram.org/bots/api#messagereactionupdated
type MessageReactionUpdated struct {
	Chat        Chat           `json:"chat"`
	MessageID   int            `json:"message_id"`
	User        *User          `json:"user,omitempty"`
	ActorChat   *Chat          `json:"actor_chat,omitempty"`
	Date        int            `json:"date"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
}

// ReactionCount https://core.telegram.org/bots/api#reactioncount
type ReactionCount struct {
	Type       ReactionType `json:"type"`
	TotalCount int          `json:"total_count"`
}

// MessageReactionCountUpdated https://core.telegram.org/bots/api#messagereactioncountupdated
type MessageReactionCountUpdated struct {
	Chat      Chat            `json:"chat"`
	MessageID int             `json:"message_id"`
	Date      int             `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}
