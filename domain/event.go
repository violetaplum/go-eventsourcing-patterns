package domain

import (
	"context"
	"time"
)

type EventType string

const (
	AccountCreated EventType = "AccountCreated"
	MoneyDeposited EventType = "MoneyDeposited"
	MoneyWithdrawn EventType = "MoneyWithdrawn"
)

type Event struct {
	ID        string    `gorm:"column:id;primaryKey"`
	AccountID string    `gorm:"column:account_id"`
	EventType string    `gorm:"column:event_type"`
	EventData []byte    `gorm:"column:event_data"`
	CreatedAt time.Time `gorm:"column:created_at"`
	Amount    int64     `gorm:"-"`
}

func (e Event) TableName() string {
	return "events"
}

// Event 인터페이스 메서드 구현
func (e Event) GetAccountID() string {
	return e.AccountID
}

func (e Event) GetEventType() string {
	return e.EventType
}

func (e Event) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e Event) GetData() interface{} {
	return e.EventData
}

type EventPublisher interface {
	Publish(ctx context.Context, event Event) error
	PublishAll(ctx context.Context, events []Event) error
}

type EventConsumer interface {
	ProcessEvent(ctx context.Context, event Event) error
}

type EventHandler interface {
	Handle(ctx context.Context, event Event) error
}

type EventBus interface {
	Subscribe(eventType string, handler EventHandler)
	Publish(ctx context.Context, event Event) error
}

type EventStore interface {
	Save(ctx context.Context, accountId string, events []Event) error
	Load(ctx context.Context, accountId string) ([]Event, error)
}
