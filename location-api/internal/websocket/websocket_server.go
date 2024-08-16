package websocket

import (
	"crypto/tls"
	"log"
	"net/http"
)

func StartServer(address, crtSSL, keySSL string) {
	http.HandleFunc("/ws", HandleMessage)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13,
	}

	server := &http.Server{
		Addr:      address,
		TLSConfig: tlsConfig,
	}

	log.Printf("WebSocket server starting at %s", address)
	log.Fatal(server.ListenAndServeTLS(crtSSL, keySSL))
}
