package domain

import (
	"context"
	"time"
)

const (
	AccountCreated EventType = "AccountCreated"
	MoneyDeposited EventType = "MoneyDeposited"
	MoneyWithdrawn EventType = "MoneyWithdrawn"
)

type Event interface {
	AggregateID() string
	EventType() string
	Version() int
	Timestamp() time.Time
	Data() interface{}
}

// domain/event.go
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
