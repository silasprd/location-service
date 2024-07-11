package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func SendMessage(key, message []byte) error {
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   key,
			Value: message,
		},
	)
	return err
}
