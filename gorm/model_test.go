package gorm

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	db := db.Debug()

	now := time.Now()
	r := &Resource{
		Value: "value",
		Events: Events{
			{
				Value: "value",
				Date:  now,
			},
		},
	}
	require.NoError(t, db.Create(r).Error)
	r.Value = "value2"
	require.NoError(t, db.Save(r).Error)
	r.Events[0].Value = "value2" // これは反映されない
	require.NoError(t, db.Save(r).Error)

	var r2 *Resource
	require.NoError(t, db.First(&r2, r.ID).Error)
	require.Empty(t, r2.Events)
	r2.Events = append(r2.Events, &Event{Value: "appended"})
	require.NoError(t, db.Save(r2).Error)

	var r3 *Resource
	require.NoError(t, db.Preload("Events").First(&r3, r.ID).Error)
	require.Len(t, r3.Events, 2)
	// Verify that the date in the first event is stored correctly
	require.WithinDuration(t, now, r3.Events[0].Date, time.Second)
}

func TestCreate2(t *testing.T) {
	id := uuid.New().String()

	eg, _ := errgroup.WithContext(context.Background())
	for i := range 2 {
		eg.Go(func() error {
			return db.Transaction(func(tx *gorm.DB) (err error) {
				o := &DupTest{ID: id}
				if err = tx.Create(o).Error; err != nil {
					t.Logf("%d: error: %+v", i, err)
					return xerrors.Errorf("create: %w", err)
				}
				t.Logf("%d: waiting...", i)
				time.Sleep(time.Second * 5)
				return nil
			})
		})
	}
	if err := eg.Wait(); err != nil {
		t.Logf("%+v", err)
	}
}
