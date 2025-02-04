package uuid

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		arg     string
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "uuid v7",
			arg:     "0194a9fe-f1ec-7646-b1f9-ed3e4dba9f52",
			wantErr: require.NoError,
		},
		{
			name:    "invalid uuid",
			arg:     "00000000-0000-0000-0000-000000000000",
			wantErr: require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, uuid.Validate(tt.arg))
		})
	}
}

func TestVersion(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			name: "uuid v7",
			arg:  "0194a9fe-f1ec-7646-b1f9-ed3e4dba9f52",
			want: 7,
		},
		{
			name: "uuid v4",
			arg:  "b3f2a7a2-0e7b-4a8e-8a9e-7b7b4b9b0f0b",
			want: 4,
		},
		{
			name: "invalid uuid",
			arg:  "00000000-0000-0000-0000-000000000000",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := uuid.Parse(tt.arg)
			require.NoError(t, err)
			require.EqualValues(t, tt.want, u.Version())
		})
	}
}
