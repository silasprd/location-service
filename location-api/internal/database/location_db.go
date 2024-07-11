package database

import (
	"github.com/silasprd/sailor-location-service/location-api/internal/entity"
	"gorm.io/gorm"
)

type Location struct {
	DB *gorm.DB
}

func (l *Location) Upsert(location *entity.Location) error {
	err := l.DB.Where(entity.Location{DeviceId: location.DeviceId}).Assign(location).FirstOrCreate(location).Error
	if err != nil {
		return err
	}
	return nil
}
