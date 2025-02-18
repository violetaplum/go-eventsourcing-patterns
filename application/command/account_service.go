// application/command/account_service.go
package command

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"go-eventsourcing-patterns/domain"
	"time"
)

type AccountCommandService struct {
	accountStore   domain.AccountStore
	eventStore     domain.EventStore
	eventPublisher domain.EventPublisher
	txManager      domain.TransactionManager
}

func NewAccountCommandService(accountStore domain.AccountStore, eventStore domain.EventStore,
	eventPublisher domain.EventPublisher, txManager domain.TransactionManager) *AccountCommandService {
	return &AccountCommandService{
		accountStore:   accountStore,
		eventStore:     eventStore,
		eventPublisher: eventPublisher,
		txManager:      txManager,
	}
}

// CreateAccount는 Command를 받아서 처리
func (s *AccountCommandService) CreateAccount(ctx context.Context, cmd domain.CreateAccountCommand) error {
	ctx, err := s.txManager.Begin(ctx)
	if err != nil {
		return err
	}
	defer s.txManager.Rollback(ctx)

	account := &domain.Account{
		ID:        cmd.AccountId,
		Balance:   cmd.InitialBalance,
		UserName:  cmd.UserName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.accountStore.Create(ctx, account); err != nil {
		return err
	}

	eventData := map[string]interface{}{
		"account_id":      cmd.AccountId,
		"user_name":       cmd.UserName,
		"initial_balance": 1000,
	}
	byteData, err := json.Marshal(eventData)
	if err != nil {
		return err
	}

	event := domain.Event{
		AccountID: account.ID,
		CreatedAt: time.Now(),
		EventType: string(domain.AccountCreated),
		EventData: byteData,
	}

	if err := s.eventPublisher.Publish(ctx, event); err != nil {
		return err
	}

	return s.txManager.Commit(ctx)
}

// Deposit은 Command를 받아서 처리
func (s *AccountCommandService) Deposit(ctx context.Context, cmd domain.DepositCommand) error {
	ctx, err := s.txManager.Begin(ctx)
	if err != nil {
		return err
	}
	defer s.txManager.Rollback(ctx)

	account, err := s.accountStore.FindByID(ctx, cmd.AccountID)
	if err != nil {
		return err
	}

	originalBalance := account.Balance

	account.Balance += cmd.Amount
	account.UpdatedAt = time.Now()

	if err := s.accountStore.Update(ctx, account); err != nil {
		return err
	}
	eventId := uuid.New().String()
	eventData := map[string]interface{}{
		"id":               eventId,
		"account_id":       cmd.AccountID,
		"amount":           cmd.Amount,
		"original_balance": originalBalance,
	}
	byteData, err := json.Marshal(eventData)
	if err != nil {
		return err
	}

	event := domain.Event{
		ID:        eventId,
		AccountID: account.ID,
		Amount:    cmd.Amount,
		CreatedAt: time.Now(),
		EventData: byteData,
		EventType: string(domain.MoneyDeposited),
	}

	if err := s.eventPublisher.Publish(ctx, event); err != nil {
		return err
	}

	return s.txManager.Commit(ctx)
}

// Withdraw은 Command를 받아서 처리
func (s *AccountCommandService) Withdraw(ctx context.Context, cmd domain.WithdrawCommand) error {
	ctx, err := s.txManager.Begin(ctx)
	if err != nil {
		return err
	}
	defer s.txManager.Rollback(ctx)

	account, err := s.accountStore.FindByID(ctx, cmd.AccountID)
	if err != nil {
		return err
	}

	originalBalance := account.Balance

	if account.Balance < cmd.Amount {
		return domain.ErrInsufficientBalance
	}

	account.Balance -= cmd.Amount
	account.UpdatedAt = time.Now()

	if err := s.accountStore.Update(ctx, account); err != nil {
		return err
	}

	eventId := uuid.New().String()

	eventData := map[string]interface{}{
		"id":               eventId,
		"account_id":       cmd.AccountID,
		"amount":           cmd.Amount,
		"original_balance": originalBalance,
	}
	byteData, err := json.Marshal(eventData)
	if err != nil {
		return err
	}

	event := domain.Event{
		ID:        eventId,
		AccountID: account.ID,
		Amount:    cmd.Amount,
		CreatedAt: time.Now(),
		EventType: string(domain.MoneyWithdrawn),
		EventData: byteData,
	}

	if err := s.eventPublisher.Publish(ctx, event); err != nil {
		return err
	}

	return s.txManager.Commit(ctx)
}
