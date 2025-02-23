package gorm

import (
	"time"

	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Value  string
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
