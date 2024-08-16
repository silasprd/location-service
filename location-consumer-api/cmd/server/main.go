package main

import (
	"log"
	"os"
	"path/filepath"

	database "github.com/silasprd/sailor-location-service/location-consumer-api/internal/database/connection"
	"github.com/silasprd/sailor-location-service/location-consumer-api/internal/kafka"
	"github.com/silasprd/sailor-location-service/location-consumer-api/internal/websocket"
)

func main() {
	kafkaBrokerURL := "localhost:9092"
	topic := "locations"

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Erro ao obter diretório de trabalho: %v", err)
	}

	crtPath := filepath.Join(cwd, "/internal/ssl/locationSSL.crt")
	keyPath := filepath.Join(cwd, "/internal/ssl/locationSSL.key")

	database.MySQLInitDB()

	database.MongoInitDB()

	kafka.InitConsumer(kafkaBrokerURL, topic)

	go websocket.StartServer(":8082", crtPath, keyPath)

	go kafka.GetMessages()

	select {}
}
