package models

import (
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	PokemonId     int
	CompetitionId int
	Rank1stCount  int `gorm:"column:rank_1st_count"`
	Rank2ndCount  int `gorm:"column:rank_2nd_count"`
	Rank3rdCount  int `gorm:"column:rank_3rd_count"`
	Rank4thCount  int `gorm:"column:rank_4th_count"`
	Rank5thCount  int `gorm:"column:rank_5th_count"`
	Points        int

	// DataPokemon     Pokemon     `gorm:"foreignKey:PokemonId"`
	DataCompetition Competition `gorm:"foreignKey:CompetitionId"`

	//gom migrator drop table
	TotalPoints int
	SeasonId    int
}

func (Score) TableName() string {
	return "score"
}
