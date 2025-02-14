package infraKafka

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go-eventsourcing-patterns/domain"
	"log"
)

type EventConsumer struct {
	consumer *kafka.Consumer
	topic    string
	//eventStore domain.EventStore
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

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to topic: %v", err)
	}
	return &EventConsumer{
		consumer: c,
		topic:    topic,
		//eventStore: eventStore,
	}, nil
}

// Subscribe 메서드 추가 - 이벤트 처리를 위한 핸들러 함수를 받음
func (ec *EventConsumer) Subscribe(handler func(event domain.Event) error) error {
	err := ec.consumer.Subscribe(ec.topic, nil)
	if err != nil {
		return fmt.Errorf("failed to subscribe to topic %s: %v", ec.topic, err)
	}

	// 이벤트 처리 루프
	for {
		msg, err := ec.consumer.ReadMessage(-1) // -1은 무한 대기
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		var event domain.Event
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("Error unmarshaling event: %v", err)
			continue
		}

		if err := handler(event); err != nil {
			log.Printf("Error handling event: %v", err)
			// 에러 처리 정책에 따라 추가 로직 구현 가능
			// 예: Dead Letter Queue로 보내기, 재시도 등
		}
	}
}

func (ec *EventConsumer) Close() error {
	return ec.consumer.Close()
}

//func (e *EventConsumer) ProcessEvent(ctx context.Context, event domain.Event) error {
//	switch event.GetEventType() {
//	case string(domain.AccountCreated):
//
//	case string(domain.MoneyDeposited):
//	case string(domain.MoneyWithdrawn):
//	default:
//		return fmt.Errorf("unknown event type: %v", event)
//	}
//
//	return nil
//}
