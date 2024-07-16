package kafka

import (
	"context"
	"encoding/json"
	"log"

	database "github.com/silasprd/sailor-location-service/location-consumer-api/internal/database/connection"
	location_db "github.com/silasprd/sailor-location-service/location-consumer-api/internal/database/location"
	"github.com/silasprd/sailor-location-service/location-consumer-api/internal/entity"
	"github.com/silasprd/sailor-location-service/location-consumer-api/internal/websocket"
)

// var consumers = make(map[string]*websocket.Conn)
// var mu sync.Mutex

func GetMessages() {

	reader := GetReader()
	if reader == nil {
		log.Println("Kafka reader is nil")
		return
	}

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		log.Printf("Received message: %s", msg.Value)

		var location entity.Location
		if err := json.Unmarshal(msg.Value, &location); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		locationDB := location_db.Location{DB: database.DB}
		if err := locationDB.Upsert(&location); err != nil {
			log.Printf("Error upserting location: %v", err)
		}

		websocket.BroadcastMessage(msg.Value)

	}

}
