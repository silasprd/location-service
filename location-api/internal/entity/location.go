package entity

type Location struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	DeviceId  string  `gorm:"deviceId"`
	Latitude  float64 `gorm:"latitude"`
	Longitude float64 `gorm:"longitude"`
	Speed     float64 `gorm:"speed"`
	Heading   float64 `gorm:"heading"`
	Accuracy  float64 `gorm:"accuracy"`
	Altitude  float64 `gorm:"altitude"`
	Timestamp int64   `gorm:"timestamp"`
}
