package postgres

import (
	"github.com/golang/mock/gomock"
	"go-eventsourcing-patterns/domain"
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_AccountStore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := &gorm.DB{}

	accountStore := NewAccountStore(mockDB)
	account := domain.Account{
		ID:      "test-account-id",
		Balance: 100,
	}

	t.Run("Create Account", func(t *testing.T) {
		mockDB.EXPECT().Create
	})
}
