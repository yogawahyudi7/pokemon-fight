package models

import "gorm.io/gorm"

type Score struct {
	gorm.Model
	PokemonId     int
	Rank          int
	Points        int
	CompetitionId int
}

func (Score) TableName() string {
	return "score"
}
