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

type Event interface {
	GetAggregateID() string
	GetEventType() string
	GetVersion() int
	GetCreatedAt() time.Time
	GetData() interface{}
}

type EventPublisher interface {
	Publish(ctx context.Context, event Event) error
	PublishAll(ctx context.Context, events []Event) error
}

type EventHandler interface {
	Handle(ctx context.Context, event Event) error
}

type EventBus interface {
	Subscribe(eventType string, handler EventHandler)
	Publish(ctx context.Context, event Event) error
}

// Event 관련 구조체들 추가
type AccountCreatedEvent struct {
	AccountId string
	CreatedAt time.Time
}

type MoneyDepositedEvent struct {
	AccountId string
	Amount    int64
	CreatedAt time.Time
}

type MoneyWithdrawnEvent struct {
	AccountId string
	Amount    int64
	CreatedAt time.Time
}

// Event 인터페이스 구현
func (e AccountCreatedEvent) GetAggregateID() string  { return e.AccountId }
func (e AccountCreatedEvent) GetEventType() string    { return string(AccountCreated) }
func (e AccountCreatedEvent) GetVersion() int         { return 1 }
func (e AccountCreatedEvent) GetCreatedAt() time.Time { return e.CreatedAt }
func (e AccountCreatedEvent) GetData() interface{}    { return e }

func (e MoneyDepositedEvent) GetAggregateID() string  { return e.AccountId }
func (e MoneyDepositedEvent) GetEventType() string    { return string(MoneyDeposited) }
func (e MoneyDepositedEvent) GetVersion() int         { return 1 }
func (e MoneyDepositedEvent) GetCreatedAt() time.Time { return e.CreatedAt }
func (e MoneyDepositedEvent) GetData() interface{}    { return e }

func (e MoneyWithdrawnEvent) GetAggregateID() string  { return e.AccountId }
func (e MoneyWithdrawnEvent) GetEventType() string    { return string(MoneyWithdrawn) }
func (e MoneyWithdrawnEvent) GetVersion() int         { return 1 }
func (e MoneyWithdrawnEvent) GetCreatedAt() time.Time { return e.CreatedAt }
func (e MoneyWithdrawnEvent) GetData() interface{}    { return e }

type EventStore interface {
	Save(ctx context.Context, accountId string, events []Event) error
	Load(ctx context.Context, accountId string) ([]Event, error)
}
