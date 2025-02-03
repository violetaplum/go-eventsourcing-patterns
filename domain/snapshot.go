package domain

import "context"

// domain/snapshot.go
type Snapshot struct {
	AggregateID string
	Version     int
	State       []byte
}

type SnapshotStore interface {
	Save(ctx context.Context, snapshot *Snapshot) error
	Get(ctx context.Context, aggregateID string) (*Snapshot, error)
}
