package gorm

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTypeDate(t *testing.T) {
	t.Parallel()
	t.Run("DBのタイムゾーン設定に関わらずUTCが返ってくる", func(t *testing.T) {
		var timezone string
		require.NoError(t, db.Raw("SHOW timezone").Scan(&timezone).Error)
		assert.Equal(t, "Asia/Tokyo", timezone)
		now := time.Now().In(JST)
		o := &DateTest{
			Time: now,
			Date: time.Date(2006, 1, 2, 0, 0, 0, 0, JST),
		}
		require.NoError(t, db.Create(o).Error)

		var got *DateTest
		require.NoError(t, db.Find(&got, o.ID).Error)

		t.Log(got.Time.Zone())
		assert.Equal(t, time.UTC, got.Date.Location())
		assert.Empty(t, cmp.Diff(
			&DateTest{
				Time: now,
				Date: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			got,
			cmpopts.IgnoreFields(DateTest{}, "ID", "CreatedAt", "UpdatedAt"),
		))
	})
}
