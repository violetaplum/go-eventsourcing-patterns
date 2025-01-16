package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-eventsourcing-patterns/domain"
	"time"
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

func (s *EventStore) Load(ctx context.Context, accountId string) ([]domain.Event, error) {
	query := `SELECT event_type, event_data, created_at 
				FROM events 
				WHERE aggregated_id = $1 
				ORDER BY created_at ASC `

	rows, err := s.db.QueryContext(ctx, query, accountId)
	if err != nil {
		return nil, fmt.Errorf("error quering events: %w", err)
	}

	defer rows.Close()
	var events []domain.Event
	for rows.Next() {
		var eventType string
		var eventData []byte
		var createdAt time.Time

		if err := rows.Scan(&eventType, &eventData, &createdAt); err != nil {
			return nil, fmt.Errorf("error scanning event row: %w", err)
		}

		var event domain.Event
		switch domain.EventType(eventType) {
		case domain.AccountCreated:
			var accountCreated domain.AccountCreatedEvent
			if err := json.Unmarshal(eventData, &accountCreated); err != nil {
				return nil, fmt.Errorf("error unmarshaling AccountCreatedEvent: %w", err)
			}
			accountCreated.Timestamp = createdAt
			event = accountCreated
		case domain.MoneyDeposited:
			var moneyDeposited domain.MoneyDepositedEvent
			if err := json.Unmarshal(eventData, &moneyDeposited); err != nil {
				return nil, fmt.Errorf("error unmarshaling MoneyDepositedEvent: %w", err)
			}
			moneyDeposited.Timestamp = createdAt
			event = moneyDeposited
		case domain.MoneyWithdrawn:
			var moneyWithdrawn domain.MoneyWithdrawnEvent
			if err := json.Unmarshal(eventData, &moneyWithdrawn); err != nil {
				return nil, fmt.Errorf("error unmarshaling MoneyWithdrawnEvent: %w", err)
			}
			moneyWithdrawn.Timestamp = createdAt
			event = moneyWithdrawn
		default:
			return nil, fmt.Errorf("unknown event type: %s", eventType)
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error interating rows: %w", err)
	}

	return events, nil

}
