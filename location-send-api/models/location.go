package models

type Location struct {
	DeviceId  string `json:"deviceId"`
	Coords    Coords `json:"coords"`
	Timestamp int64  `json:"timestamp"`
}
