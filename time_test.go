package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go-sandbox/common"
)

func TestParseTime(t *testing.T) {
	t.Parallel()

	type args struct {
		layout string
		value  string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr require.ErrorAssertionFunc
	}{
		{
			name: "hour out of range",
			args: args{
				layout: "15:04:05",
				value:  "31:01:02",
			},
			wantErr: func(t require.TestingT, err error, i ...interface{}) {
				var parseErr *time.ParseError
				require.ErrorAs(t, err, &parseErr)
				assert.Equal(t, "parsing time \"31:01:02\": hour out of range", err.Error())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := time.Parse(tt.args.layout, tt.args.value)
			tt.wantErr(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestTime_IsZero(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		t    time.Time
		want bool
	}{
		{
			name: "zero",
			t:    time.Time{},
			want: true,
		},
		{
			name: "",
			t:    (time.Time{}).In(common.JST),
			want: true,
		},
		{
			name: "",
			t:    time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "not zero",
			t:    time.Now(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.t.IsZero())
		})
	}
}
