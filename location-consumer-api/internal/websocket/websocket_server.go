package websocket

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/silasprd/sailor-location-service/location-consumer-api/internal/database/location"
)

func StartServer(port, crtSSL, keySSL string) {
	http.HandleFunc("/locations", location.FindAllHandler)
	http.HandleFunc("/ws", HandleMessage)
	go handleMessages()

	cors := corsMiddleware(http.DefaultServeMux)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13,
	}

	server := &http.Server{
		Addr:      port,
		Handler:   cors,
		TLSConfig: tlsConfig,
	}

	log.Printf("WebSocket server starting at %s", port)
	log.Fatal(server.ListenAndServeTLS(crtSSL, keySSL))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Defina os cabeçalhos CORS
		w.Header().Set("Access-Control-Allow-Origin", "https://192.168.68.102:4200")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Access-Control-Allow-Headers, Access-Control-Allow-Methods")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Responda imediatamente a solicitações OPTIONS
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Passe para o próximo handler
		next.ServeHTTP(w, r)
	})
}
