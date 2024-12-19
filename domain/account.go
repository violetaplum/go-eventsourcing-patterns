package domain

import (
	"context"
	"time"
)

type Account struct {
	ID      string
	Balance int64
	Events  []Event
}

type CreateAccountCommand struct {
	AccountID string
}

type DepositMoneyCommand struct {
	AccountID string
	Amount    int32
}

type WithdrawMoneyCommand struct {
	AccountID string
	Amount    int32
}

type GetAccountQuery struct {
	AccountID string
}

type GetBalanceQuery struct {
	AccountID string
}

type AccountView struct {
	ID       string
	Balance  int32
	UpdateAt time.Time
}

type CommandUseCase interface {
	CreateAccount(ctx context.Context, req CreateAccountCommand) error
	DepositMoney(ctx context.Context, req DepositMoneyCommand) error
	WithDrawMoney(ctx context.Context, req WithdrawMoneyCommand) error
}

type QueryUseCase interface {
	GetAccount(ctx context.Context, req GetAccountQuery) (*AccountView, error)
	GetBalance(ctx context.Context, req GetBalanceQuery) (*AccountView, error)
}

type EventStore interface {
	Save(ctx context.Context, accountId string, events []Event) error
	Load(ctx context.Context, accountId string) ([]Event, error)
}

type ReadStore interface {
	GetAccount(ctx context.Context, accountId string) (*AccountView, error)
}
