package excelize

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuri/excelize/v2"
)

func TestOpenFile(t *testing.T) {
	t.Parallel()
	f, err := excelize.OpenFile("testdata/v16.92.xlsx")
	require.NoError(t, err)
	require.NoError(t, f.Close())
}

func TestGetDefaultFont(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		setup func(*testing.T) *excelize.File
		want  string
	}{
		{
			name:  "file created by excelize",
			setup: newFileAndSetCleanup,
			want:  "Calibri",
		},
		{
			name:  "file created by Excel",
			setup: openTestDataAndSetCleanup,
			want:  "游ゴシック",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.setup(t)
			font, err := f.GetDefaultFont()
			require.NoError(t, err)
			assert.Equal(t, tt.want, font)
		})
	}
}

func TestStylesFont(t *testing.T) {
	t.Parallel()
	type wantItem struct {
		Name string
		Sz   float64
	}
	type want []wantItem

	tests := []struct {
		name  string
		setup func(*testing.T) *excelize.File
		want  want
	}{
		{
			name:  "file created by excelize",
			setup: newFileAndSetCleanup,
			want: want{
				{
					Name: "Calibri",
					Sz:   11,
				},
			},
		},
		{
			name:  "file created by Excel",
			setup: openTestDataAndSetCleanup,
			want: want{
				{
					Name: "游ゴシック",
					Sz:   12,
				},
				{
					Name: "游ゴシック",
					Sz:   6,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.setup(t)
			for i, font := range f.Styles.Fonts.Font {
				assert.Equal(t, tt.want[i].Name, font.Name.Value())
				assert.Equal(t, tt.want[i].Sz, font.Sz.Value())

				// デバッグログ
				a, err := json.MarshalIndent(font, "", "    ")
				assert.NoError(t, err)
				if err != nil {
					continue
				}
				log.Printf("font[%d]: %s", i, a)
			}
		})
	}
}

// TestChangeDefaultFont
// File.StylesFonts.Font を操作することで、ファイル全体のデフォルトフォントを制御できる
func TestChangeDefaultFont(t *testing.T) {
	t.Parallel()
	f := excelize.NewFile()
	f.Styles.Fonts.Font[0].Name.Val = ptr("游ゴシック")
	f.Styles.Fonts.Font[0].Sz.Val = ptr[float64](12)
	require.NoError(t, f.SaveAs("result.xlsx"))
}
