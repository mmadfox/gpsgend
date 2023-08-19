package generator

import (
	"context"

	"github.com/mmadfox/gpsgend/internal/types"
)

type Storage interface {
	FindTracker(ctx context.Context, trackerID types.ID) (*Tracker, error)
	Insert(ctx context.Context, t *Tracker) error
	Delete(ctx context.Context, trackerID types.ID) error
	Update(ctx context.Context, t *Tracker) error
}
