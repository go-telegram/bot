package models

// VideoChatScheduled https://core.telegram.org/bots/api#videochatscheduled
type VideoChatScheduled struct {
	StartDate int `json:"start_date"`
}

// VideoChatStarted https://core.telegram.org/bots/api#videochatstarted
type VideoChatStarted struct {
	Duration int `json:"duration"`
}

// VideoChatEnded https://core.telegram.org/bots/api#videochatended
type VideoChatEnded struct {
	Duration int `json:"duration"`
}

// VideoChatParticipantsInvited https://core.telegram.org/bots/api#videochatparticipantsinvited
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}
