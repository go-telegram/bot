package models

// ProximityAlertTriggered https://core.telegram.org/bots/api#proximityalerttriggered
type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"`
	Watcher  User `json:"watcher"`
	Distance int  `json:"distance"`
}
