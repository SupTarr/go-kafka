package main

import (
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/SupTarr/go-kafka/consumer"
	"github.com/SupTarr/go-kafka/producer"
)

func consume() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := []string{"localhost:9092"}
	topics := []string{"topic"}
	group := "group"
	consumer := consumer.NewConsumer(config, brokers, topics, group)
	consumer.Consume()
}

func main() {
	go func() {
		consume()
	}()
	time.Sleep(2 * time.Second)

	i := 0
	config := sarama.NewConfig()
	config.Producer.Idempotent = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Net.MaxOpenRequests = 1
	brokers := []string{"localhost:9092"}
	topic := "topic"
	producer := producer.NewProducer(config, topic, brokers)
	for {
		producer.Produce(fmt.Sprintf("Message: %d", i))
		i += 1
		time.Sleep(1 * time.Second)
	}
}
