package types

import (
	"fmt"

	"github.com/lucasb-eyer/go-colorful"
)

type Color struct {
	value colorful.Color
	ok    uint8
}

func (t Color) String() string {
	return t.value.Hex()
}

func (t Color) RGB() colorful.Color {
	return t.value
}

func ParseColor(color string) (Color, error) {
	rgb, err := colorful.Hex(color)
	if err != nil {
		return Color{}, err
	}
	return Color{value: rgb, ok: 1}, nil
}

func RandomColor() Color {
	rgb := colorful.FastHappyColor()
	return Color{value: rgb, ok: 1}
}

func (t Color) Validate() error {
	if t.ok != 1 {
		return fmt.Errorf("gpsgend/types: invalid color")
	}
	return nil
}
