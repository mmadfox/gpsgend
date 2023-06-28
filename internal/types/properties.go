package types

const MaxPropertiesValue = 10

type Properties struct {
	val map[string]string
}

func NewProperties(props map[string]string) (Properties, error) {
	p := Properties{val: props}
	if err := p.validate(); err != nil {
		return Properties{}, err
	}
	return p, nil
}

func (p Properties) ToMap() map[string]string {
	props := make(map[string]string, len(p.val))
	for k, v := range p.val {
		props[k] = v
	}
	return props
}

func (p Properties) validate() error {
	if len(p.val) > MaxPropertiesValue {
		return ErrInvalidMaxValue
	}
	return nil
}
