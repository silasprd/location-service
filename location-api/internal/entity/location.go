package entity

import "errors"

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

func (location *Location) validate() error {

	if location.DeviceId == "" {
		return errors.New("deviceId cannot be empty")
	}

	if location.Latitude < -90 || location.Latitude > 90 {
		return errors.New("latitude must be between -90 and 90")
	}

	if location.Longitude < -180 || location.Longitude > 180 {
		return errors.New("longitude must be between -180 and 180")
	}

	if location.Accuracy < 0 {
		return errors.New("accuracy cannot be negative")
	}

	if location.Speed < 0 {
		return errors.New("speed cannot be negative")
	}

	if location.Heading < 0 || location.Heading >= 360 {
		return errors.New("heading must be between 0 and 360 degrees")
	}

	if location.Altitude < 0 {
		return errors.New("altitude cannot be negative")
	}

	if location.Timestamp <= 0 {
		return errors.New("timestamp must be positive")
	}

	return nil
}
