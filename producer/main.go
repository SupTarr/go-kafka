package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(brokers, nil)
	if err != nil {
		log.Panicf("Error creating producer: %v", err)
	}
	defer producer.Close()

	msg := sarama.ProducerMessage{
		Topic: "suptarr",
		Value: sarama.StringEncoder("Hello, Sarama!"),
	}

	p, o, err := producer.SendMessage(&msg)
	if err != nil {
		log.Panicf("Error send message from producer: %v", err)
	}
	log.Printf("Message sent to partition %d at offset %d\n", p, o)
}
