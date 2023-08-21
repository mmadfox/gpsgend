package generator

import (
	"context"

	"github.com/mmadfox/gpsgend/internal/types"
)

type Storage interface {
	// Find retrieves a Tracker based on the provided trackerID from the storage.
	// If the Tracker is found, it is returned along with a nil error.
	// If the Tracker is not found, nil is returned along with an error describing the situation.
	Find(ctx context.Context, trackerID types.ID) (*Tracker, error)

	// Insert adds a new Tracker to the storage.
	// Returns an error if the insertion process fails.
	Insert(ctx context.Context, t *Tracker) error

	// Delete removes a Tracker from the storage based on the provided trackerID.
	// Returns an error if the deletion process fails.
	Delete(ctx context.Context, trackerID types.ID) error

	// Update updates the information of an existing Tracker in the storage.
	// Returns an error if the update process fails.
	Update(ctx context.Context, t *Tracker) error
}
