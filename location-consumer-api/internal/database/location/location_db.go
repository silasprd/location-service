package database

import (
	"context"
	"log"
	"time"

	"github.com/silasprd/sailor-location-service/location-consumer-api/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Location struct {
	DB *gorm.DB
}

type PastLocation struct {
	MongoDB *mongo.Database
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

func (m *PastLocation) Save(location *entity.Location) error {
	collection := m.MongoDB.Collection("locations")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, bson.M{
		"deviceId":  location.DeviceId,
		"latitude":  location.Latitude,
		"longitude": location.Longitude,
		"speed":     location.Speed,
		"heading":   location.Heading,
		"accuracy":  location.Accuracy,
		"altitude":  location.Altitude,
		"timestamp": location.Timestamp,
	})
	if err != nil {
		log.Printf("Failed to insert location: %v", err)
		return err
	}

	log.Println("Location saved to MongoDB")
	return nil
}
