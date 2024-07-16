package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

func HandleMessage(w http.ResponseWriter, r *http.Request) {
	conn, err := WebSocketConnect(w, r)
	if err != nil {
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			delete(clients, conn)
			break
		}

		log.Printf("Received message: %s", message)
	}
}

func BroadcastMessage(message []byte) {
	broadcast <- message
}

func handleMessages() {
	for {
		message := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("Error writing message: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
