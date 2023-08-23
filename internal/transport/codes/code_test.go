package codes

import (
	"errors"
	"testing"

	"github.com/mmadfox/gpsgend/internal/types"
)

func TestCode(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return valid code 1 ",
			args: args{
				err: types.ErrInvalidMinValue,
			},
			want: 1,
		},
		{
			name: "should return valid code 7",
			args: args{
				err: types.ErrInvalidMaxAmplitude,
			},
			want: 7,
		},
		{
			name: "should return invalid code when error not matches",
			args: args{
				err: errors.New("some errro"),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromError(tt.args.err); got != tt.want {
				t.Errorf("Code() = %v, want %v", got, tt.want)
			}
		})
	}
}
