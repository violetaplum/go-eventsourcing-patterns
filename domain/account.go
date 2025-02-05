package domain

import (
	"context"
	"time"
)

// Account 도메인 모델
type Account struct {
	ID        string    `json:"id"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Command 구조체들
type CreateAccountCommand struct {
	InitialBalance int64
}

type DepositCommand struct {
	AccountID string
	Amount    int64
}

type WithdrawCommand struct {
	AccountID string
	Amount    int64
}

// HTTP 요청/응답을 위한 DTO 구조체들
type CreateAccountRequest struct {
	InitialBalance int64 `json:"initial_balance"`
}

type CreateAccountResponse struct {
	ID        string    `json:"id"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type DepositRequest struct {
	Amount int64 `json:"amount"`
}

type WithdrawRequest struct {
	Amount int64 `json:"amount"`
}

type AccountResponse struct {
	ID        string    `json:"id"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	TotalDeposits    int64 `json:"total_deposits"`
	TotalWithdrawals int64 `json:"total_withdrawals"`
	TransactionCount int   `json:"transaction_count"`
}

//go:generate mockgen -source=account.go -destination=../mock/mock_account.go -package=mock

// Account 서비스 인터페이스
type AccountCommandService interface {
	CreateAccount(ctx context.Context, cmd CreateAccountCommand) error
	Deposit(ctx context.Context, cmd DepositCommand) error
	Withdraw(ctx context.Context, cmd WithdrawCommand) error
}

type AccountQueryService interface {
	GetAccountByID(ctx context.Context, accountID string) (*AccountResponse, error)
	ListAccounts(ctx context.Context) ([]AccountResponse, error)
	GetAccountHistory(ctx context.Context, accountID string) ([]Event, error)
}

// Account 저장소 인터페이스
type AccountStore interface {
	Create(ctx context.Context, account *Account) error
	FindByID(ctx context.Context, id string) (*Account, error)
	Update(ctx context.Context, account *Account) error
	ListAll(ctx context.Context) ([]*Account, error)
}
