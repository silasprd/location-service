package websocket

import (
	"log"
	"net/http"

	"github.com/silasprd/sailor-location-service/location-api/internal/kafka"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {

	conn, err := WebSocketConnect(w, r)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		log.Printf("Received message: %s", message)

		// Enviar mensagem ao Kafka
		if err := kafka.SendMessage(nil, message); err != nil {
			log.Printf("Error sending message to Kafka: %v", err)
		}
	}
}
