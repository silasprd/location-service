package database

import (
	"github.com/silasprd/sailor-location-service/location-consumer-api/internal/entity"
	"gorm.io/gorm"
)

type Location struct {
	DB *gorm.DB
}

func (l *Location) Upsert(location *entity.Location) error {
	var existingLocation entity.Location
	result := l.DB.Where(entity.Location{DeviceId: location.DeviceId}).First(&existingLocation)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}

	if result.Error == gorm.ErrRecordNotFound {
		result = l.DB.Create(location)
	} else {
		result = l.DB.Model(&existingLocation).Updates(location)
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
