package gorm

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestPreload
// 構造体同士の参照で循環参照が発生する場合、すでにどこかで利用した構造体は除外されるっぽい
func TestPreload(t *testing.T) {

	// テストデータの作成
	//users := []User{
	//	{
	//		Name:  "User1",
	//		Email: "user1@example.com",
	//		Posts: []Post{
	//			{
	//				Title:   "User1's first post",
	//				Content: "Content of first post",
	//				Comments: []Comment{
	//					{Content: "First comment on first post"},
	//					{Content: "Second comment on first post"},
	//				},
	//			},
	//			{
	//				Title:   "User1's second post",
	//				Content: "Content of second post",
	//				Comments: []Comment{
	//					{Content: "First comment on second post"},
	//				},
	//			},
	//		},
	//	},
	//	{
	//		Name:  "User2",
	//		Email: "user2@example.com",
	//		Posts: []Post{
	//			{
	//				Title:   "User2's first post",
	//				Content: "Content of first post by user2",
	//				Comments: []Comment{
	//					{Content: "Comment on User2's post"},
	//				},
	//			},
	//		},
	//	},
	//}

	// データの保存
	//for _, user := range users {
	//	require.NoError(t, db.Create(&user).Error)
	//}

	// ユーザーとその関連するPostとCommentを一度に取得
	t.Run("ユーザーとその関連するPostとCommentを一度に取得", func(t *testing.T) {
		db := db.Debug() // デバッグモードで実行
		var user User
		err := db.
			Preload("Posts").
			Preload("Posts.Comments").
			First(&user).
			Error
		require.NoError(t, err)
		data, err := json.MarshalIndent(user, "", "  ")
		require.NoError(t, err)
		t.Log(string(data))
	})

	t.Run("ネストされたPreloadを一度に指定", func(t *testing.T) {
		db := db.Debug()
		var user User
		err := db.
			Preload("Posts.Comments").
			First(&user).
			Error
		require.NoError(t, err)
		data, err := json.MarshalIndent(user, "", "  ")
		require.NoError(t, err)
		t.Log(string(data))
	})

	t.Run("複数ユーザー取得時にPreloadを設定する", func(t *testing.T) {
		db := db.Debug()
		var users []User
		err := db.
			Preload("Posts.Comments").
			Find(&users).
			Error
		require.NoError(t, err)
		data, err := json.MarshalIndent(users, "", "  ")
		require.NoError(t, err)
		t.Log(string(data))
	})

	t.Run("Post起点", func(t *testing.T) {
		db := db.Debug()
		var posts []Post
		err := db.
			Preload("User").
			Preload("Comments").
			Find(&posts).
			Error
		require.NoError(t, err)
		data, err := json.MarshalIndent(posts, "", "  ")
		require.NoError(t, err)
		t.Log(string(data))

		for _, post := range posts {
			for _, comment := range post.Comments {
				assert.Empty(t, comment.Post)
			}
		}
	})
}
