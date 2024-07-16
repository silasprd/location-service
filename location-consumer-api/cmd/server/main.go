package main

import (
	database "github.com/silasprd/sailor-location-service/location-consumer-api/internal/database/connection"
	"github.com/silasprd/sailor-location-service/location-consumer-api/internal/kafka"
	"github.com/silasprd/sailor-location-service/location-consumer-api/internal/websocket"
)

func main() {
	kafkaBrokerURL := "localhost:9092"
	topic := "locations"

	database.MySQLInitDB()

	kafka.InitConsumer(kafkaBrokerURL, topic)

	go websocket.StartServer(":8081")

	go kafka.GetMessages()

	select {}
}
