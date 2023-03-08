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
	GetAll(limit, offset string) (data models.Pokemons, err error)
	GetByUrl(url string) (data models.Pokemon, err error)

	AddCompetition(params models.Competition) (data models.Competition, err error)
	AddScore(params models.Score) (data models.Score, err error)

	AddCompetitionScoreTrx(params models.Competition) (data models.Competition, err error)
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

		err = tx.Create(&score).Error

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

func (pr *PokemonRepositories) GetCompetitions() (data []models.Competition, err error) {

	query := pr.db.Debug()

	err = query.Find(&data).Error

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
