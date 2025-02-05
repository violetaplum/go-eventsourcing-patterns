package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go-eventsourcing-patterns/domain"
	"log"
)

type EventPublisher struct {
	producer *kafka.Producer
	topic    string
}

func NewEventPublisher(brokers string, topic string) (*EventPublisher, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"client.id":         "account-service-producer",
		"acks":              "all", // 모든 ISR 에 복제 되었는지 확인
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %v", err)
	}

	return &EventPublisher{
		producer: p,
		topic:    topic,
	}, nil
}

func (ep *EventPublisher) Publish(ctx context.Context, event domain.Event) error {
	jsonEvent, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %v", err)
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &ep.topic,
			Partition: kafka.PartitionAny, // 카프카가 적절한 파티션 선택
		},
		Key:   []byte(event.GetAggregateID()), // 집계 ID 를 키로 사용
		Value: jsonEvent,
	}

	// 메시지 전송 및 전달 확인 채널 생성
	deliveryChan := make(chan kafka.Event, 1)
	err = ep.producer.Produce(msg, deliveryChan)
	if err != nil {
		return fmt.Errorf("error queuing message: %v", err)
	}

	// 메시지 전달 결과 확인
	e := <-deliveryChan
	switch ev := e.(type) {
	case *kafka.Message:
		if ev.TopicPartition.Error != nil {
			return fmt.Errorf("message delivery failed: %v", ev.TopicPartition.Error)
		}

		log.Printf("Event published: AccountID=%s, Type=%s",
			event.GetAggregateID(),
			event.GetEventType())
	}
	return nil
}

func (kp EventPublisher) PublishAll(ctx context.Context, events []domain.Event) error {
	for _, event := range events {
		if err := kp.Publish(ctx, event); err != nil {
			return err
		}
	}
	return nil
}

// Close Kafka 프로듀서 종료
func (kp *EventPublisher) Close() {
	// Flush는 아직 전송되지 않은 메시지가 있다면 모두 전송
	kp.producer.Flush(10 * 1000)
	kp.producer.Close()
}
