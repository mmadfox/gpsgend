package codes

import (
	"errors"
	"net/http"
	"testing"

	"github.com/mmadfox/gpsgend/internal/generator"
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
			name: "should return http.StatusBadRequest when error is types.ErrInvalidMinValue",
			args: args{
				err: types.ErrInvalidMinValue,
			},
			want: http.StatusBadRequest,
		},
		{
			name: "should return http.StatusBadRequest when error is types.ErrInvalidMaxAmplitude",
			args: args{
				err: types.ErrInvalidMaxAmplitude,
			},
			want: http.StatusBadRequest,
		},
		{
			name: "should return http.StatusUnprocessableEntity when error is generator.ErrTrackerIsAlreadyRunning",
			args: args{
				err: generator.ErrTrackerIsAlreadyRunning,
			},
			want: http.StatusUnprocessableEntity,
		},
		{
			name: "should return http.StatusNotFound when error is generator.ErrTrackNotFound",
			args: args{
				err: generator.ErrTrackNotFound,
			},
			want: http.StatusNotFound,
		},
		{
			name: "should return  http.StatusInternalServerError when error not matches",
			args: args{
				err: errors.New("some errro"),
			},
			want: http.StatusInternalServerError,
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
