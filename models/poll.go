package models

// PollAnswer https://core.telegram.org/bots/api#pollanswer
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	VoterChat *Chat  `json:"voter_chat,omitempty"`
	User      *User  `json:"user"`
	OptionIDs []int  `json:"option_ids,omitempty"`
}

// InputPollOption https://core.telegram.org/bots/api#inputpolloption
type InputPollOption struct {
	Text          string          `json:"text"`
	TextParseMode ParseMode       `json:"text_parse_mode,omitempty"`
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`
}

// PollOption https://core.telegram.org/bots/api#polloption
type PollOption struct {
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	VoterCount   int             `json:"voter_count"`
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
	CorrectOptionID       int             `json:"correct_option_id,omitempty"`
	Explanation           string          `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int             `json:"open_period,omitempty"`
	CloseDate             int             `json:"close_date,omitempty"`
}
