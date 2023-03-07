package models

import "gorm.io/gorm"

type Blacklist struct {
	gorm.Model
	PokemonId int
}

func (Blacklist) TableName() string {
	return "blacklist"
}
