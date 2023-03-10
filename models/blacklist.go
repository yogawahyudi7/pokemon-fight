package models

import "gorm.io/gorm"

type Blacklist struct {
	gorm.Model
	PokemonId int `gorm:"unique"`

	// DataPokemon Pokemon `gorm:"foreignKey:PokemonId"`
}

func (Blacklist) TableName() string {
	return "blacklist"
}
