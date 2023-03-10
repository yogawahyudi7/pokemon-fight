package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pokemon-fight/constants"
	"pokemon-fight/models"
	"strings"

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

type PokemonRepositoriesInterface interface {
	//POKEMON
	GetAll(limit, offset string) (data models.Pokemons, err error)
	GetByUrl(url string) (data models.Pokemon, err error)
	GetByString(str interface{}) (data models.Pokemon, err error)
	//TODO: setting login dan search by name
	//SEASON
	AddSeason(params models.Season) (err error)
	GetSeasons() (data []models.Season, err error)
	GetSeasonById(id int) (data models.Season, err error)

	//COMPETITION
	// AddCompetition(params models.Competition) (data models.Competition, err error)
	// AddScore(params models.Score) (data models.Score, err error)
	AddCompetitionScoreTrx(params models.Competition) (data models.Competition, err error)
	GetCompetitions(seasonId, filterScore int) (data []models.Competition, err error)
	GetScores(seasonId int) (data []models.Score, err error)
	// GetCompetitionsScore() (data []models.Competition, err error)

	//BLACKLIST
	AddBlackList(pokemonId int) (err error)
	DeleteScoreById(pokemonId int) (err error)
	GetBlackList(pokemonId int) (data []models.Score, err error)
	GetBlackListById(pokemonId int) (data []models.Blacklist, err error)
}

type PokemonRepositories struct {
	db *gorm.DB
}

func NewPokemonRepositories(db *gorm.DB) *PokemonRepositories {
	return &PokemonRepositories{
		db: db,
	}
}

func (pr *PokemonRepositories) GetAll(limit, offset string) (data models.Pokemons, err error) {

	api := constants.PokemonAPIV2
	path := fmt.Sprintf("pokemon?limit=%v&offset=%v", limit, offset)
	element := []string{api, path}
	apiPath := strings.Join(element, "")

	// fmt.Println("--API PATH--")
	// fmt.Println(apiPath)

	response, err := http.Get(apiPath)

	if err != nil {
		return data, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return data, err
	}

	json.Unmarshal(responseData, &data)

	// fmt.Println("--DATA RESULT--")
	// fmt.Println(data)

	return data, err
}

func (pr *PokemonRepositories) GetByUrl(url string) (data models.Pokemon, err error) {

	apiPath := url

	// fmt.Println("--API PATH--")
	// fmt.Println(apiPath)

	response, err := http.Get(apiPath)

	if err != nil {
		return data, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return data, err
	}

	json.Unmarshal(responseData, &data)

	// fmt.Println("--DATA RESULT--")
	// fmt.Println(data)

	return data, err
}

func (pr *PokemonRepositories) GetByString(str interface{}) (data models.Pokemon, err error) {

	api := constants.PokemonAPIV2
	path := fmt.Sprintf("pokemon/%v", str)
	element := []string{api, path}
	apiPath := strings.Join(element, "")

	// fmt.Println("--API PATH--")
	// fmt.Println(apiPath)

	response, err := http.Get(apiPath)

	if err != nil {
		return data, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return data, err
	}

	json.Unmarshal(responseData, &data)

	// fmt.Println("--DATA RESULT--")
	// fmt.Println(data)

	return data, err
}

func (pr *PokemonRepositories) AddCompetition(params models.Competition) (data models.Competition, err error) {

	data = models.Competition{
		SeasonId: params.SeasonId,
	}

	query := pr.db.Debug()

	err = query.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, err
}

func (pr *PokemonRepositories) AddScore(params models.Score) (data models.Score, err error) {

	data = models.Score{
		PokemonId:     params.PokemonId,
		CompetitionId: params.CompetitionId,
		Rank1stCount:  params.Rank1stCount,
		Rank2ndCount:  params.Rank2ndCount,
		Rank3rdCount:  params.Rank3rdCount,
		Rank4thCount:  params.Rank4thCount,
		Rank5thCount:  params.Rank5thCount,
		Points:        params.Points,
	}

	query := pr.db.Debug()

	err = query.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, err
}

func (pr *PokemonRepositories) AddCompetitionScoreTrx(params models.Competition) (data models.Competition, err error) {

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

func (pr *PokemonRepositories) GetCompetitions(seasonId, filterScore int) (data []models.Competition, err error) {

	query := pr.db.Debug().Preload("DataScore")

	if seasonId != 0 {
		query = query.Where("season_id = ?", seasonId)
	}

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

	query = query.Order("season_id DESC")

	query = query.Find(&data)

	err = query.Error

	if query.Error != nil {
		return data, err
	}

	return data, err
}

func (pr *PokemonRepositories) GetCompetitionsScore() (data []models.Competition, err error) {

	// var
	selectedField := []string{
		"sc.pokemon_id as pokemon_id",
	}
	query := pr.db.Debug()

	query = query.Table(models.Competition{}.TableName() + " AS co")

	query = query.Select(selectedField)

	query = query.Joins("JOIN " + models.Score{}.TableName() + " AS sc on sc.CompetiionId = co.Id")

	err = query.Find(&data).Error

	if query.Error != nil {
		return data, err
	}

	return data, err
}

func (pr *PokemonRepositories) GetSeasons() (data []models.Season, err error) {

	query := pr.db.Debug()
	err = query.Find(&data).Error

	if query.Error != nil {
		return data, err
	}

	return data, err
}

func (pr *PokemonRepositories) GetSeasonById(id int) (data models.Season, err error) {

	query := pr.db.Debug()
	query = query.Where("id = ?", id)
	err = query.Find(&data).Error

	if query.Error != nil {
		return data, err
	}

	return data, err
}

func (pr *PokemonRepositories) GetScores(seasonId int) (data []models.Score, err error) {

	selectedField := []string{
		"pokemon_id",
		"SUM(rank1st_count) AS rank1st_count",
		"SUM(rank2nd_count) AS rank2nd_count",
		"SUM(rank3rd_count) AS rank3rd_count",
		"SUM(rank4th_count) AS rank4th_count",
		"SUM(rank5th_count) AS rank5th_count",
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

func (pr *PokemonRepositories) AddBlackList(pokemonId int) (err error) {

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

func (pr *PokemonRepositories) DeleteScoreById(pokemonId int) (err error) {

	data := []models.Score{}

	query := pr.db.Debug()

	query = query.Where("pokemon_id = ?", pokemonId)

	err = query.Delete(&data).Error
	if err != nil {
		return err
	}

	return err
}

func (pr *PokemonRepositories) GetBlackList(pokemonId int) (data []models.Score, err error) {

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

func (pr *PokemonRepositories) GetBlackListById(pokemonId int) (data []models.Blacklist, err error) {

	query := pr.db.Debug()

	query = query.Where("pokemon_id = ?", pokemonId)

	err = query.Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, err
}

func (pr *PokemonRepositories) AddSeason(params models.Season) (err error) {

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
