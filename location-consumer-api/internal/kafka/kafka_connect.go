package kafka

import (
	"log"

	"github.com/segmentio/kafka-go"
)

var reader *kafka.Reader

func InitConsumer(brokerURL, topic string) {
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{brokerURL},
		Topic:       topic,
		GroupID:     "location-processor-group",
		StartOffset: kafka.LastOffset,
	})
	log.Println("Kafka consumer initialized")
}

func CloseConsumer() {
	if reader != nil {
		reader.Close()
		log.Println("Kafka consumer closed")
	}
}

func GetReader() *kafka.Reader {
	return reader
}
