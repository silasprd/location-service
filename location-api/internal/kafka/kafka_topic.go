package kafka

import (
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

const kafkaBrokerURL = "localhost:9092"

func CreateTopic(topic string, numPartitions int, replicationFactor int) error {
	conn, err := kafka.Dial("tcp", kafkaBrokerURL)
	if err != nil {
		return fmt.Errorf("failed to connect to Kafka broker: %v", err)
	}
	defer conn.Close()

	// Listar os tópicos existentes
	partitions, err := conn.ReadPartitions()
	if err != nil {
		return fmt.Errorf("failed to read partitions: %v", err)
	}

	for _, p := range partitions {
		if p.Topic == topic {
			log.Printf("Topic %s exists", topic)
			return nil
		}
	}

	// Criar o tópico se não existir
	controller, err := conn.Controller()
	if err != nil {
		return fmt.Errorf("failed to get Kafka controller: %v", err)
	}

	controllerConn, err := kafka.Dial("tcp", fmt.Sprintf("%s:%d", controller.Host, controller.Port))
	if err != nil {
		return fmt.Errorf("failed to connect to Kafka controller: %v", err)
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
		return fmt.Errorf("failed to create Kafka topic: %v", err)
	}

	log.Printf("Topic %s created", topic)

	return nil
}
