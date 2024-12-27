package excelize

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xuri/excelize/v2"
)

func ptr[T any](v T) *T { return &v }
func openTestDataAndSetCleanup(t *testing.T) *excelize.File {
	t.Helper()
	f, err := excelize.OpenFile("testdata/v16.92.xlsx")
	require.NoError(t, err)
	t.Cleanup(func() { _ = f.Close() })

	return f
}

func newFileAndSetCleanup(t *testing.T) *excelize.File {
	t.Helper()
	f := excelize.NewFile()
	t.Cleanup(func() { _ = f.Close() })

	return f
}
