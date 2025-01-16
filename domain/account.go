package domain

import (
	"context"
	"time"
)

type Account struct {
	ID        string    `gorm:"primaryKey;type:uuid"` // UUID 타입의 기본키 설정
	Balance   int64     `gorm:"not null"`             // 잔액을 나타내며 null을 허용하지 않음
	Events    []Event   // 관련된 도메인 이벤트
	CreatedAt time.Time `gorm:"autoCreateTime"` // 생성 시간 자동 설정
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // 업데이트 시간 자동 설정
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

//go:generate mockgen -source=account.go -destination=../mock/mock_account.go -package=mock

type CommandUseCase interface {
	CreateAccount(ctx context.Context, req CreateAccountCommand) error
	DepositMoney(ctx context.Context, req DepositMoneyCommand) error
	WithDrawMoney(ctx context.Context, req WithdrawMoneyCommand) error
}

type QueryUseCase interface {
	GetAccount(ctx context.Context, req GetAccountQuery) (*AccountView, error)
	GetBalance(ctx context.Context, req GetBalanceQuery) (*AccountView, error)
}

type AccountStore interface {
	Create(ctx context.Context, account Account) error                      // 계좌 생성
	GetByID(ctx context.Context, accountID string) (Account, error)         // 계좌 ID로 조회
	Update(ctx context.Context, account Account) error                      // 계좌 업데이트
	Delete(ctx context.Context, accountID string) error                     // 계좌 삭제
	GetAccount(ctx context.Context, accountId string) (*AccountView, error) // account_view 조회
}
