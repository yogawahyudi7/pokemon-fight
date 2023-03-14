package repositories

import (
	"fmt"
	"pokemon-fight/models"

	"gorm.io/gorm"
)

type ScoreRepositoriesInterface interface {
	GetScores(seasonId int) (data []models.Score, err error)
}

type ScoreRepositories struct {
	db *gorm.DB
}

func NewScoreRepositories(db *gorm.DB) *ScoreRepositories {
	return &ScoreRepositories{
		db: db,
	}
}

func (pr *ScoreRepositories) GetScores(seasonId int) (data []models.Score, err error) {

	selectedField := []string{
		"pokemon_id",
		"SUM(rank_1st_count) AS rank_1st_count",
		"SUM(rank_2nd_count) AS rank_2nd_count",
		"SUM(rank_3rd_count) AS rank_3rd_count",
		"SUM(rank_4th_count) AS rank_4th_count",
		"SUM(rank_5th_count) AS rank_5th_count",
		"SUM(points) AS total_points",
	}

	query := pr.db.Debug()

	if seasonId != 0 {
		selectedField = append(selectedField, "season_id")

		query = query.Where("season_id = ?", seasonId)

		query = query.Group("season_id")
	}

	query = query.Select(selectedField)

	query = query.Table(models.Score{}.TableName() + " AS sc")

	query = query.Joins("JOIN " + models.Competition{}.TableName() + " AS co on co.ID = sc.competition_id")

	query = query.Joins("JOIN " + models.Season{}.TableName() + " AS se on se.ID = co.season_id")

	query = query.Group("pokemon_id")

	query = query.Order("SUM(points) DESC")

	query = query.Find(&data)

	// if filterScore == 1 {
	// 	query = query.Preload("DataScore", func(db *gorm.DB) *gorm.DB {
	// 		return db.Order("points asc")
	// 	}).Find(&data)
	// } else if filterScore == 2 {
	// 	query = query.Preload("DataScore", func(db *gorm.DB) *gorm.DB {
	// 		return db.Order("points desc")
	// 	}).Find(&data)
	// } else {
	// query = query.Find(&data)
	// }

	err = query.Error

	if query.Error != nil {
		return data, err
	}
	fmt.Println(&data)
	return data, err
}
