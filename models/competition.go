package models

import "gorm.io/gorm"

type Competition struct {
	gorm.Model
	Rank1st int `gorm:"column:rank1st"`
	Rank2nd int `gorm:"column:rank2nd"`
	Rank3rd int `gorm:"column:rank3rd"`
	Rank4th int `gorm:"column:rank4th"`
	Rank5th int `gorm:"column:rank5th"`
	Season  int `gorm:"column:Season"`
}

func (Competition) TableName() string {
	return "competition"
}
