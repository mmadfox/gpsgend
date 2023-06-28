package types

const (
	MinDescriptionValue = 0
	MaxDescriptionValue = 256
)

type Description struct {
	val string
}

func (d Description) String() string {
	return d.val
}

func NewDescription(val string) (Description, error) {
	descr := Description{val: val}
	if err := descr.validate(); err != nil {
		return Description{}, err
	}
	return descr, nil
}

func (d Description) validate() error {
	if len(d.val) < MinDescriptionValue {
		return ErrInvalidMinValue
	}
	if len(d.val) > MaxDescriptionValue {
		return ErrInvalidMaxValue
	}
	return nil
}
