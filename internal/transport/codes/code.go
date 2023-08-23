package codes

import (
	stderrors "errors"
)

func FromError(err error) int {
	for i := 0; i < len(table); i++ {
		code := table[i]
		if stderrors.Is(err, code.Err) {
			return code.Code
		}
	}
	return 0
}
