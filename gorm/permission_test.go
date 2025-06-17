package gorm

import (
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPermission(t *testing.T) {
	t.Parallel()
	opts := cmpopts.IgnoreFields(Permission{}, "ID")
	db := db.Debug()

	p := &Permission{V1: "a", V2: "a", V3: "a"}
	require.NoError(t, db.Save(&p).Error)
	log.Printf("created permission: %+v", p)
	{
		var got *Permission
		require.NoError(t, db.Where("id = ?", p.ID).First(&got).Error)
		assert.Empty(t, cmp.Diff(&Permission{V1: "a"}, got, opts))

	}
	t.Run("<-:updateは、更新時のみwriteされる", func(t *testing.T) {
		var tmp *Permission
		require.NoError(t, db.Where("id = ?", p.ID).First(&tmp).Error)
		tmp.V1 = "save"
		tmp.V2 = "save"
		tmp.V3 = "save"
		require.NoError(t, db.Save(tmp).Error)

		var got *Permission
		require.NoError(t, db.Where("id = ?", p.ID).First(&got).Error)
		assert.Empty(t, cmp.Diff(&Permission{V1: "a", V2: "save"}, got, opts))
	})
}
