package models

// VoiceChatScheduled https://core.telegram.org/bots/api#videochatscheduled
type VoiceChatScheduled struct {
	StartDate int `json:"start_date"`
}

// VoiceChatStarted https://core.telegram.org/bots/api#videochatstarted
type VoiceChatStarted struct {
	Duration int `json:"duration"`
}

// VoiceChatEnded https://core.telegram.org/bots/api#videochatended
type VoiceChatEnded struct {
	Duration int `json:"duration"`
}

// VoiceChatParticipantsInvited https://core.telegram.org/bots/api#videochatparticipantsinvited
type VoiceChatParticipantsInvited struct {
	Users []User `json:"users"`
}
