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

type CompetitionControllers struct {
	Competition repositories.CompetitionRepositoriesInterface
	Pokemon     repositories.PokemonRepositoriesInterface
	User        repositories.UserRepositoriesInterface
	Blacklist   repositories.BlacklistRepositoriesInterface
	Season      repositories.SeasonRepositoriesInterface
}

func NewCompetitionControllers(
	competition repositories.CompetitionRepositoriesInterface,
	pokemon repositories.PokemonRepositoriesInterface,
	season repositories.SeasonRepositoriesInterface,
	user repositories.UserRepositoriesInterface,
	blacklist repositories.BlacklistRepositoriesInterface,
) *CompetitionControllers {
	return &CompetitionControllers{
		Competition: competition,
		Pokemon:     pokemon,
		User:        user,
		Blacklist:   blacklist,
		Season:      season,
	}
}

func (pc CompetitionControllers) AddCompetition(ctx echo.Context) error {
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

	rank1st := ctx.FormValue("rank_1st")
	rank2nd := ctx.FormValue("rank_2nd")
	rank3rd := ctx.FormValue("rank_3rd")
	rank4th := ctx.FormValue("rank_4th")
	rank5th := ctx.FormValue("rank_5th")
	seasonId := ctx.FormValue("season_id")

	//VALIDASI NUMERIC
	formValue := []string{
		rank1st,
		rank2nd,
		rank3rd,
		rank4th,
		rank5th,
	}
	for i, vData := range formValue {
		if vData == "" {
			return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Parameter [rank...] Tidak Boleh Kosong."))
		}

		err := validate.Var(vData, "number")
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Parameter [rank...] Hanya Boleh Disi Dengan Angka."))
		}

		//CHECK DUPLCIATE ID INPUT
		for j := 1 + i; j < len(formValue); j++ {
			if formValue[i] == formValue[j] {
				return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Tidak Boleh Menginputkan Id Yang Sama Pada Parameter [rank...]."))
			}
		}
	}

	//VALIDASI SEASON ID
	if seasonId == "" {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Parameter Season Id Tidak Boleh Kosong."))
	}
	err = validate.Var(seasonId, "number")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Parameter Season Id Hanya Boleh Disi Dengan Angka."))
	}

	// fmt.Println("rank1st", rank1st)

	rank1stInt, _ := strconv.Atoi(rank1st)
	rank2ndInt, _ := strconv.Atoi(rank2nd)
	rank3rdInt, _ := strconv.Atoi(rank3rd)
	rank4thInt, _ := strconv.Atoi(rank4th)
	rank5thInt, _ := strconv.Atoi(rank5th)
	seasonIdInt, _ := strconv.Atoi(seasonId)

	listPokemons := []int{
		rank1stInt,
		rank2ndInt,
		rank3rdInt,
		rank4thInt,
		rank5thInt,
	}
	for _, vData := range listPokemons {
		//CHECKING AVAILABLE POKEMON ID
		availablePokemons, err := pc.Pokemon.GetPokemon(vData)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
		}

		fmt.Println(availablePokemons)
		if availablePokemons.Id == 0 {
			responseMessage := fmt.Sprintf("Maaf, Pokemon Dengan Id [%v] Tidak Ditemukan.", vData)
			return ctx.JSON(http.StatusNotFound, response.NotFound(responseMessage))

		}
		//CHECK BLACK LIST
		blacklist, err := pc.Blacklist.GetBlackListById(vData)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
		}

		fmt.Println(blacklist)
		if len(blacklist) > 0 {
			responseMessage := fmt.Sprintf("Maaf, Pokemon Dengan Id [%v] Terdaftar Dalam Blacklist.", vData)
			return ctx.JSON(http.StatusBadRequest, response.BadRequest(responseMessage))

		}
	}

	params := models.Competition{
		Rank1st:  rank1stInt,
		Rank2nd:  rank2ndInt,
		Rank3rd:  rank3rdInt,
		Rank4th:  rank4thInt,
		Rank5th:  rank5thInt,
		SeasonId: seasonIdInt,
	}

	season, err := pc.Season.GetSeasonById(seasonIdInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}
	if season.ID == 0 {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Season Id Tidak Ditemukan"))
	}

	conpetition, err := pc.Competition.AddCompetitionScoreTrx(params)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	competition := common.CompetitionData{
		Id:       int(conpetition.ID),
		Rank1st:  conpetition.Rank1st,
		Rank2nd:  conpetition.Rank2nd,
		Rank3rd:  conpetition.Rank3rd,
		Rank4th:  conpetition.Rank4th,
		Rank5th:  conpetition.Rank5th,
		SeasonId: conpetition.SeasonId,
	}
	return ctx.JSON(http.StatusOK, response.Saved(competition))
}

func (pc CompetitionControllers) GetCompetitions(ctx echo.Context) error {
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

	dataCompetition, err := pc.Competition.GetCompetitions(seasonIdInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	result := []common.DataCompetition{}

	for _, vData := range dataCompetition {

		id := vData.ID
		pokemon1st, err := pc.Pokemon.GetPokemon(vData.Rank1st)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("pokemon1st :"+err.Error()))
		}

		pokemon2nd, err := pc.Pokemon.GetPokemon(vData.Rank2nd)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("pokemon2nd :"+err.Error()))
		}

		pokemon3rd, err := pc.Pokemon.GetPokemon(vData.Rank3rd)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("pokemon3rd :"+err.Error()))
		}

		pokemon4th, err := pc.Pokemon.GetPokemon(vData.Rank4th)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("pokemon4th :"+err.Error()))
		}

		pokemon5th, err := pc.Pokemon.GetPokemon(vData.Rank5th)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("pokemon5th :"+err.Error()))
		}

		season, err := pc.Season.GetSeasonById(vData.SeasonId)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("season :"+err.Error()))
		}

		data := common.DataCompetition{
			Id: int(id),
			Rank1st: common.Pokemon{
				Id:   pokemon1st.Id,
				Name: pokemon1st.Name,
			},
			Rank2nd: common.Pokemon{
				Id:   pokemon2nd.Id,
				Name: pokemon2nd.Name,
			},
			Rank3rd: common.Pokemon{
				Id:   pokemon3rd.Id,
				Name: pokemon3rd.Name,
			},
			Rank4th: common.Pokemon{
				Id:   pokemon4th.Id,
				Name: pokemon4th.Name,
			},
			Rank5th: common.Pokemon{
				Id:   pokemon5th.Id,
				Name: pokemon5th.Name,
			},
			Season: common.Season{
				Id:        int(season.ID),
				Name:      season.Name,
				StartDate: season.StartDate.Format(constants.LayoutYMD),
				EndDate:   season.EndDate.Format(constants.LayoutYMD),
			},
		}

		result = append(result, data)

	}

	if len(result) < 1 {
		return ctx.JSON(http.StatusNotFound, response.NotFound(""))
	}

	return ctx.JSON(http.StatusOK, response.Found(result))
}
