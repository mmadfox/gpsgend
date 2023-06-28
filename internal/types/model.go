package types

const (
	MinModelValue = 3
	MaxModelValue = 60
)

type Model struct {
	val string
}

func (m Model) String() string {
	return m.val
}

func NewModel(val string) (Model, error) {
	model := Model{val: val}
	if err := model.validate(); err != nil {
		return Model{}, err
	}
	return model, nil
}

func (m Model) validate() error {
	if len(m.val) < MinModelValue {
		return ErrInvalidMinValue
	}
	if len(m.val) > MaxModelValue {
		return ErrInvalidMaxValue
	}
	return nil
}
