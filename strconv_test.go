package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAtoi(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		args    string
		want    int
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "success",
			args:    "00012",
			want:    12,
			wantErr: require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := strconv.Atoi(tt.args)
			tt.wantErr(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestQuote(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "success",
			args: "abc",
			want: "\"abc\"",
		},
		{
			name: "quote",
			args: "\"",
			want: "\"\\\"\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strconv.Quote(tt.args)
			require.Equal(t, tt.want, got)
		})
	}
}
