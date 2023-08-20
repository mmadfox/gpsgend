package types

import (
	"github.com/google/uuid"
)

type ID struct {
	value uuid.UUID
}

func (t ID) String() string {
	return t.value.String()
}

func ParseID(id string) (ID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return ID{}, err
	}
	return ID{value: uid}, nil
}

func NewID() ID {
	return ID{value: uuid.New()}
}

func (t ID) Validate() error {
	if t.value == uuid.Nil {
		return ErrInvalidID
	}
	return nil
}
