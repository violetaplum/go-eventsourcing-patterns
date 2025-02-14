package main

import (
	"go-eventsourcing-patterns/domain"
	infraKafka "go-eventsourcing-patterns/infrastructure/kafka"
	"log"
	"os"
)

func main() {
	//
	//db, err := store.NewPostgresDB(&domain.Config{
	//	DBHost:     "postgres", // docker 서비스명
	//	DBPort:     "5432",
	//	DBUser:     "user",
	//	DBPassword: "password",
	//	DBName:     "eventstore",
	//	SSLMode:    "disable", // 로컬 개발환경이므로 SSL 비활성화
	//})
	//if err != nil {
	//	log.Fatalf("Failed to conncting databse: %v", err)
	//}
	//
	//eventStore := store.NewEventStore(db)
	brokers := os.Getenv("KAFKA_BROKERS")
	topic := os.Getenv("KAFKA_TOPIC")

	if brokers == "" || topic == "" {
		log.Fatalf("Empty kafka info")
	}

	groupId := os.Getenv("KAFKA_GROUP_ID")
	consumer, err := infraKafka.NewEventConsumer(brokers, topic, groupId)
	if err != nil {
		log.Fatalf("error creating event consumer: %v", err)
	}
	defer consumer.Close()

	err = consumer.Subscribe(func(event domain.Event) error {
		log.Printf("Received event: %+v", event)
		return nil
	})

	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}
}
