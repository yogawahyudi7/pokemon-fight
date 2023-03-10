package seeders

import (
	"pokemon-fight/models"

	"gorm.io/gorm"
)

func LevelSeeder(db *gorm.DB) {

	level1 := models.Level{
		ID:   1,
		Name: "Bos",
	}
	db.Create(&level1)

	level2 := models.Level{
		ID:   2,
		Name: "Operasional",
	}
	db.Create(&level2)

	level3 := models.Level{
		ID:   3,
		Name: "Pengedar",
	}
	db.Create(&level3)
}
