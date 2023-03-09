package seeders

import (
	"pokemon-fight/models"

	"gorm.io/gorm"
)

func CompetitionSeeder(db *gorm.DB) {

	data1 := models.Competition{
		Rank1st:  11,
		Rank2nd:  12,
		Rank3rd:  13,
		Rank4th:  14,
		Rank5th:  15,
		SeasonId: 1,
	}
	db.Create(&data1)

	data2 := models.Competition{
		Rank1st:  21,
		Rank2nd:  22,
		Rank3rd:  23,
		Rank4th:  24,
		Rank5th:  25,
		SeasonId: 2,
	}
	db.Create(&data2)

}
