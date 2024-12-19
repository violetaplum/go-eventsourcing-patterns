package domain

import "time"

type EventType string

const (
	AccountCreated EventType = "AccountCreated"
	MoneyDeposited EventType = "MoneyDeposited"
	MoneyWithdrawn EventType = "MoneyWithdrawn"
)

type Event interface {
	GetType() EventType
	GetAggregateId() string
	GetData() interface{}
	GetTimestamp() time.Time
}

type AccountCreatedEvent struct {
	AccountId string
	Timestamp time.Time
}

type MoneyDepositedEvent struct {
	AccountId string
	Amount    int32
	Timestamp time.Time
}

type MoneyWithdrawnEvent struct {
	AccountId string
	Amount    int32
	Timestamp time.Time
}

func (e AccountCreatedEvent) GetType() EventType {
	return AccountCreated
}

func (e AccountCreatedEvent) GetAggregatedId() string {
	return e.AccountId
}

func (e AccountCreatedEvent) GetData() interface{} {
	return e
}

func (e AccountCreatedEvent) GetTimestamp() time.Time {
	return e.Timestamp
}
