package websocket

import (
	"log"
	"net/http"
)

func StartServer(address string) {
	http.HandleFunc("/ws", SendMessage)
	log.Printf("WebSocket server starting at %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
