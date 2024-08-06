package kafka

import (
	"log"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitProducer(kafkaBrokerURL string, topic string) {

	go func() {
		for {
			err := CreateTopic(topic, 1, 1)
			if err != nil {
				log.Fatalf("Failed to create Kafka topic: %v", err)
			}

			writer = kafka.NewWriter(kafka.WriterConfig{
				Brokers:  []string{kafkaBrokerURL},
				Topic:    topic,
				Balancer: &kafka.LeastBytes{},
			})
			log.Println("Kafka producer initialized!")
			break
		}
	}()
}

func CloseProducer() {
	if writer != nil {
		writer.Close()
		log.Println("Kafka writer closed")
	}
}
