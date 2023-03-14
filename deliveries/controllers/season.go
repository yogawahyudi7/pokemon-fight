package controllers

import (
	"net/http"
	"pokemon-fight/constants"
	"pokemon-fight/deliveries/common"
	"pokemon-fight/deliveries/middleware"
	"pokemon-fight/helpers"
	"pokemon-fight/models"
	"pokemon-fight/repositories"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type SeasonControllers struct {
	Blacklist repositories.BlacklistRepositoriesInterface
	User      repositories.UserRepositoriesInterface
	Pokemon   repositories.PokemonRepositoriesInterface
	Season    repositories.SeasonRepositoriesInterface
	Score     repositories.ScoreRepositoriesInterface
}

func NewSeasonControllers(
	blacklist repositories.BlacklistRepositoriesInterface,
	user repositories.UserRepositoriesInterface,
	pokemon repositories.PokemonRepositoriesInterface,
	season repositories.SeasonRepositoriesInterface,
	score repositories.ScoreRepositoriesInterface,
) *SeasonControllers {
	return &SeasonControllers{
		Blacklist: blacklist,
		User:      user,
		Pokemon:   pokemon,
		Season:    season,
		Score:     score,
	}
}

func (pc SeasonControllers) AddSeason(ctx echo.Context) error {
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
	if user.LevelId != 3 {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(3))
	}

	name := ctx.FormValue("name")
	startDate := ctx.FormValue("start_date")
	endDate := ctx.FormValue("end_date")

	_, errStr := helpers.ValidatorGeneralName(name)
	if errStr != "" {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest(errStr))
	}

	_, errStr = helpers.ValidatorDate(startDate)
	if errStr != "" {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest(errStr))
	}

	// fmt.Println("==", code)
	// fmt.Println("==", errStr)

	_, errStr = helpers.ValidatorDate(endDate)
	if errStr != "" {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest(errStr))
	}

	startDateParse, _ := time.Parse(constants.LayoutYMD, startDate)
	endDateParse, _ := time.Parse(constants.LayoutYMD, endDate)

	errTrue := startDateParse.After(endDateParse)
	if errTrue {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Tanggal Mulai Tidak Boleh Melampaui Tanggal Selesai Season."))
	}

	paramsSeason := models.Season{
		Name:      name,
		StartDate: startDateParse,
		EndDate:   endDateParse,
	}

	err = pc.Season.AddSeason(paramsSeason)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Nama Season Sudah Pernah Terdaftar, Mohon Menggunakan Nama Lain."))
		}
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.Saved(nil))
}

func (pc SeasonControllers) GetSeasons(ctx echo.Context) error {
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
	if user.LevelId != 3 {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(3))
	}

	seasonData, err := pc.Season.GetSeasons()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	data := []common.Season{}

	for _, vData := range seasonData {

		id := vData.ID
		name := vData.Name
		startDate := vData.StartDate.Format(constants.LayoutYMD)
		endDate := vData.EndDate.Format(constants.LayoutYMD)

		seasonData := common.Season{
			Id:        int(id),
			Name:      name,
			StartDate: startDate,
			EndDate:   endDate,
		}

		data = append(data, seasonData)

	}

	if len(data) < 1 {
		return ctx.JSON(http.StatusNotFound, response.NotFound(""))
	}

	return ctx.JSON(http.StatusOK, response.Found(data))

}
