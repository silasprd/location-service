package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/segmentio/kafka-go"

	models "location-send-api/models"
)

// Configuração do Kafka
const (
	kafkaBrokerURL = "localhost:9092"
	topic          = "locations"
)

// Configuração do WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Permitir conexões de qualquer origem
	},
}

// Kafka writer (producer)
var writer *kafka.Writer

func main() {
	// Criar o tópico no Kafka, se necessário
	err := createTopic(topic, 1, 1)
	if err != nil {
		log.Fatalf("Failed to create Kafka topic: %v", err)
	}

	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaBrokerURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	http.HandleFunc("/ws", handleLocation)

	fmt.Println("Server WebSocket running in port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createTopic(topic string, numPartitions int, replicationFactor int) error {
	conn, err := kafka.Dial("tcp", kafkaBrokerURL)
	if err != nil {
		return fmt.Errorf("Failed to connect to Kafka broker: %v", err)
	}
	defer conn.Close()

	// Listar os tópicos existentes
	partitions, err := conn.ReadPartitions()
	if err != nil {
		return fmt.Errorf("Failed to read partitions: %v", err)
	}

	for _, p := range partitions {
		if p.Topic == topic {
			log.Printf("Tópico %s exists", topic)
			return nil
		}
	}

	// Criar o tópico se não existir
	controller, err := conn.Controller()
	if err != nil {
		return fmt.Errorf("Failed to get Kafka controller: %v", err)
	}

	controllerConn, err := kafka.Dial("tcp", fmt.Sprintf("%s:%d", controller.Host, controller.Port))
	if err != nil {
		return fmt.Errorf("Failed to connect to Kafka controller: %v", err)
	}
	defer controllerConn.Close()

	err = controllerConn.CreateTopics(
		kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     numPartitions,
			ReplicationFactor: replicationFactor,
		},
	)
	if err != nil {
		return fmt.Errorf("Failed to create Kafka topic: %v", err)
	}

	log.Printf("Topic %s created", topic)

	return nil
}

func handleLocation(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erro ao atualizar para WebSocket:", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected")

	for {
		// Lê a mensagem do WebSocket
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error to read message:", err)
			if websocket.IsCloseError(err, websocket.CloseGoingAway) {
				log.Println("WebSocket closing: Client disconnected")
			} else if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket unexpected close error: %v", err)
			}
			break
		}

		log.Printf("Message send: %s", message)

		// Deserializar a mensagem para extrair o deviceId
		var location models.Location
		if err := json.Unmarshal(message, &location); err != nil {
			log.Println("Error parsing message:", err)
			break
		}

		// Envia a mensagem para o Kafka
		kafkaMessage := kafka.Message{
			Key:   []byte(location.DeviceId), // Timestamp como chave
			Value: message,
		}

		err = writer.WriteMessages(r.Context(), kafkaMessage)
		if err != nil {
			log.Println("Error to send message for Kafka:", err)
			break
		}

		log.Println("Success to send message for Kafka")
	}

	log.Println("Client disconnected")

}
