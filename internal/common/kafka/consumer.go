package kafka

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

// ConsumerHandler provides a simple consumer group wrapper for Go consumers.
type ConsumerHandler struct {
	ready chan bool
	fn    func(message *sarama.ConsumerMessage) error
}

func NewConsumerHandler(fn func(message *sarama.ConsumerMessage) error) *ConsumerHandler {
	return &ConsumerHandler{ready: make(chan bool), fn: fn}
}

func (h *ConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error {
	close(h.ready)
	return nil
}

func (h *ConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if err := h.fn(message); err != nil {
			logrus.Errorf("consumer handler failed: %v", err)
			continue
		}
		session.MarkMessage(message, "")
	}
	return nil
}

func StartConsumerGroup(ctx context.Context, brokers []string, groupID string, topics []string, handler *ConsumerHandler) error {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return fmt.Errorf("create consumer group: %w", err)
	}

	go func() {
		for {
			if err := client.Consume(ctx, topics, handler); err != nil {
				logrus.Errorf("consumer group error: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			handler.ready = make(chan bool)
		}
	}()

	<-handler.ready
	return nil
}
