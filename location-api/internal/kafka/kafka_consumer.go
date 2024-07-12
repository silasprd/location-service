package kafka

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/segmentio/kafka-go"
	"github.com/silasprd/sailor-location-service/location-api/internal/database"
	"github.com/silasprd/sailor-location-service/location-api/internal/entity"
	"gorm.io/gorm"
)

var consumers = make(map[string]*websocket.Conn)
var mu sync.Mutex

func GetMessages(brokerURL, topic string, db *gorm.DB) {

	log.Printf("GETMESSAGES CALL")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerURL},
		Topic:   topic,
	})

	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		var location entity.Location
		if err := json.Unmarshal(msg.Value, &location); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		locationDB := database.Location{DB: db}
		if err := locationDB.Upsert(&location); err != nil {
			log.Printf("Error upserting location: %v", err)
		}

		broadcastMessage(msg.Value)
	}

}

func RegisterConsumer(id string, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	consumers[id] = conn
}

func UnregisterConsumer(id string) {
	mu.Lock()
	defer mu.Unlock()
	delete(consumers, id)
}

func broadcastMessage(message []byte) {
	mu.Lock()
	defer mu.Unlock()
	for _, conn := range consumers {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}
