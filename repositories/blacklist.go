package repositories

import (
	"pokemon-fight/models"

	"gorm.io/gorm"
)

type BlacklistRepositoriesInterface interface {
	//BLACKLIST
	AddBlackList(pokemonId int) (err error)
	DeleteScoreById(pokemonId int) (err error)
	GetBlackList(pokemonId int) (data []models.Score, err error)
	GetBlackListById(pokemonId int) (data []models.Blacklist, err error)
}

type BlacklistRepositories struct {
	db *gorm.DB
}

func NewBlacklistRepositories(db *gorm.DB) *BlacklistRepositories {
	return &BlacklistRepositories{
		db: db,
	}
}

func (pr *BlacklistRepositories) AddBlackList(pokemonId int) (err error) {

	data := models.Blacklist{
		PokemonId: pokemonId,
	}

	query := pr.db.Debug()

	err = query.Create(&data).Error
	if err != nil {
		return err
	}

	return err
}

func (pr *BlacklistRepositories) DeleteScoreById(pokemonId int) (err error) {

	data := []models.Score{}

	query := pr.db.Debug()

	query = query.Where("pokemon_id = ?", pokemonId)

	err = query.Delete(&data).Error
	if err != nil {
		return err
	}

	return err
}

func (pr *BlacklistRepositories) GetBlackList(pokemonId int) (data []models.Score, err error) {

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

	query = query.Select(selectedField)

	// query = query.Table(models.Score{})

	if pokemonId != 0 {
		query = query.Where("pokemon_id = ?", pokemonId)
	}

	query = query.Group("pokemon_id")

	query = query.Order("SUM(points) DESC")

	err = query.Unscoped().Where("deleted_at IS NOT NULL").Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, err
}

func (pr *BlacklistRepositories) GetBlackListById(pokemonId int) (data []models.Blacklist, err error) {

	query := pr.db.Debug()

	query = query.Where("pokemon_id = ?", pokemonId)

	err = query.Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, err
}
