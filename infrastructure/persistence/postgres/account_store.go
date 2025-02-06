package postgres

import (
	"context"
	"go-eventsourcing-patterns/domain"
)

type AccountStore struct {
	db *PostgresDB
}

func NewAccountStore(db *PostgresDB) *AccountStore {
	return &AccountStore{
		db: db,
	}
}

// Save 새 계좌 생성
func (r *AccountStore) Create(ctx context.Context, account *domain.Account) error {
	tx := r.db.db.WithContext(ctx).Create(account)
	return tx.Error
}

// FindByID ID로 계좌 조회
func (r *AccountStore) FindByID(ctx context.Context, id string) (*domain.Account, error) {
	var account domain.Account
	tx := r.db.db.WithContext(ctx).First(&account, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &account, nil
}

// Update 계좌 정보 업데이트
func (r *AccountStore) Update(ctx context.Context, account *domain.Account) error {
	tx := r.db.db.WithContext(ctx).Save(account)
	return tx.Error
}

// Delete 계좌 삭제
func (r *AccountStore) Delete(ctx context.Context, id string) error {
	tx := r.db.db.WithContext(ctx).Delete(&domain.Account{}, "id = ?", id)
	return tx.Error
}

// ListAll 모든 계좌 조회
func (s *AccountStore) ListAll(ctx context.Context) ([]*domain.Account, error) {
	var accounts []*domain.Account
	tx := s.db.db.WithContext(ctx).Find(&accounts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return accounts, nil
}
