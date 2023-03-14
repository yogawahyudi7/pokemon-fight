package controllers

import (
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

type ScoreControllers struct {
	Blacklist repositories.BlacklistRepositoriesInterface
	User      repositories.UserRepositoriesInterface
	Pokemon   repositories.PokemonRepositoriesInterface
	Season    repositories.SeasonRepositoriesInterface
	Score     repositories.ScoreRepositoriesInterface
}

func NewScoreControllers(
	blacklist repositories.BlacklistRepositoriesInterface,
	user repositories.UserRepositoriesInterface,
	pokemon repositories.PokemonRepositoriesInterface,
	season repositories.SeasonRepositoriesInterface,
	score repositories.ScoreRepositoriesInterface,
) *ScoreControllers {
	return &ScoreControllers{
		Blacklist: blacklist,
		User:      user,
		Pokemon:   pokemon,
		Season:    season,
		Score:     score,
	}
}

func (pc ScoreControllers) GetScores(ctx echo.Context) error {
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

	seasonId := ctx.QueryParam("season_id")

	seasonIdInt, _ := strconv.Atoi(seasonId)

	dataScore, err := pc.Score.GetScores(seasonIdInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	result := []common.DataScores{}

	for _, vData := range dataScore {

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
