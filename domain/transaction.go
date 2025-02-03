package domain

import "context"

// domain/transaction.go
type TransactionManager interface {
	// 트랜잭션 시작
	Begin(ctx context.Context) (context.Context, error)
	// 트랜잭션 커밋
	Commit(ctx context.Context) error
	// 트랜잭션 롤백
	Rollback(ctx context.Context) error
}
