package xerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

var childErr = errors.New("child error")

func TestErrorf(t *testing.T) {
	type args struct {
		format string
		any    []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "parent error: child error",
			args: args{
				format: "parent error: %w",
				any:    []any{childErr},
			},
			want: "parent error: child error",
		},
		{
			name: "parent error: child error",
			args: args{
				format: "%w: parent error",
				any:    []any{childErr},
			},
			want: "child error: parent error: child error", // ここがおかしい
		},
		{
			name: "empty",
			args: args{
				format: "%w",
				any:    []any{childErr},
			},
			want: "child error: child error", // %wだけの場合、2回出力される
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := xerrors.Errorf(tt.args.format, tt.args.any...)
			assert.Equal(t, tt.want, err.Error())
			t.Logf("%+v", err)
		})
	}
}
