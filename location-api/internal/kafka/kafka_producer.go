package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func SendMessage(key, message []byte) error {
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   key,
			Value: message,
		},
	)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}

	log.Printf("Key: %s ", key)
	log.Printf("Message sent: %s", message)

	return err
}
