package main

import (
	"github.com/gin-gonic/gin"
	appCommand "go-eventsourcing-patterns/application/command"
	"go-eventsourcing-patterns/application/query"
	"go-eventsourcing-patterns/domain"
	infraKafka "go-eventsourcing-patterns/infrastructure/kafka"
	store "go-eventsourcing-patterns/infrastructure/persistence/postgres"
	"go-eventsourcing-patterns/interface/http"
	"log"
	"os"
)

func main() {
	router := gin.Default()

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

	accountStore := store.NewAccountStore(db)

	brokers := os.Getenv("KAFKA_BROKERS")
	topic := os.Getenv("KAFKA_TOPIC")

	if brokers == "" || topic == "" {
		log.Fatalf("Empty kafka info")
	}

	eventPublisher, err := infraKafka.NewEventPublisher(brokers, topic)
	if err != nil {
		log.Fatalf("Failed to create event publisher: %v", err)
	}

	eventStore := store.NewEventStore(db)
	commandService := appCommand.NewAccountCommandService(accountStore, eventStore, eventPublisher, db)
	queryService := query.NewAccountQueryService(accountStore, eventStore)

	accountHandler := http.NewAccountHandler(commandService, queryService)
	accountHandler.SetupRoutes(router)

	router.Run(":8080")

}
