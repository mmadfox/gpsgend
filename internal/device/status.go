package device

import "fmt"

type Status uint16

const (
	Running Status = iota + 1
	Stopped
	Paused
	Stored
)

func (ds Status) String() string {
	switch ds {
	case Running:
		return "Running"
	case Stopped:
		return "Stopped"
	case Paused:
		return "Paused"
	case Stored:
		return "Stored"
	default:
		return "Unknown"
	}
}

func validateStatus(s Status) error {
	switch s {
	case Running, Stopped, Paused, Stored:
		return nil
	default:
		return fmt.Errorf("unknown device status")
	}
}
