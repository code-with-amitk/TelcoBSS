package kafka

import (
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

// Producer wraps a Sarama async producer.
type Producer struct {
	producer sarama.AsyncProducer
}

func NewProducer(brokers []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Timeout = 10 * time.Second

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("create async producer: %w", err)
	}

	go func() {
		for {
			select {
			case success := <-producer.Successes():
				logrus.Infof("message sent to topic %s partition %d", success.Topic, success.Partition)
			case err := <-producer.Errors():
				logrus.Errorf("kafka producer error: %v", err)
			}
		}
	}()

	return &Producer{producer: producer}, nil
}

func (p *Producer) Publish(topic string, key string, value []byte) error {
	message := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(value),
	}
	p.producer.Input() <- message
	return nil
}

func (p *Producer) Close() error {
	return p.producer.Close()
}
