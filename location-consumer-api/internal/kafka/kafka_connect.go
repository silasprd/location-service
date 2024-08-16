package kafka

import (
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var reader *kafka.Reader

func InitConsumer(brokerURL, topic string) {
	for {
		reader = kafka.NewReader(kafka.ReaderConfig{
			Brokers:     []string{brokerURL},
			Topic:       topic,
			GroupID:     "location-processor-group",
			StartOffset: kafka.LastOffset,
		})
		log.Println("Kafka consumer initialized")
		break
	}
	log.Println("Failed to initialize Kafka, retrying in 15 seconds...")
	time.Sleep(15 * time.Second)
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
