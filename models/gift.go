package models

// Gifts https://core.telegram.org/bots/api#gifts
type Gifts struct {
	Gifts []Gift `json:"gifts"`
}

// Gift https://core.telegram.org/bots/api#gift
type Gift struct {
	ID             string  `json:"id"`
	Sticker        Sticker `json:"sticker"`
	StarCount      int     `json:"star_count"`
	TotalCount     int     `json:"total_count,omitempty"`
	RemainingCount int     `json:"remaining_count,omitempty"`
}
