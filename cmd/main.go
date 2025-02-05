package cmd

import (

	"github.com/gin-gonic/gin"
	"go-eventsourcing-patterns/application/command"
	"go-eventsourcing-patterns/application/query"
	"go-eventsourcing-patterns/domain"
	store "go-eventsourcing-patterns/infrastructure/persistence/postgres"
	"go-eventsourcing-patterns/interface/http"
	"log"
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
	eventStore := store.NewEventStore(db)

	eventPublisher := 

	//eventPublisher domain.EventPublisher,
	//	txManager domain.TransactionManager,
	commandService := command.NewAccountCommandService(accountStore, eventStore)
	queryService := query.NewAccountQueryService(accountStore, eventStore)

	accountHandler := http.NewAccountHandler(commandService, queryService)
	accountHandler.SetupRoutes(router)

	router.Run(":8080")

}
