package postgres

import (
	"context"
	"fmt"
	"go-eventsourcing-patterns/domain"
)

type EventStore struct {
	db *PostgresDB
}

func NewEventStore(db *PostgresDB) *EventStore {
	return &EventStore{
		db: db,
	}
}

// Save 이벤트들을 저장
func (r *EventStore) Save(ctx context.Context, accountId string, events []domain.Event) error {
	tx := r.db.db.WithContext(ctx)
	for _, event := range events {
		if err := tx.Create(&event).Error; err != nil {
			return fmt.Errorf("failed to save AccountCreatedEvent: %v", err)
		}
	}
	return nil
}

// Load 특정 계좌의 모든 이벤트 조회
func (r *EventStore) Load(ctx context.Context, accountId string) ([]domain.Event, error) {
	var events []domain.Event
	tx := r.db.db.WithContext(ctx).
		Where("account_id = ?", accountId).
		Order("created_at desc").
		Find(&events)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return events, nil
}
