package kafka

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

type KafkaEventPublisher struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaEventPublisher(brokers string, topic string) (*KafkaEventPublisher, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"client.id":         "account-service-producer",
		"acks":              "all", // 모든 ISR 에 복제 되었는지 확인
	})
}
