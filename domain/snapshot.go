package domain

import (
	"context"
	"time"
)

// Snapshot 구조체 정의
type Snapshot struct {
	AggregateID string    `json:"aggregate_id"`
	Version     int       `json:"version"`
	Data        []byte    `json:"data"`
	CreatedAt   time.Time `json:"created_at"`
}

//go:generate mockgen -source=snapshot.go -destination=../mock/mock_snapshot.go -package=mock

// SnapshotStore 인터페이스
type SnapshotStore interface {
	// 스냅샷 저장
	Save(ctx context.Context, snapshot *Snapshot) error
	// 특정 애그리게이트 ID의 최신 스냅샷 조회
	Get(ctx context.Context, aggregateID string) (*Snapshot, error)
	// 특정 버전의 스냅샷 조회
	GetByVersion(ctx context.Context, aggregateID string, version int) (*Snapshot, error)
}

// SnapshotManager 인터페이스
type SnapshotManager interface {
	// 스냅샷 생성이 필요한지 확인
	ShouldTakeSnapshot(aggregateID string, eventCount int) bool
	// 새로운 스냅샷 생성
	TakeSnapshot(ctx context.Context, account *Account) error
	// 스냅샷으로부터 계정 상태 복원
	RestoreFromSnapshot(ctx context.Context, aggregateID string) (*Account, error)
}
