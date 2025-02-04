package domain

import (
	"context"
	"time"
)

// 읽기 모델을 위한 View 구조체
type AccountView struct {
	ID        string    `json:"id"`
	Balance   int64     `json:"balance"`
	UpdatedAt time.Time `json:"updated_at"`
}

//go:generate mockgen -source=projection.go -destination=../mock/mock_projection.go -package=mock

// ProjectionStore는 읽기 모델 저장소 인터페이스
type ProjectionStore interface {
	// 읽기 모델 저장
	SaveView(ctx context.Context, view *AccountView) error
	// ID로 읽기 모델 조회
	GetView(ctx context.Context, accountID string) (*AccountView, error)
	// 모든 읽기 모델 조회
	GetAllViews(ctx context.Context) ([]*AccountView, error)
}

// ProjectionHandler는 이벤트를 받아서 읽기 모델을 업데이트하는 인터페이스
type ProjectionHandler interface {
	// 이벤트 처리
	HandleEvent(ctx context.Context, event Event) error
	// 현재 프로젝션 상태 조회
	GetCurrentState(ctx context.Context, accountID string) (*AccountView, error)
}

// ProjectionManager는 프로젝션 처리를 관리하는 인터페이스
type ProjectionManager interface {
	// 프로젝션 재생성
	Rebuild(ctx context.Context) error
	// 프로젝션 상태 확인
	GetStatus(ctx context.Context) (string, error)
	// 프로젝션 처리 중지
	Stop(ctx context.Context) error
	// 프로젝션 처리 시작
	Start(ctx context.Context) error
}
