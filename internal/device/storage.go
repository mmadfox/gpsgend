package device

import (
	"context"

	"github.com/google/uuid"
)

type Storage interface {
	FindByID(ctx context.Context, deviceID uuid.UUID) (*Device, error)
	Insert(ctx context.Context, d *Device) error
	Update(ctx context.Context, d *Device) error
	Delete(ctx context.Context, deviceID uuid.UUID) error
	WalkByStatus(ctx context.Context, s []Status, fn func(*Device) error) error
}

func Version(d *Device) int {
	return d.version
}

func SetVersion(d *Device, version int) {
	d.version = version
}
