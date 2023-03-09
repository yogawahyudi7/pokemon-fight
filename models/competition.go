package models

import (
	"time"

	"gorm.io/gorm"
)

type Competition struct {
	gorm.Model
	Rank1st   int `gorm:"column:rank1st"`
	Rank2nd   int `gorm:"column:rank2nd"`
	Rank3rd   int `gorm:"column:rank3rd"`
	Rank4th   int `gorm:"column:rank4th"`
	Rank5th   int `gorm:"column:rank5th"`
	SeasonId  int `gorm:"column:Season_id"`
	CreatedAt time.Time
	UpdatedAt time.Time

	DataScore  []Score `gorm:"foreignKey:CompetitionId;references:ID"`
	DataSeason Season  `gorm:"foreignKey:SeasonId"`
}

func (Competition) TableName() string {
	return "competition"
}
