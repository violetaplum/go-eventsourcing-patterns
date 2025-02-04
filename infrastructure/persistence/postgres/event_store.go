package postgres

import (
	"context"
	"go-eventsourcing-patterns/domain"
	"gorm.io/gorm"
)

type EventStore struct {
	db *gorm.DB
}

func NewEventStore(db *gorm.DB) *EventStore {
	return &EventStore{
		db: db,
	}
}

// Save 이벤트들을 저장
func (r *EventStore) Save(ctx context.Context, accountId string, events []domain.Event) error {
	tx := r.db.WithContext(ctx)
	for _, event := range events {
		result := tx.Create(event)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// Load 특정 계좌의 모든 이벤트 조회
func (r *EventStore) Load(ctx context.Context, accountId string) ([]domain.Event, error) {
	var events []domain.Event
	tx := r.db.WithContext(ctx).
		Where("aggregate_id = ?", accountId).
		Order("version asc").
		Find(&events)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return events, nil
}
