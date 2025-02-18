package main

import (
	"context"
	"go-eventsourcing-patterns/domain"
	infraKafka "go-eventsourcing-patterns/infrastructure/kafka"
	store "go-eventsourcing-patterns/infrastructure/persistence/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	db, err := store.NewPostgresDB(&domain.Config{
		DBHost:     "postgres", // docker 서비스명
		DBPort:     "5432",
		DBUser:     "user",
		DBPassword: "password",
		DBName:     "eventstore",
		SSLMode:    "disable", // 로컬 개발환경이므로 SSL 비활성화
	})
	if err != nil {
		log.Fatalf("Failed to conncting databse: %v", err)
	}

	eventStore := store.NewEventStore(db)
	brokers := os.Getenv("KAFKA_BROKERS")
	topic := os.Getenv("KAFKA_TOPIC")

	if brokers == "" || topic == "" {
		log.Fatalf("Empty kafka info")
	}

	groupId := os.Getenv("KAFKA_GROUP_ID")
	// 환경변수 로깅 추가
	log.Printf("Kafka configuration - Brokers: %s, Topic: %s, GroupID: %s", brokers, topic, groupId)

	// 순서 확인 (brokers, groupId, topic)
	consumer, err := infraKafka.NewEventConsumer(brokers, groupId, topic)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}

	// Subscribe 전에 토픽 로깅
	log.Printf("Attempting to subscribe to topic: %s", topic)
	defer consumer.Close()

	accountCreatedHandler := infraKafka.NewAccountCreatedHandler(eventStore)
	moneyDepositedHandler := infraKafka.NewMoneyDepositHandler(eventStore)
	moneyWithdrawnHandler := infraKafka.NewMoneyWithdrawHandler(eventStore)

	consumer.RegisterHandler(string(domain.AccountCreated), accountCreatedHandler)
	consumer.RegisterHandler(string(domain.MoneyDeposited), moneyDepositedHandler)
	consumer.RegisterHandler(string(domain.MoneyWithdrawn), moneyWithdrawnHandler)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	if err := consumer.Subscribe(ctx); err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	log.Println("Event consumer started successfully")

	//시그널 대기
	<-sigChan
	log.Println("Shutting down...")
}
