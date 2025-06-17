package gorm

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	models = append(models, &DateTest{})
}

type DateTest struct {
	gorm.Model
	Time time.Time
	Date time.Time `gorm:"type:DATE"`
}
