package models

// See Telegram API docs: https://core.telegram.org/bots/api#videochatscheduled
type VoiceChatScheduled struct {
	StartDate int `json:"start_date"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#videochatstarted
type VoiceChatStarted struct {
	Duration int `json:"duration"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#videochatended
type VoiceChatEnded struct {
	Duration int `json:"duration"`
}

// See Telegram API docs: https://core.telegram.org/bots/api#videochatparticipantsinvited
type VoiceChatParticipantsInvited struct {
	Users []User `json:"users"`
}
