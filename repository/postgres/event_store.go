package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-eventsourcing-patterns/domain"
)

type EventStore struct {
	db *sql.DB
}

func NewEventStore(db *sql.DB) *EventStore {
	return &EventStore{db: db}
}

func (s *EventStore) Save(ctx context.Context, accountId string, events []domain.Event) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error beginning transaction: %w", err)
	}

	defer tx.Rollback()

	query := `INSERT INTO events (aggregate_id, event_type, event_data, created_at) 
				VALUES ($1, $2, $3, $4)`

	for _, event := range events {
		eventData, err := json.Marshal(event.GetData())
		if err != nil {
			return fmt.Errorf("error marshaling event: %w", err)
		}

		_, err = tx.ExecContext(ctx, query, accountId, string(event.GetType()),
			eventData, event.GetTimestamp())

		if err != nil {
			return fmt.Errorf("error inserting event: %w", err)
		}
	}

	return tx.Commit()
}
