package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/silasprd/sailor-location-service/location-api/internal/kafka"
	"github.com/silasprd/sailor-location-service/location-api/internal/websocket"
)

func main() {

	kafkaBrokerURL := "localhost:9092"
	topic := "locations"

	kafka.InitProducer(kafkaBrokerURL, topic)
	defer kafka.CloseProducer()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Erro ao obter diretório de trabalho: %v", err)
	}

	certPath := filepath.Join(cwd, "/internal/ssl/locationSSL.crt")
	keyPath := filepath.Join(cwd, "/internal/ssl/locationSSL.key")

	websocket.StartServer(":8080", certPath, keyPath)

	select {}

}
