package types

import "github.com/mmadfox/go-gpsgen/types"

type Model struct {
	val string
}

func (m Model) String() string {
	return m.val
}

func ParseModel(val string) (Model, error) {
	model := Model{val: val}
	if err := model.Validate(); err != nil {
		return Model{}, err
	}
	return model, nil
}

func (m Model) Validate() error {
	if len(m.val) < types.MinModelLen {
		return ErrInvalidMinValue
	}
	if len(m.val) > types.MaxModelLen {
		return ErrInvalidMaxValue
	}
	return nil
}
