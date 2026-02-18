package models

// UserRating https://core.telegram.org/bots/api#userrating
type UserRating struct {
	Level              int `json:"level"`
	Rating             int `json:"rating"`
	CurrentLevelRating int `json:"current_level_rating"`
	NextLevelRating    int `json:"next_level_rating,omitempty"`
}
