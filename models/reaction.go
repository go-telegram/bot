package models

import (
	"encoding/json"
	"fmt"
)

// ReactionTypeType https://core.telegram.org/bots/api#reactiontype
type ReactionTypeType int

const (
	ReactionTypeTypeEmoji ReactionTypeType = iota
	ReactionTypeTypeCustomEmoji
)

// ReactionType https://core.telegram.org/bots/api#reactiontype
type ReactionType struct {
	Type ReactionTypeType

	ReactionTypeEmoji       *ReactionTypeEmoji
	ReactionTypeCustomEmoji *ReactionTypeCustomEmoji
}

func (rt *ReactionType) UnmarshalJSON(data []byte) error {
	v := struct {
		Type string `json:"type"`
	}{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	switch v.Type {
	case "emoji":
		rt.Type = ReactionTypeTypeEmoji
		rt.ReactionTypeEmoji = &ReactionTypeEmoji{}
		return json.Unmarshal(data, rt.ReactionTypeEmoji)
	case "custom_emoji":
		rt.Type = ReactionTypeTypeCustomEmoji
		rt.ReactionTypeCustomEmoji = &ReactionTypeCustomEmoji{}
		return json.Unmarshal(data, rt.ReactionTypeCustomEmoji)
	}

	return fmt.Errorf("unsupported ReactionType type")
}

// ReactionTypeEmoji https://core.telegram.org/bots/api#reactiontypeemoji
type ReactionTypeEmoji struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}

// ReactionTypeCustomEmoji https://core.telegram.org/bots/api#reactiontypecustomemoji
type ReactionTypeCustomEmoji struct {
	Type          string `json:"type"`
	CustomEmojiID string `json:"custom_emoji_id"`
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
