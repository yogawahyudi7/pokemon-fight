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

type PokemonRepositoriesInterface interface {
	//POKEMON
	GetPokemons(limit, offset string) (data models.Pokemons, err error)
	GetByUrl(url string) (data models.Pokemon, err error)
	GetPokemon(str interface{}) (data models.Pokemon, err error)
}

type PokemonRepositories struct {
	db *gorm.DB
}

func NewPokemonRepositories(db *gorm.DB) *PokemonRepositories {
	return &PokemonRepositories{
		db: db,
	}
}

func (pr *PokemonRepositories) GetPokemons(limit, offset string) (data models.Pokemons, err error) {

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

func (pr *PokemonRepositories) GetPokemon(str interface{}) (data models.Pokemon, err error) {

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
