package device

import (
	"github.com/google/uuid"
	gpsgen "github.com/mmadfox/go-gpsgen"
)

type Generator interface {
	Attach(*gpsgen.Device)
	Lookup(deviceID uuid.UUID) (*gpsgen.Device, error)
	Detach(id uuid.UUID)
}
