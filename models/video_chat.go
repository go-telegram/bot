package models

// See Telegram API docs: https://core.telegram.org/bots/api#videochatscheduled
type VideoChatScheduled struct {
	StartDate int `json:"start_date"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#videochatstarted
type VideoChatStarted struct {
	Duration int `json:"duration"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#videochatended
type VideoChatEnded struct {
	Duration int `json:"duration"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#videochatparticipantsinvited
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}
