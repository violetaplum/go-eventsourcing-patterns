// domain/account.go

package domain

import (
	"context"
	"time"
)

// Account 모델
type Account struct {
	ID        string    `json:"id"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Command 요청 구조체들
type CreateAccountCommand struct {
	AccountID string
}

type DepositMoneyCommand struct {
	AccountID string
	Amount    int64
}

type WithdrawMoneyCommand struct {
	AccountID string
	Amount    int64
}

// Query 요청 구조체들
type GetAccountQuery struct {
	AccountID string
}

type GetBalanceQuery struct {
	AccountID string
}

// AccountView는 조회용 모델
type AccountView struct {
	ID      string
	Balance int64
}

//go:generate mockgen -source=account.go -destination=../mock/mock_account.go -package=mock

// AccountCommandService 커맨드(쓰기) 작업을 위한 인터페이스
type AccountCommandService interface {
	CreateAccount(ctx context.Context, cmd CreateAccountCommand) error
	DepositMoney(ctx context.Context, cmd DepositMoneyCommand) error
	WithdrawMoney(ctx context.Context, cmd WithdrawMoneyCommand) error
}

// AccountQueryService 쿼리(읽기) 작업을 위한 인터페이스
type AccountQueryService interface {
	GetAccount(ctx context.Context, query GetAccountQuery) (*AccountView, error)
	GetBalance(ctx context.Context, query GetBalanceQuery) (*AccountView, error)
}

// AccountRepository 저장소 작업을 위한 인터페이스
type AccountRepository interface {
	Save(ctx context.Context, account *Account) error
	FindByID(ctx context.Context, id string) (*Account, error)
	Update(ctx context.Context, account *Account) error
}

// EventStore 이벤트 저장소 작업을 위한 인터페이스
type EventStore interface {
	SaveEvents(ctx context.Context, accountID string, events []Event) error
	GetEvents(ctx context.Context, accountID string) ([]Event, error)
}
