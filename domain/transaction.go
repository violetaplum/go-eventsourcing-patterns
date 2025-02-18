package domain

import (
	"context"
)

//go:generate mockgen -source=transaction.go -destination=domain/mock/mock_transaction.go -package=mock

// TransactionManager 인터페이스
type TransactionManager interface {
	// 트랜잭션 시작
	Begin(ctx context.Context) (context.Context, error)
	// 트랜잭션 커밋
	Commit(ctx context.Context) error
	// 트랜잭션 롤백
	Rollback(ctx context.Context) error
}

// UnitOfWork 인터페이스
type UnitOfWork interface {
	// 트랜잭션 내에서 실행될 함수를 받아 처리
	RunInTransaction(ctx context.Context, fn func(ctx context.Context) error) error
	// 트랜잭션 컨텍스트 가져오기
	GetTransactionContext(ctx context.Context) context.Context
}

// TransactionContext는 트랜잭션 정보를 컨텍스트에 저장하기 위한 키 타입
type TransactionContext struct{}

// TxKey는 컨텍스트에서 트랜잭션을 식별하기 위한 키
var TxKey = TransactionContext{}
