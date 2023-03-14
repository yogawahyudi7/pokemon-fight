package repositories

import (
	"pokemon-fight/models"

	"gorm.io/gorm"
)

type SeasonRepositoriesInterface interface {
	//SEASON
	AddSeason(params models.Season) (err error)
	GetSeasons() (data []models.Season, err error)
	GetSeasonById(id int) (data models.Season, err error)
}

type SeasonRepositories struct {
	db *gorm.DB
}

func NewSeasonRepositories(db *gorm.DB) *SeasonRepositories {
	return &SeasonRepositories{
		db: db,
	}
}

func (pr *SeasonRepositories) GetSeasons() (data []models.Season, err error) {

	query := pr.db.Debug()
	err = query.Find(&data).Error

	if query.Error != nil {
		return data, err
	}

	return data, err
}

func (pr *SeasonRepositories) GetSeasonById(id int) (data models.Season, err error) {

	query := pr.db.Debug()
	query = query.Where("id = ?", id)
	err = query.Find(&data).Error

	if query.Error != nil {
		return data, err
	}

	return data, err
}

func (pr *SeasonRepositories) AddSeason(params models.Season) (err error) {

	data := models.Season{
		Name:      params.Name,
		StartDate: params.StartDate,
		EndDate:   params.EndDate,
	}

	query := pr.db.Debug()

	err = query.Create(&data).Error
	if err != nil {
		return err
	}

	return err
}
