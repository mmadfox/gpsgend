package types

import "errors"

var (
	ErrInvalidMinValue      = errors.New("invalid minimum value")
	ErrInvalidMaxValue      = errors.New("invalid maximum value")
	ErrInvalidRangeValue    = errors.New("invalid range values")
	ErrInvalidMinChargeTime = errors.New("invalid minimum charge time")
	ErrInvalidMaxChargeTime = errors.New("invlaid maximum charge time")
	ErrInvalidMinAmplitude  = errors.New("invalid minimum amplitude")
	ErrInvalidMaxAmplitude  = errors.New("invalid maximum amplitude")
	ErrInvalidName          = errors.New("invalid name value")
)
