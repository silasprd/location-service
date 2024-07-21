package entity

type Location struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	DeviceId  string  `gorm:"deviceId" json:"deviceId"`
	Latitude  float64 `gorm:"latitude" json:"latitude"`
	Longitude float64 `gorm:"longitude" json:"longitude"`
	Speed     float64 `gorm:"speed" json:"speed"`
	Heading   float64 `gorm:"heading" json:"heading"`
	Accuracy  float64 `gorm:"accuracy" json:"accuracy"`
	Altitude  float64 `gorm:"altitude" json:"altitude"`
	Timestamp int64   `gorm:"timestamp" json:"timestamp"`
}
