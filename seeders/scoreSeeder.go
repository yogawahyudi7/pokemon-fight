package seeders

import (
	"pokemon-fight/models"

	"gorm.io/gorm"
)

func ScoreSeeder(db *gorm.DB) {

	selectedFilend := []string{
		"PokemonId",
		"CompetitionId",
		"Rank1stCount",
		"Rank2ndCount",
		"Rank3rdCount",
		"Rank4thCount",
		"Rank5thCount",
		"Points",
	}

	data1 := models.Score{
		PokemonId:     11,
		CompetitionId: 1,
		Rank1stCount:  1,
		Rank2ndCount:  0,
		Rank3rdCount:  0,
		Rank4thCount:  0,
		Rank5thCount:  0,
		Points:        5,
	}
	db.Select(selectedFilend).Create(&data1)

	data2 := models.Score{
		PokemonId:     12,
		CompetitionId: 1,
		Rank1stCount:  0,
		Rank2ndCount:  1,
		Rank3rdCount:  0,
		Rank4thCount:  0,
		Rank5thCount:  0,
		Points:        4,
	}
	db.Select(selectedFilend).Create(&data2)

	data3 := models.Score{
		PokemonId:     13,
		CompetitionId: 1,
		Rank1stCount:  0,
		Rank2ndCount:  0,
		Rank3rdCount:  1,
		Rank4thCount:  0,
		Rank5thCount:  0,
		Points:        3,
	}
	db.Select(selectedFilend).Create(&data3)

	data4 := models.Score{
		PokemonId:     14,
		CompetitionId: 1,
		Rank1stCount:  0,
		Rank2ndCount:  0,
		Rank3rdCount:  0,
		Rank4thCount:  1,
		Rank5thCount:  0,
		Points:        2,
	}
	db.Select(selectedFilend).Create(&data4)

	data5 := models.Score{
		PokemonId:     15,
		CompetitionId: 1,
		Rank1stCount:  0,
		Rank2ndCount:  0,
		Rank3rdCount:  0,
		Rank4thCount:  0,
		Rank5thCount:  1,
		Points:        1,
	}
	db.Select(selectedFilend).Create(&data5)

	data21 := models.Score{
		PokemonId:     21,
		CompetitionId: 2,
		Rank1stCount:  1,
		Rank2ndCount:  0,
		Rank3rdCount:  0,
		Rank4thCount:  0,
		Rank5thCount:  0,
		Points:        5,
	}
	db.Select(selectedFilend).Create(&data21)

	data22 := models.Score{
		PokemonId:     22,
		CompetitionId: 2,
		Rank1stCount:  0,
		Rank2ndCount:  1,
		Rank3rdCount:  0,
		Rank4thCount:  0,
		Rank5thCount:  0,
		Points:        4,
	}
	db.Select(selectedFilend).Create(&data22)

	data23 := models.Score{
		PokemonId:     23,
		CompetitionId: 2,
		Rank1stCount:  0,
		Rank2ndCount:  0,
		Rank3rdCount:  1,
		Rank4thCount:  0,
		Rank5thCount:  0,
		Points:        3,
	}
	db.Select(selectedFilend).Create(&data23)

	data24 := models.Score{
		PokemonId:     24,
		CompetitionId: 2,
		Rank1stCount:  0,
		Rank2ndCount:  0,
		Rank3rdCount:  0,
		Rank4thCount:  1,
		Rank5thCount:  0,
		Points:        2,
	}
	db.Select(selectedFilend).Create(&data24)

	data25 := models.Score{
		PokemonId:     25,
		CompetitionId: 2,
		Rank1stCount:  0,
		Rank2ndCount:  0,
		Rank3rdCount:  0,
		Rank4thCount:  0,
		Rank5thCount:  1,
		Points:        1,
	}
	db.Select(selectedFilend).Create(&data25)
}
