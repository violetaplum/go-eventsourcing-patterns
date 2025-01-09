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

// AccountCreatedEvent implements Event
func (e AccountCreatedEvent) GetType() EventType {
	return AccountCreated
}

func (e AccountCreatedEvent) GetAggregateId() string {
	return e.AccountId
}

func (e AccountCreatedEvent) GetData() interface{} {
	return e
}

func (e AccountCreatedEvent) GetTimestamp() time.Time {
	return e.Timestamp
}

// MoneyDepositedEvent implements Event
func (e MoneyDepositedEvent) GetType() EventType {
	return MoneyDeposited
}

func (e MoneyDepositedEvent) GetAggregateId() string {
	return e.AccountId
}

func (e MoneyDepositedEvent) GetData() interface{} {
	return e
}

func (e MoneyDepositedEvent) GetTimestamp() time.Time {
	return e.Timestamp
}

// MoneyWithdrawnEvent implements Event
func (e MoneyWithdrawnEvent) GetType() EventType {
	return MoneyWithdrawn
}

func (e MoneyWithdrawnEvent) GetAggregateId() string {
	return e.AccountId
}

func (e MoneyWithdrawnEvent) GetData() interface{} {
	return e
}

func (e MoneyWithdrawnEvent) GetTimestamp() time.Time {
	return e.Timestamp
}
