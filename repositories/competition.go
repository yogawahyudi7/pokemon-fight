package repositories

import (
	"fmt"
	"pokemon-fight/models"

	"gorm.io/gorm"
)

type CompetitionScoreTrx struct {
	Id       int
	Rank1st  int
	Rank2nd  int
	Rank3rd  int
	Rank4th  int
	Rank5th  int
	SeasonId int
}

type CompetitionRepositoriesInterface interface {
	AddCompetitionScoreTrx(params models.Competition) (data models.Competition, err error)
	GetCompetitions(seasonId int) (data []models.Competition, err error)
}

type CompetitionRepositories struct {
	db *gorm.DB
}

func NewCompetitionRepositories(db *gorm.DB) *CompetitionRepositories {
	return &CompetitionRepositories{
		db: db,
	}
}

func (pr *CompetitionRepositories) AddCompetitionScoreTrx(params models.Competition) (data models.Competition, err error) {

	competition := models.Competition{
		Rank1st:  params.Rank1st,
		Rank2nd:  params.Rank2nd,
		Rank3rd:  params.Rank3rd,
		Rank4th:  params.Rank4th,
		Rank5th:  params.Rank5th,
		SeasonId: params.SeasonId,
	}

	tx := pr.db.Begin().Debug()

	err = tx.Create(&competition).Error

	tx.SavePoint("sp1")

	if err != nil {
		tx.RollbackTo("sp1")
		return data, err
	}

	n := 5
	pokemonId := []int{
		params.Rank1st,
		params.Rank2nd,
		params.Rank3rd,
		params.Rank4th,
		params.Rank5th,
	}

	for i := 0; i < n; i++ {
		competitionId := competition.ID
		pokemonId := pokemonId[i]
		points := n - i

		rank1stCount := 0
		rank2ndCount := 0
		rank3rdCount := 0
		rank4thCount := 0
		rank5thCount := 0

		switch i {
		case 0:
			rank1stCount = 1
		case 1:
			rank2ndCount = 1
		case 2:
			rank3rdCount = 1
		case 3:
			rank4thCount = 1
		case 4:
			rank5thCount = 1
		}

		score := models.Score{
			PokemonId:     pokemonId,
			CompetitionId: int(competitionId),
			Rank1stCount:  rank1stCount,
			Rank2ndCount:  rank2ndCount,
			Rank3rdCount:  rank3rdCount,
			Rank4thCount:  rank4thCount,
			Rank5thCount:  rank5thCount,
			Points:        points,
		}

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
		err = tx.Select(selectedFilend).Create(&score).Error

		spCount := 2 + i
		sp := fmt.Sprintf("sp%v", spCount)
		tx.SavePoint(sp)

		if err != nil {
			tx.RollbackTo("sp1")
			return data, err
		}

	}

	tx.Commit()

	data = competition

	return data, err
}

func (pr *CompetitionRepositories) GetCompetitions(seasonId int) (data []models.Competition, err error) {

	query := pr.db.Debug().Preload("DataScore")

	if seasonId != 0 {
		query = query.Where("season_id = ?", seasonId)
	}

	query = query.Order("season_id DESC")

	query = query.Find(&data)

	err = query.Error

	if query.Error != nil {
		return data, err
	}

	return data, err
}
