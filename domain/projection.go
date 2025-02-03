package domain

import "context"

// domain/projection.go
type ProjectionStore interface {
	// 읽기 모델 저장
	SaveView(ctx context.Context, accountView *AccountView) error
	// 읽기 모델 조회
	GetView(ctx context.Context, accountID string) (*AccountView, error)
}

type ProjectionHandler interface {
	// 이벤트를 받아서 읽기 모델 업데이트
	HandleEvent(ctx context.Context, event Event) error
	// 현재 프로젝션 상태 조회
	GetCurrentState(ctx context.Context, accountID string) (*AccountView, error)
}
