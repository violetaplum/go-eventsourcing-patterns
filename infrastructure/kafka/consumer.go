package infraKafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go-eventsourcing-patterns/domain"
	"log"
)

type EventConsumer struct {
	consumer  *kafka.Consumer
	topic     string
	handlers  map[string]domain.EventHandler
	isRunning bool
}

func NewEventConsumer(brokers string, groupID string, topic string) (*EventConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":       brokers,
		"group.id":                groupID,
		"auto.offset.reset":       "earliest",
		"enable.auto.commit":      true,
		"auto.commit.interval.ms": 1000,
		"client.id":               "account-service-consumer",
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %v", err)
	}

	return &EventConsumer{
		consumer:  c,
		topic:     topic,
		handlers:  make(map[string]domain.EventHandler),
		isRunning: false,
	}, nil
}

func (ec *EventConsumer) RegisterHandler(eventType string, handler domain.EventHandler) {
	ec.handlers[eventType] = handler
}

func (ec *EventConsumer) Subscribe(ctx context.Context) error {
	if err := ec.consumer.SubscribeTopics([]string{ec.topic}, nil); err != nil {
		return fmt.Errorf("failed to subscribe to topic %s: %v", ec.topic, err)
	}

	ec.isRunning = true
	go ec.consumeMessages(ctx)
	return nil
}

func (ec *EventConsumer) consumeMessages(ctx context.Context) {
	for ec.isRunning {
		select {
		case <-ctx.Done(): // 컨텍스트 취소됐을때의 처리
			ec.isRunning = false
			return
		default:
			// 일반적 메세지 처리
			msg, err := ec.consumer.ReadMessage(100)
			if err != nil {
				if !err.(kafka.Error).IsTimeout() {
					log.Printf("Error reading message: %v", err)
				}
				continue
			}

			if err := ec.processMessage(ctx, msg); err != nil {
				log.Printf("Error processing message: %v", err)
			}

		}
	}
}

func (ec *EventConsumer) processMessage(ctx context.Context, msg *kafka.Message) error {
	var event domain.Event
	if err := json.Unmarshal(msg.Value, &event); err != nil {
		return fmt.Errorf("failed to unmarshal event: %v", err)
	}

	eventType := event.GetEventType()
	handler, exists := ec.handlers[eventType]
	if !exists {
		return fmt.Errorf("no handler registered for event type: %s", eventType)
	}

	if err := handler.Handle(ctx, event); err != nil {
		return fmt.Errorf("handler failed for event type: %v", eventType)
	}
	return nil
}

func (ec *EventConsumer) Close() error {
	ec.isRunning = false
	return ec.consumer.Close()
}
