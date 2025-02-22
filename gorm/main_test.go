package gorm

import (
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	gormConfig := &gorm.Config{
		QueryFields: true,
	}
	var err error
	db, err = gorm.Open(postgres.New(postgres.Config{DSN: "postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable"}), gormConfig)
	if err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(&Resource{}, &Event{}, &DupTest{}); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
