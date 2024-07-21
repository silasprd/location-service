package websocket

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestWebsocketConnect(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := WebSocketConnect(w, r)
		if err != nil {
			t.Errorf("Failed to upgrade to WebSocket: %v", err)
			return
		}
		defer conn.Close()

		// Lê a mensagem do WebSocket
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			t.Errorf("Failed to read message: %v", err)
			return
		}

		// Escreve a mensagem de volta para o WebSocket
		if err := conn.WriteMessage(messageType, p); err != nil {
			t.Errorf("Failed to write message: %v", err)
			return
		}
	}))

	defer server.Close()

	time.Sleep(100 * time.Millisecond)

	// Define a URL do WebSocket
	u := "ws" + server.URL[len("http"):]

	// Cria um cliente WebSocket
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer ws.Close()

	// Envia uma mensagem pelo WebSocket
	message := []byte("test message")
	err = ws.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		t.Fatalf("Failed to write message: %v", err)
	}

	// Lê a mensagem de volta do WebSocket
	_, p, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}

	// Verifica se a mensagem recebida é igual à mensagem enviada
	assert.Equal(t, message, p)
}
