package gorm

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Value  string
	Events Events
}

type Event struct {
	gorm.Model
	ResourceID uint
	Value      string
}

type Events []*Event

type DupTest struct {
	gorm.Model
	ID string `gorm:"primarykey"`
}
