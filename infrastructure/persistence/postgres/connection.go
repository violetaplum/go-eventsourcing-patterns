package postgres

import (
	"fmt"
	"go-eventsourcing-patterns/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection(config *domain.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{}) // db, error 둘다 반환
}
