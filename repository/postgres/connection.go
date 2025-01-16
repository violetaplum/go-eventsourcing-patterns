package postgres

import (
	"fmt"
	"go-eventsourcing-patterns/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewConnection(cfg domain.Config) (*gorm.DB, error) {
	// 데이터 소스 이름(DSN) 생성
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	// db 연결
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// 연결 성공
	log.Println("Connected to the database successfully")
	return db, nil
}

// fixme: gorm 을 사용하기때문에 굳이 필요없는 코드이긴 함
func CloseConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("error closing database connection: %v", err)
		return err
	}

	log.Println("Database connection closed")
	return nil
}
