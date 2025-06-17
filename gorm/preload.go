package gorm

import (
	"gorm.io/gorm"
)

func init() {
	models = append(models, &User{}, &Post{}, &Comment{})
}

// User はユーザーを表すモデルです
// One to Manyの関係で複数のPostを持ちます
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"uniqueIndex"`
	Posts []Post `gorm:"foreignKey:UserID"`
}

// Post はユーザーの投稿を表すモデルです
// UserとOne to Manyの関係になっており、また複数のCommentを持つ
type Post struct {
	gorm.Model
	Title    string
	Content  string
	UserID   uint
	User     User      `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}

// Comment は投稿に対するコメントを表すモデルです
// PostとOne to Manyの関係になっています
type Comment struct {
	gorm.Model
	Content string
	PostID  uint
	Post    Post `gorm:"foreignKey:PostID"`
}

// InitDB はテストデータを作成するヘルパー関数です
func InitDB(db *gorm.DB) error {
	// テーブルの初期化
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		return err
	}

	return nil
}
