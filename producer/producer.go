package producer

import (
	"sync"

	"github.com/IBM/sarama"
)

type Producer struct {
	topic            string
	producersLock    sync.Mutex
	producers        []sarama.AsyncProducer
	producerProvider func() sarama.AsyncProducer
}

func NewProducer(config *sarama.Config, topic string, brokers []string) *Producer {
	return &Producer{
		topic: topic,
		producerProvider: func() sarama.AsyncProducer {
			producer, err := sarama.NewAsyncProducer(brokers, config)
			if err != nil {
				return nil
			}
			return producer
		},
	}
}

func (p *Producer) Borrow() (producer sarama.AsyncProducer) {
	p.producersLock.Lock()
	defer p.producersLock.Unlock()

	if len(p.producers) == 0 {
		for {
			producer = p.producerProvider()
			if producer != nil {
				return
			}
		}
	}

	index := len(p.producers) - 1
	producer = p.producers[index]
	p.producers = p.producers[:index]
	return
}

func (p *Producer) Release(producer sarama.AsyncProducer) {
	p.producersLock.Lock()
	defer p.producersLock.Unlock()

	if producer.TxnStatus()&sarama.ProducerTxnFlagInError != 0 {
		_ = producer.Close()
		return
	}
	p.producers = append(p.producers, producer)
}

func (p *Producer) Clear() {
	p.producersLock.Lock()
	defer p.producersLock.Unlock()

	for _, producer := range p.producers {
		producer.Close()
	}
	p.producers = p.producers[:0]
}

func (p *Producer) Produce(message string) {
	producer := p.Borrow()
	defer p.Release(producer)

	producer.Input() <- &sarama.ProducerMessage{Topic: p.topic, Key: nil, Value: sarama.StringEncoder(string(message))}
}
