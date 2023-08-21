package generator

import "context"

type Bootstraper interface {
	LoadTrackers(ctx context.Context, p Processes) error
	UnloadTrackers(ctx context.Context, p Processes) error
}
