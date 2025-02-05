package store

import (
	"context"
	"fmt"
	"go-eventsourcing-patterns/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresDB는 데이터베이스 연결과 트랜잭션을 관리하는 구조체
type PostgresDB struct {
	db *gorm.DB
}

// NewPostgresDB는 새로운 PostgresDB 인스턴스를 생성합니다
func NewPostgresDB(config *domain.Config) (*PostgresDB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
		config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &PostgresDB{db: db}, nil
}

// GetDB는 gorm.DB 인스턴스를 반환합니다
func (p *PostgresDB) GetDB() *gorm.DB {
	return p.db
}

// domain.TransactionManager 인터페이스 구현
func (p *PostgresDB) Begin(ctx context.Context) (context.Context, error) {
	tx := p.db.Begin()
	if tx.Error != nil {
		return ctx, fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}
	return context.WithValue(ctx, domain.TxKey, tx), nil
}

func (p *PostgresDB) Commit(ctx context.Context) error {
	tx, ok := ctx.Value(domain.TxKey).(*gorm.DB)
	if !ok {
		return fmt.Errorf("no transaction found in context")
	}
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (p *PostgresDB) Rollback(ctx context.Context) error {
	tx, ok := ctx.Value(domain.TxKey).(*gorm.DB)
	if !ok {
		return fmt.Errorf("no transaction found in context")
	}
	if err := tx.Rollback().Error; err != nil {
		return fmt.Errorf("failed to rollback transaction: %w", err)
	}
	return nil
}

// domain.UnitOfWork 인터페이스 구현
func (p *PostgresDB) RunInTransaction(ctx context.Context, fn func(ctx context.Context) error) error {

	newCtx, err := p.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			p.Rollback(newCtx)
			panic(r) // re-throw panic after rollback
		}
	}()

	if err := fn(newCtx); err != nil {
		if rbErr := p.Rollback(newCtx); rbErr != nil {
			return fmt.Errorf("rollback failed: %v (original error: %v)", rbErr, err)
		}
		return err
	}

	if err := p.Commit(newCtx); err != nil {
		return err
	}

	return nil
}

func (p *PostgresDB) GetTransactionContext(ctx context.Context) context.Context {
	tx, ok := ctx.Value(domain.TxKey).(*gorm.DB)
	if !ok {
		return ctx
	}
	return context.WithValue(ctx, domain.TxKey, tx)
}

// Close closes the database connection
func (p *PostgresDB) Close() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}
	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}
	return nil
}
