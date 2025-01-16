package postgres

import (
	"context"
	"go-eventsourcing-patterns/domain"
	"gorm.io/gorm"
)

type AccountStore struct {
	db *gorm.DB
}

func NewAccountStore(db *gorm.DB) domain.AccountStore {
	return &AccountStore{db: db}
}

func (a AccountStore) Create(ctx context.Context, account domain.Account) error {
	return a.db.Create(&account).Error
}

func (a AccountStore) GetByID(ctx context.Context, accountID string) (domain.Account, error) {
	var account domain.Account
	if err := a.db.First(&account, "id = ?", accountID).Error; err != nil {
		return domain.Account{}, err
	}
	return account, nil
}

func (a AccountStore) Update(ctx context.Context, account domain.Account) error {
	return a.db.Save(&account).Error
}

func (a AccountStore) Delete(ctx context.Context, accountID string) error {
	return a.db.Delete(&domain.Account{}, accountID).Error
}

func (a AccountStore) GetAccount(ctx context.Context, accountId string) (*domain.AccountView, error) {
	var accountView domain.AccountView
	if err := a.db.First(&accountView, "id = ?", accountId).Error; err != nil {
		return nil, err
	}
	return &accountView, nil
}
