package infraKafka

import (
	"context"
	"go-eventsourcing-patterns/domain"
)

type AccountCreatedHandler struct {
	eventStore domain.EventStore
}

func NewAccountCreatedHandler(eventStore domain.EventStore) *AccountCreatedHandler {
	return &AccountCreatedHandler{
		eventStore: eventStore,
	}
}

func (h *AccountCreatedHandler) Handle(ctx context.Context, event domain.Event) error {
	if err := h.eventStore.Save(ctx, event.GetAccountID(), []domain.Event{event}); err != nil {
		return err
	}

	return nil
}

type MoneyDepositHandler struct {
	eventStore domain.EventStore
}

func NewMoneyDepositHandler(eventStore domain.EventStore) *MoneyDepositHandler {
	return &MoneyDepositHandler{eventStore: eventStore}
}

func (h *MoneyDepositHandler) Handle(ctx context.Context, event domain.Event) error {
	if err := h.eventStore.Save(ctx, event.GetAccountID(), []domain.Event{event}); err != nil {
		return err
	}

	return nil
}

type MoneyWithdrawHandler struct {
	eventStore domain.EventStore
}

func NewMoneyWithdrawHandler(eventStore domain.EventStore) *MoneyWithdrawHandler {
	return &MoneyWithdrawHandler{eventStore: eventStore}
}

func (h *MoneyWithdrawHandler) Handle(ctx context.Context, event domain.Event) error {
	if err := h.eventStore.Save(ctx, event.GetAccountID(), []domain.Event{event}); err != nil {
		return err
	}
	return nil
}
