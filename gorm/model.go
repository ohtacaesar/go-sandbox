package gorm

import (
	"time"

	"gorm.io/gorm"
)

var models []any

func init() {
	models = append(models,
		&Resource{},
		&Event{},
		&DupTest{},
	)
}

type Resource struct {
	gorm.Model
	Value  int
	Events Events
}

type Event struct {
	gorm.Model
	ResourceID uint
	Value      string
	Date       time.Time `grom:"type:date"`
}

type Events []*Event

type DupTest struct {
	gorm.Model
	ID string `gorm:"primarykey"`
}
