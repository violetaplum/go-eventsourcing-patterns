package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go-eventsourcing-patterns/domain"
)

type EventConsumer struct {
	consumer *kafka.Consumer
	topic    string
	handler  domain.EventHandler
}

func NewEventConsumer(brokers string, groupID string, topic string, handler domain.EventHandler) (*EventConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":       brokers,
		"group.id":                groupID,
		"auto.offset.reset":       "earliest",
		"enable.auto.commit":      true,
		"auto.commit.interval.ms": 1000,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %v", err)
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to topic: %v", err)
	}
	return &EventConsumer{
		consumer: c,
		topic:    topic,
		handler:  handler,
	}, nil
}
