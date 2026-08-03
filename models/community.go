package models

// Community https://core.telegram.org/bots/api#community
//
// Represents a community (a group of chats).
type Community struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// CommunityChatAdded https://core.telegram.org/bots/api#communitychatadded
//
// Describes a service message about a chat being added to a community.
type CommunityChatAdded struct {
	Community Community `json:"community"`
}

// CommunityChatRemoved https://core.telegram.org/bots/api#communitychatremoved
//
// Describes a service message about a chat being removed from a community.
// Currently holds no information.
type CommunityChatRemoved struct{}
