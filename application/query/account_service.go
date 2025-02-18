package query

import (
	"context"
	"go-eventsourcing-patterns/domain"
)

type AccountQueryService struct {
	eventStore   domain.EventStore
	accountStore domain.AccountStore
}

func NewAccountQueryService(
	accountStore domain.AccountStore,
	eventStore domain.EventStore,
) *AccountQueryService {
	return &AccountQueryService{
		eventStore:   eventStore,
		accountStore: accountStore,
	}
}

func (s *AccountQueryService) GetAccountByID(ctx context.Context, accountID string) (*domain.AccountResponse, error) {
	// 계정 정보 조회
	account, err := s.accountStore.FindByID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	events, err := s.eventStore.Load(ctx, accountID)
	if err != nil {
		return nil, err
	}

	totalDeposits := int64(0)
	totalWithdrawals := int64(0)
	transactionCount := 0

	for _, event := range events {
		switch event.GetEventType() {
		case string(domain.MoneyDeposited):
			totalDeposits += event.Amount
			transactionCount++
		case string(domain.MoneyWithdrawn):
			totalWithdrawals += event.Amount
			transactionCount++
		}
	}

	//이벤트 히스토리를 활용하여 계정의 추가 정보 제공
	//총 입금액, 총 출금액, 트랜잭션 횟수 등의 추가 정보 계산
	//이벤트 소싱의 장점 활용 (상태뿐만 아니라 상태 변경 이력 활용)

	return &domain.AccountResponse{
		ID:               account.ID,
		Balance:          account.Balance,
		CreatedAt:        account.CreatedAt,
		UpdatedAt:        account.UpdatedAt,
		UserName:         account.UserName,
		TotalDeposits:    totalDeposits,
		TotalWithdrawals: totalWithdrawals,
		TransactionCount: transactionCount,
	}, nil
}

// ListAccounts 모든 계정 조회
func (s *AccountQueryService) ListAccounts(ctx context.Context) ([]domain.AccountResponse, error) {
	accounts, err := s.accountStore.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []domain.AccountResponse
	for _, account := range accounts {
		// 각 계정의 이벤트 히스토리 로드
		events, err := s.eventStore.Load(ctx, account.ID)
		if err != nil {
			// 이벤트 로드 실패해도 계정 정보는 반환
			responses = append(responses, domain.AccountResponse{
				ID:        account.ID,
				Balance:   account.Balance,
				CreatedAt: account.CreatedAt,
				UpdatedAt: account.UpdatedAt,
				UserName:  account.UserName,
			})
			continue
		}

		// 이벤트 히스토리를 사용하여 추가 정보 계산
		totalDeposits := int64(0)
		totalWithdrawals := int64(0)
		transactionCount := 0

		for _, event := range events {
			switch event.GetEventType() {
			case string(domain.MoneyDeposited):
				totalDeposits += event.Amount
				transactionCount++
			case string(domain.MoneyWithdrawn):
				totalWithdrawals += event.Amount
				transactionCount++
			}
		}

		responses = append(responses, domain.AccountResponse{
			ID:               account.ID,
			Balance:          account.Balance,
			CreatedAt:        account.CreatedAt,
			UpdatedAt:        account.UpdatedAt,
			UserName:         account.UserName,
			TotalDeposits:    totalDeposits,
			TotalWithdrawals: totalWithdrawals,
			TransactionCount: transactionCount,
		})

	}

	return responses, nil
}

// GetAccountHistory 계정의 이벤트 히스토리 조회
func (s *AccountQueryService) GetAccountHistory(ctx context.Context, accountID string) ([]domain.Event, error) {
	// 계정 존재 여부 먼저 확인
	_, err := s.accountStore.FindByID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	// 이벤트 히스토리 로드
	return s.eventStore.Load(ctx, accountID)
}
