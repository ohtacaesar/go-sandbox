package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
