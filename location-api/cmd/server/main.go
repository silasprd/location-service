package main

import (
	"github.com/silasprd/sailor-location-service/location-api/internal/kafka"
	"github.com/silasprd/sailor-location-service/location-api/internal/websocket"
)

func main() {

	kafkaBrokerURL := "localhost:9092"
	topic := "locations"

	kafka.InitProducer(kafkaBrokerURL, topic)
	defer kafka.CloseProducer()

	websocket.StartServer(":8080")

	select {}

}
