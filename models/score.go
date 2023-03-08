package models

import "gorm.io/gorm"

type Score struct {
	gorm.Model
	PokemonId     int
	CompetitionId int
	Rank          int
	Points        int

	// DataPokemon     Pokemon     `gorm:"foreignKey:PokemonId"`
	DataCompetition Competition `gorm:"foreignKey:CompetitionId"`
}

func (Score) TableName() string {
	return "score"
}
