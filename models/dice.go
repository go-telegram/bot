package models

// See Telegram API docs: https://core.telegram.org/bots/api#dice
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}
