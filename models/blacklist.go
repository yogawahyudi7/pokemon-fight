package models

import "gorm.io/gorm"

type Blacklist struct {
	gorm.Model
	PokemonId int

	// DataPokemon Pokemon `gorm:"foreignKey:PokemonId"`
}

func (Blacklist) TableName() string {
	return "blacklist"
}
