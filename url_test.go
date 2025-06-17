package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryEscape(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "success",
			args: "abc",
			want: "abc",
		},
		{
			name: "space",
			args: "abc def",
			want: "abc+def",
		},
		{
			name: "quote",
			args: "\"",
			want: "%22",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := url.QueryEscape(tt.args)
			require.Equal(t, tt.want, got)
		})
	}
}
