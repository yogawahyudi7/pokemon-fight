package seeders

import (
	"pokemon-fight/models"
	"time"

	"gorm.io/gorm"
)

func SeasonSeeder(db *gorm.DB) {

	season1 := models.Season{
		Name:      "First Season",
		StartDate: time.Now().AddDate(0, -1, 0),
		EndDate:   time.Now().AddDate(0, 0, 0),
	}
	db.Create(&season1)

	season2 := models.Season{
		Name:      "Second Season",
		StartDate: time.Now().AddDate(0, 1, 0),
		EndDate:   time.Now().AddDate(0, 2, 0),
	}
	db.Create(&season2)

	season3 := models.Season{
		Name:      "Third Season",
		StartDate: time.Now().AddDate(0, 2, 0),
		EndDate:   time.Now().AddDate(0, 3, 0),
	}
	db.Create(&season3)
}
