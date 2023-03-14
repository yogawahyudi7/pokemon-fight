package models

import (
	"time"

	"gorm.io/gorm"
)

type Season struct {
	gorm.Model
	Name      string `gorm:"unique"`
	StartDate time.Time
	EndDate   time.Time
}

func (Season) TableName() string {
	return "season"
}
