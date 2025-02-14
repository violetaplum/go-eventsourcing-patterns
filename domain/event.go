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
	GetAccountID() string
	GetEventType() string
	GetVersion() int
	GetCreatedAt() time.Time
	GetData() interface{}
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

// Event 관련 구조체들 추가
type AccountCreatedEvent struct {
	ID        string    `gorm:"column:id"`
	AccountId string    `gorm:"column:account_id"`
	EventType string    `gorm:"column:event_type"`
	Amount    int64     `gorm:"-"`
	EventData []byte    `gorm:"column:event_data"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (AccountCreatedEvent) TableName() string {
	return "events"
}

type MoneyDepositedEvent struct {
	AccountId string    `gorm:"column:account_id"`
	Amount    int64     `gorm:"column:amount"`
	EventType string    `gorm:"column:event_type"`
	EventData []byte    `gorm:"column:event_data"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (MoneyDepositedEvent) TableName() string {
	return "events"
}

type MoneyWithdrawnEvent struct {
	AccountId string    `gorm:"column:account_id"`
	Amount    int64     `gorm:"column:amount"`
	EventType string    `gorm:"column:event_type"`
	EventData []byte    `gorm:"column:event_data"`
	CreatedAt time.Time `gorm:"column:created"`
}

func (MoneyWithdrawnEvent) TableName() string {
	return "events"
}

// Event 인터페이스 구현
func (e AccountCreatedEvent) GetAccountID() string    { return e.AccountId }
func (e AccountCreatedEvent) GetEventType() string    { return string(AccountCreated) }
func (e AccountCreatedEvent) GetVersion() int         { return 1 }
func (e AccountCreatedEvent) GetCreatedAt() time.Time { return e.CreatedAt }
func (e AccountCreatedEvent) GetData() interface{}    { return e }

func (e MoneyDepositedEvent) GetAccountID() string    { return e.AccountId }
func (e MoneyDepositedEvent) GetEventType() string    { return string(MoneyDeposited) }
func (e MoneyDepositedEvent) GetVersion() int         { return 1 }
func (e MoneyDepositedEvent) GetCreatedAt() time.Time { return e.CreatedAt }
func (e MoneyDepositedEvent) GetData() interface{}    { return e }

func (e MoneyWithdrawnEvent) GetAccountID() string    { return e.AccountId }
func (e MoneyWithdrawnEvent) GetEventType() string    { return string(MoneyWithdrawn) }
func (e MoneyWithdrawnEvent) GetVersion() int         { return 1 }
func (e MoneyWithdrawnEvent) GetCreatedAt() time.Time { return e.CreatedAt }
func (e MoneyWithdrawnEvent) GetData() interface{}    { return e }
