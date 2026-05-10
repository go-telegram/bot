package models

import "encoding/json"

// PollAnswer https://core.telegram.org/bots/api#pollanswer
type PollAnswer struct {
	PollID              string   `json:"poll_id"`
	VoterChat           *Chat    `json:"voter_chat,omitempty"`
	User                *User    `json:"user,omitempty"`
	OptionIDs           []int    `json:"option_ids,omitempty"`
	OptionPersistentIDs []string `json:"option_persistent_ids,omitempty"`
}

// InputPollOption https://core.telegram.org/bots/api#inputpolloption
type InputPollOption struct {
	Text          string          `json:"text"`
	TextParseMode ParseMode       `json:"text_parse_mode,omitempty"`
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`
	Media         InputMedia      `json:"media,omitempty"`
}

func (o InputPollOption) MarshalJSON() ([]byte, error) {
	type base struct {
		Text          string          `json:"text"`
		TextParseMode ParseMode       `json:"text_parse_mode,omitempty"`
		TextEntities  []MessageEntity `json:"text_entities,omitempty"`
	}
	b := base{Text: o.Text, TextParseMode: o.TextParseMode, TextEntities: o.TextEntities}
	if o.Media == nil {
		return json.Marshal(b)
	}
	raw, err := o.Media.MarshalInputMedia()
	if err != nil {
		return nil, err
	}
	return json.Marshal(struct {
		base
		Media json.RawMessage `json:"media"`
	}{
		base:  b,
		Media: raw,
	})
}

// PollOption https://core.telegram.org/bots/api#polloption
type PollOption struct {
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	VoterCount   int             `json:"voter_count"`
	PersistentID string          `json:"persistent_id,omitempty"`
	Media        *PollMedia      `json:"media,omitempty"`
	AddedByUser  *User           `json:"added_by_user,omitempty"`
	AddedByChat  *Chat           `json:"added_by_chat,omitempty"`
	AdditionDate int             `json:"addition_date,omitempty"`
}

// PollMedia https://core.telegram.org/bots/api#pollmedia
type PollMedia struct {
	Animation *Animation  `json:"animation,omitempty"`
	Audio     *Audio      `json:"audio,omitempty"`
	Document  *Document   `json:"document,omitempty"`
	LivePhoto *LivePhoto  `json:"live_photo,omitempty"`
	Location  *Location   `json:"location,omitempty"`
	Photo     []PhotoSize `json:"photo,omitempty"`
	Sticker   *Sticker    `json:"sticker,omitempty"`
	Venue     *Venue      `json:"venue,omitempty"`
	Video     *Video      `json:"video,omitempty"`
}

// Poll https://core.telegram.org/bots/api#poll
type Poll struct {
	ID                    string          `json:"id"`
	Question              string          `json:"question"`
	QuestionEntities      []MessageEntity `json:"question_entities,omitempty"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int             `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOptionIDs      []int           `json:"correct_option_ids,omitempty"`
	Explanation           string          `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	Description           string          `json:"description,omitempty"`
	DescriptionEntities   []MessageEntity `json:"description_entities,omitempty"`
	Media                 *PollMedia      `json:"media,omitempty"`
	ExplanationMedia      *PollMedia      `json:"explanation_media,omitempty"`
	MembersOnly           bool            `json:"members_only,omitempty"`
	CountryCodes          []string        `json:"country_codes,omitempty"`
	AllowsRevoting        bool            `json:"allows_revoting,omitempty"`
	OpenPeriod            int             `json:"open_period,omitempty"`
	CloseDate             int             `json:"close_date,omitempty"`
}

// PollOptionAdded https://core.telegram.org/bots/api#polloptionadded
type PollOptionAdded struct {
	PollMessage        *MaybeInaccessibleMessage `json:"poll_message,omitempty"`
	OptionPersistentID string                    `json:"option_persistent_id"`
	OptionText         string                    `json:"option_text"`
	OptionTextEntities []MessageEntity           `json:"option_text_entities,omitempty"`
}

// PollOptionDeleted https://core.telegram.org/bots/api#polloptiondeleted
type PollOptionDeleted struct {
	PollMessage        *MaybeInaccessibleMessage `json:"poll_message,omitempty"`
	OptionPersistentID string                    `json:"option_persistent_id"`
	OptionText         string                    `json:"option_text"`
	OptionTextEntities []MessageEntity           `json:"option_text_entities,omitempty"`
}
