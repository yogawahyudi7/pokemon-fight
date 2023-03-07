package controllers

import (
	"fmt"
	"net/http"
	"pokemon-fight/helpers"
	"pokemon-fight/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (r Response) Saved(data interface{}) Response {
	return Response{
		Code:    200,
		Message: "Data Berhasil Disimpan.",
		Data:    data,
	}
}

func (r Response) Found(data interface{}) Response {
	return Response{
		Code:    200,
		Message: "Data Ditemukan.",
		Data:    data,
	}
}

func (r Response) BadRequest(message string) Response {
	return Response{
		Code:    400,
		Message: message,
		Data:    nil,
	}
}

func (r Response) NotFound() Response {
	return Response{
		Code:    404,
		Message: "Data Tidak Ditemukan.",
		Data:    nil,
	}
}

func (r Response) InternalServerError() Response {
	return Response{
		Code:    500,
		Message: "Maaf, Server Sedang Dalam Perbaikan Cobalah Beberapa Saat Lagi.",
		Data:    nil,
	}
}

type PokemonData struct {
	Id             int      `json:"id"`
	Name           string   `json:"name"`
	Abilities      []string `json:"abilities"`
	Height         int      `json:"height"`
	Weight         int      `json:"weight"`
	BaseExperience int      `json:"base_experience"`
}

var validate = validator.New()

type PokemonControllers struct {
	Repositories repositories.PokemonRepositoriesInterface
}

func NewPokemonControllers(repositories repositories.PokemonRepositoriesInterface) *PokemonControllers {
	return &PokemonControllers{Repositories: repositories}
}

func (pc PokemonControllers) GetAll(ctx echo.Context) error {
	response := Response{}

	page := ctx.QueryParam("page")
	fmt.Println("PAGE :", page)
	if page == "" {
		page = "1"
	}
	err := validate.Var(page, "number")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Parameter [page] Hanya Boleh Disi Dengan Angka."))

	}

	limit := ctx.QueryParam("limit")
	fmt.Println("LIMIT :", limit)
	if limit == "" {
		limit = "10"
	}
	err = validate.Var(limit, "number")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Parameter [limit] Hanya Boleh Disi Dengan Angka."))

	}

	offset := helpers.Pagination(limit, page)

	fmt.Println("-- DATA OFFSET --")
	fmt.Println("=", offset)
	dataGetAll, err := pc.Repositories.GetAll(limit, offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError())
	}

	// fmt.Println("-- DATA GET ALL --")
	// fmt.Println("=", dataGetAll)

	pokemons := []PokemonData{}

	for _, data := range dataGetAll.Results {
		url := data.Url

		pokemon, err := pc.Repositories.GetByUrl(url)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError())
		}

		var abilities []string

		for _, vData := range pokemon.Abilities {
			data := vData.Ability.Name

			abilities = append(abilities, data)
		}

		data := PokemonData{
			Id:             pokemon.Id,
			Name:           pokemon.Name,
			Abilities:      abilities,
			Height:         pokemon.Height,
			Weight:         pokemon.Weight,
			BaseExperience: pokemon.BaseExperience,
		}
		pokemons = append(pokemons, data)
	}

	if len(pokemons) < 1 {
		return ctx.JSON(http.StatusNotFound, response.NotFound())
	}
	// fmt.Println("-- DATA POKEMON --")
	// fmt.Println("=", pokemons)

	return ctx.JSON(http.StatusOK, response.Found(pokemons))
}
