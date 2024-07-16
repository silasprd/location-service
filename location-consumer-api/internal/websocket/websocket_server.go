package websocket

import (
	"log"
	"net/http"
)

func StartServer(port string) {
	http.HandleFunc("/ws", HandleMessage)
	go handleMessages()
	log.Printf("WebSocket server starting at %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
