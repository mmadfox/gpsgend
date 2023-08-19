package types

type CustomID struct {
	value string
}

func (t CustomID) String() string {
	return t.value
}

func ParseCustomID(id string) (CustomID, error) {
	cid := CustomID{value: id}
	if err := cid.Validate(); err != nil {
		return CustomID{}, err
	}
	return cid, nil
}

func (t CustomID) Validate() error {
	if len(t.value) < 3 {
		return ErrInvalidMinValue
	}
	if len(t.value) > 256 {
		return ErrInvalidMaxValue
	}
	return nil
}
