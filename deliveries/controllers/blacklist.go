package controllers

import (
	"fmt"
	"net/http"
	"pokemon-fight/constants"
	"pokemon-fight/deliveries/common"
	"pokemon-fight/deliveries/middleware"
	"pokemon-fight/models"
	"pokemon-fight/repositories"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type BlacklistControllers struct {
	Blacklist repositories.BlacklistRepositoriesInterface
	User      repositories.UserRepositoriesInterface
	Pokemon   repositories.PokemonRepositoriesInterface
	Season    repositories.SeasonRepositoriesInterface
}

func NewBlacklistControllers(
	blacklist repositories.BlacklistRepositoriesInterface,
	user repositories.UserRepositoriesInterface,
	pokemon repositories.PokemonRepositoriesInterface,
	season repositories.SeasonRepositoriesInterface,
) *BlacklistControllers {
	return &BlacklistControllers{
		Blacklist: blacklist,
		User:      user,
		Pokemon:   pokemon,
		Season:    season,
	}
}

func (pc BlacklistControllers) AddBlackList(ctx echo.Context) error {
	response := common.Response{}

	//check otorisasi
	tokenUserId := middleware.ExtractToken(ctx)
	if tokenUserId == 0 {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(1))
	}
	user, err := pc.User.GetUserById(tokenUserId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(6))
		}
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("User Id : "+err.Error()))
	}
	if user.Token == "" {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(2))
	}
	if user.LevelId != 1 {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(3))
	}

	pokemonId := ctx.QueryParam("pokemon_id")

	pokemonIdInt, _ := strconv.Atoi(pokemonId)

	dataPokemon, err := pc.Pokemon.GetPokemon(pokemonIdInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}
	if dataPokemon.Id == 0 {
		responseMessage := fmt.Sprintf("Maaf, Pokemon Dengan Id [%v] Tidak Ditemukan.", pokemonIdInt)
		return ctx.JSON(http.StatusNotFound, response.NotFound(responseMessage))
	}

	err = pc.Blacklist.AddBlackList(pokemonIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return ctx.JSON(http.StatusBadRequest, response.BadRequest(fmt.Sprintf("Maaf, Id Pokemon %v Sudah Terdaftar Dalam Blacklist", pokemonIdInt)))
		}
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	err = pc.Blacklist.DeleteScoreById(pokemonIdInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.Saved(nil))
}

func (pc BlacklistControllers) GetBlackList(ctx echo.Context) error {
	response := common.Response{}

	///check otorisasi
	tokenUserId := middleware.ExtractToken(ctx)
	if tokenUserId == 0 {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(1))
	}
	user, err := pc.User.GetUserById(tokenUserId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(6))
		}
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("User Id : "+err.Error()))
	}
	if user.Token == "" {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(2))
	}
	if user.LevelId != 1 {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(3))
	}

	pokemonId := ctx.FormValue("pokemon_id")

	pokemonIdInt, _ := strconv.Atoi(pokemonId)

	data, err := pc.Blacklist.GetBlackList(pokemonIdInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	result := []common.DataScores{}

	for _, vData := range data {

		id := vData.ID
		pokmeonId := vData.PokemonId
		pokemon, err := pc.Pokemon.GetPokemon(pokmeonId)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
		}

		seasonId := vData.SeasonId
		season := models.Season{}
		if seasonId != 0 {
			season, err = pc.Season.GetSeasonById(int(seasonId))
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
			}
		}

		pokemonData := common.Pokemon{
			Id:   pokemon.Id,
			Name: pokemon.Name,
		}

		startDate := season.StartDate.Format(constants.LayoutYMD)
		endDate := season.EndDate.Format(constants.LayoutYMD)

		// if startDate == "0001-01-01" {
		// 	startDate = ""
		// }

		// if endDate == "0001-01-01" {
		// 	endDate = ""
		// }

		seasonData := common.Season{
			Id:        int(season.ID),
			Name:      season.Name,
			StartDate: startDate,
			EndDate:   endDate,
		}

		data := common.DataScores{
			Id:           int(id),
			Pokemon:      pokemonData,
			Rank1stCount: vData.Rank1stCount,
			Rank2ndCount: vData.Rank2ndCount,
			Rank3rdCount: vData.Rank3rdCount,
			Rank4thCount: vData.Rank4thCount,
			Rank5thCount: vData.Rank5thCount,
			TotalPoints:  vData.TotalPoints,
			Season:       seasonData,
		}
		if season.ID == 0 {
			data.Season = "All Season"
		}

		result = append(result, data)

	}

	if len(result) < 1 {
		return ctx.JSON(http.StatusNotFound, response.NotFound(""))
	}

	return ctx.JSON(http.StatusOK, response.Found(result))

}
