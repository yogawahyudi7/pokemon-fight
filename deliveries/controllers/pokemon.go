package controllers

import (
	"fmt"
	"log"
	"net/http"
	"pokemon-fight/constants"
	"pokemon-fight/deliveries/common"
	"pokemon-fight/deliveries/middleware"
	"pokemon-fight/helpers"
	"pokemon-fight/models"
	"pokemon-fight/repositories"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

type PokemonControllers struct {
	Repositories repositories.PokemonRepositoriesInterface
}

func NewPokemonControllers(repositories repositories.PokemonRepositoriesInterface) *PokemonControllers {
	return &PokemonControllers{Repositories: repositories}
}

func (pc PokemonControllers) GetPokemons(ctx echo.Context) error {
	response := common.Response{}

	//check otorisasi
	logged_in_user_id := middleware.ExtractToken(ctx)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := pc.Repositories.GetUserById(logged_in_user_id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}
	if user.LevelID != 2 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Operasional")
	}

	page := ctx.QueryParam("page")
	fmt.Println("PAGE :", page)
	if page == "" {
		page = "1"
	}
	err = validate.Var(page, "number")
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
	dataGetAll, err := pc.Repositories.GetPokemons(limit, offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	// fmt.Println("-- DATA GET ALL --")
	// fmt.Println("=", dataGetAll)

	pokemons := []common.PokemonData{}

	for _, data := range dataGetAll.Results {
		url := data.Url

		pokemon, err := pc.Repositories.GetByUrl(url)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
		}

		var abilities []string

		for _, vData := range pokemon.Abilities {
			data := vData.Ability.Name

			abilities = append(abilities, data)
		}

		data := common.PokemonData{
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
		return ctx.JSON(http.StatusNotFound, response.NotFound(""))
	}
	// fmt.Println("-- DATA POKEMON --")
	// fmt.Println("=", pokemons)

	return ctx.JSON(http.StatusOK, response.Found(pokemons))
}

func (pc PokemonControllers) GetPokemon(ctx echo.Context) error {
	response := common.Response{}

	//check otorisasi
	logged_in_user_id := middleware.ExtractToken(ctx)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := pc.Repositories.GetUserById(logged_in_user_id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}
	if user.LevelID != 2 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Operasional")
	}

	search := ctx.QueryParam("search")
	if search == "" {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Parameter Search Tidak Boleh Kosong."))
	}

	data, err := pc.Repositories.GetPokemon(search)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	fmt.Println("-- DATA GET ALL --")
	fmt.Println("=", data)

	abilities := []string{}
	for _, vData := range data.Abilities {
		data := vData.Ability.Name

		abilities = append(abilities, data)
	}

	types := []string{}
	for _, vData := range data.Types {
		data := vData.Type.Name

		types = append(types, data)
	}

	stats := []common.Stats{}
	for _, vData := range data.Stats {
		data := common.Stats{
			Name:     vData.Stat.Name,
			BaseStat: vData.BaseStat,
			Effort:   vData.Effort,
		}

		stats = append(stats, data)
	}

	pokemon := common.PokemonData{
		Id:             data.Id,
		Name:           data.Name,
		Abilities:      abilities,
		Height:         data.Height,
		Weight:         data.Weight,
		Types:          types,
		Stats:          stats,
		BaseExperience: data.BaseExperience,
	}

	if data.Id == 0 {
		return ctx.JSON(http.StatusNotFound, response.NotFound(""))
	}
	// fmt.Println("-- DATA POKEMON --")
	// fmt.Println("=", pokemons)

	return ctx.JSON(http.StatusOK, response.Found(pokemon))
}

func (pc PokemonControllers) AddCompetition(ctx echo.Context) error {
	response := common.Response{}

	logged_in_user_id := middleware.ExtractToken(ctx)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := pc.Repositories.GetUserById(logged_in_user_id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}
	if user.LevelID != 3 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Pengedar")
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
		seasonId,
	}
	for _, vData := range formValue {
		if vData == "" {
			return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Parameter [rank...] Tidak Boleh Kosong."))
		}

		err := validate.Var(vData, "number")
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, response.BadRequest("Maaf, Parameter [rank...] Hanya Boleh Disi Dengan Angka."))
		}
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
		availablePokemons, err := pc.Repositories.GetPokemon(vData)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
		}

		fmt.Println(availablePokemons)
		if availablePokemons.Id == 0 {
			responseMessage := fmt.Sprintf("Maaf, Pokemon Dengan Id [%v] Tidak Ditemukan.", vData)
			return ctx.JSON(http.StatusNotFound, response.NotFound(responseMessage))

		}
		//CHECK BLACK LIST
		blacklist, err := pc.Repositories.GetBlackListById(vData)
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

	season, err := pc.Repositories.GetSeasonById(seasonIdInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}
	if season.ID == 0 {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest("Season Id Tidak Ditemukan"))
	}

	conpetition, err := pc.Repositories.AddCompetitionScoreTrx(params)
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

func (pc PokemonControllers) GetCompetitions(ctx echo.Context) error {
	response := common.Response{}

	//check otorisasi
	logged_in_user_id := middleware.ExtractToken(ctx)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := pc.Repositories.GetUserById(logged_in_user_id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}
	if user.LevelID != 1 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Bos")
	}

	seasonId := ctx.QueryParam("season_id")

	seasonIdInt, _ := strconv.Atoi(seasonId)

	dataCompetition, err := pc.Repositories.GetCompetitions(seasonIdInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	result := []common.DataCompetition{}

	for _, vData := range dataCompetition {

		id := vData.ID
		pokemon1st, err := pc.Repositories.GetPokemon(vData.Rank1st)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("pokemon1st :"+err.Error()))
		}

		pokemon2nd, err := pc.Repositories.GetPokemon(vData.Rank2nd)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("pokemon2nd :"+err.Error()))
		}

		pokemon3rd, err := pc.Repositories.GetPokemon(vData.Rank3rd)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("pokemon3rd :"+err.Error()))
		}

		pokemon4th, err := pc.Repositories.GetPokemon(vData.Rank4th)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("pokemon4th :"+err.Error()))
		}

		pokemon5th, err := pc.Repositories.GetPokemon(vData.Rank5th)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("pokemon5th :"+err.Error()))
		}

		season, err := pc.Repositories.GetSeasonById(vData.SeasonId)
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

func (pc PokemonControllers) GetScores(ctx echo.Context) error {
	response := common.Response{}

	//check otorisasi
	logged_in_user_id := middleware.ExtractToken(ctx)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := pc.Repositories.GetUserById(logged_in_user_id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}
	if user.LevelID != 1 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Bos")
	}

	seasonId := ctx.QueryParam("season_id")

	seasonIdInt, _ := strconv.Atoi(seasonId)

	dataScore, err := pc.Repositories.GetScores(seasonIdInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	result := []common.DataScores{}

	for _, vData := range dataScore {

		id := vData.ID
		pokmeonId := vData.PokemonId
		pokemon, err := pc.Repositories.GetPokemon(pokmeonId)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
		}

		seasonId := vData.SeasonId
		season := models.Season{}
		if seasonId != 0 {
			season, err = pc.Repositories.GetSeasonById(int(seasonId))
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

func (pc PokemonControllers) AddBlackList(ctx echo.Context) error {
	response := common.Response{}

	//check otorisasi
	logged_in_user_id := middleware.ExtractToken(ctx)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := pc.Repositories.GetUserById(logged_in_user_id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}
	if user.LevelID != 1 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Bos")
	}

	pokemonId := ctx.QueryParam("pokemon_id")

	pokemonIdInt, _ := strconv.Atoi(pokemonId)

	dataPokemon, err := pc.Repositories.GetPokemon(pokemonIdInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}
	if dataPokemon.Id == 0 {
		responseMessage := fmt.Sprintf("Maaf, Pokemon Dengan Id [%v] Tidak Ditemukan.", pokemonIdInt)
		return ctx.JSON(http.StatusNotFound, response.NotFound(responseMessage))
	}

	err = pc.Repositories.AddBlackList(pokemonIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return ctx.JSON(http.StatusBadRequest, response.BadRequest(fmt.Sprintf("Maaf, Id Pokemon %v Sudah Terdaftar Dalam Blacklist", pokemonIdInt)))
		}
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	err = pc.Repositories.DeleteScoreById(pokemonIdInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.Saved(nil))
}

func (pc PokemonControllers) GetBlackList(ctx echo.Context) error {
	response := common.Response{}

	//check otorisasi
	logged_in_user_id := middleware.ExtractToken(ctx)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := pc.Repositories.GetUserById(logged_in_user_id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}
	if user.LevelID != 1 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Bos")
	}

	pokemonId := ctx.FormValue("pokemon_id")

	pokemonIdInt, _ := strconv.Atoi(pokemonId)

	data, err := pc.Repositories.GetBlackList(pokemonIdInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	result := []common.DataScores{}

	for _, vData := range data {

		id := vData.ID
		pokmeonId := vData.PokemonId
		pokemon, err := pc.Repositories.GetPokemon(pokmeonId)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
		}

		seasonId := vData.SeasonId
		season := models.Season{}
		if seasonId != 0 {
			season, err = pc.Repositories.GetSeasonById(int(seasonId))
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

func (pc PokemonControllers) AddSeason(ctx echo.Context) error {
	response := common.Response{}

	logged_in_user_id := middleware.ExtractToken(ctx)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := pc.Repositories.GetUserById(logged_in_user_id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}
	if user.LevelID != 3 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Pengedar")
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

	_, errStr = helpers.ValidatorDate(endDate)
	if errStr != "" {
		return ctx.JSON(http.StatusBadRequest, response.BadRequest(errStr))
	}

	startDateParse, _ := time.Parse(constants.LayoutYMD, startDate)
	endDateParse, _ := time.Parse(constants.LayoutYMD, endDate)

	paramsSeason := models.Season{
		Name:      name,
		StartDate: startDateParse,
		EndDate:   endDateParse,
	}

	err = pc.Repositories.AddSeason(paramsSeason)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError(err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.Saved(nil))
}

func (pc PokemonControllers) GetSeasons(ctx echo.Context) error {
	response := common.Response{}

	logged_in_user_id := middleware.ExtractToken(ctx)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := pc.Repositories.GetUserById(logged_in_user_id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}
	if user.LevelID != 3 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Pengedar")
	}

	seasonData, err := pc.Repositories.GetSeasons()
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

//AUTH

type UserOutput struct {
	ID    uint   `json:"id"`
	Level string `json:"level"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserOutput1 struct {
	ID    uint   `json:"id"`
	Level string `json:"level"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func EncryptPwd(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func (pc PokemonControllers) RegisterBos(c echo.Context) error {
	//get user's input
	input_user := models.User{}
	input_user.LevelID = 1
	c.Bind(&input_user)

	//check is data nil?
	if input_user.Email == "" || input_user.Password == "" || input_user.Name == "" || input_user.LevelID == 0 {
		return c.JSON(http.StatusBadRequest, "Maaf, Dimohon Untuk Melengkapi Semua Data.")
	}

	//check is email exists?
	is_email_exists, _ := pc.Repositories.CheckEmail(input_user.Email)
	if is_email_exists {
		return c.JSON(http.StatusBadRequest, "Maaf, Email Sudah Pernah Terdaftar.")
	}

	//encrypt pass user
	convert_pwd := []byte(input_user.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	input_user.Password = hashed_pwd //set new pass

	//create new user
	user, err := pc.Repositories.Register(input_user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal Menginputkan Data")
	}

	//get level name
	level, err := pc.Repositories.GetLevel(int(user.LevelID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}

	//customize output
	output := UserOutput{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
	}

	return c.JSON(http.StatusOK, output)
}

func (pc PokemonControllers) RegisterOperasional(c echo.Context) error {
	//get user's input
	input_user := models.User{}
	input_user.LevelID = 2
	c.Bind(&input_user)

	//check is data nil?
	if input_user.Email == "" || input_user.Password == "" || input_user.Name == "" || input_user.LevelID == 0 {
		return c.JSON(http.StatusBadRequest, "Maaf, Dimohon Untuk Melengkapi Semua Data.")
	}

	//check is email exists?
	is_email_exists, _ := pc.Repositories.CheckEmail(input_user.Email)
	if is_email_exists {
		return c.JSON(http.StatusBadRequest, "Email already exist")
	}

	//encrypt pass user
	convert_pwd := []byte(input_user.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	input_user.Password = hashed_pwd //set new pass

	//create new user
	user, err := pc.Repositories.Register(input_user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal Menginputkan Data")
	}

	//get level name
	level, err := pc.Repositories.GetLevel(int(user.LevelID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}

	//customize output
	output := UserOutput{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
	}

	return c.JSON(http.StatusOK, output)
}

func (pc PokemonControllers) RegisterPengedar(c echo.Context) error {
	//get user's input
	input_user := models.User{}
	input_user.LevelID = 3
	c.Bind(&input_user)

	//check is data nil?
	if input_user.Email == "" || input_user.Password == "" || input_user.Name == "" || input_user.LevelID == 0 {
		return c.JSON(http.StatusBadRequest, "Maaf, Dimohon Untuk Melengkapi Semua Data.")
	}

	//check is email exists?
	is_email_exists, _ := pc.Repositories.CheckEmail(input_user.Email)
	if is_email_exists {
		return c.JSON(http.StatusBadRequest, "Email already exist")
	}

	//encrypt pass user
	convert_pwd := []byte(input_user.Password) //convert pass from string to byte
	hashed_pwd := EncryptPwd(convert_pwd)
	input_user.Password = hashed_pwd //set new pass

	//create new user
	user, err := pc.Repositories.Register(input_user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal Menginputkan Data")
	}

	//get level name
	level, err := pc.Repositories.GetLevel(int(user.LevelID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}

	//customize output
	output := UserOutput{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
	}

	return c.JSON(http.StatusOK, output)
}

func (pc PokemonControllers) Login(c echo.Context) error {
	//get user's input
	input_user := models.User{}
	c.Bind(&input_user)

	//check is data nil?
	if input_user.Email == "" || input_user.Password == "" {
		return c.JSON(http.StatusBadRequest, "Maaf, Dimohon Untuk Melengkapi Semua Data.")
	}

	//compare password on form with db
	get_pwd, x := pc.Repositories.GetPassword(input_user.Email) //get password
	if x != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}
	err := bcrypt.CompareHashAndPassword([]byte(get_pwd), []byte(input_user.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "User Unauthorized. Email or Password not equal")
	}

	//login
	user, err := pc.Repositories.Login(input_user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}

	//get level name
	level, err := pc.Repositories.GetLevel(int(user.LevelID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}

	//customize output
	output := UserOutput1{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
		Token: user.Token,
	}

	return c.JSON(http.StatusOK, output)
}

func (pc PokemonControllers) GetProfile(c echo.Context) error {
	//get id user login
	logged_in_user_id := middleware.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}

	//get customer by id
	user, err := pc.Repositories.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}

	//get level name
	level, err := pc.Repositories.GetLevel(int(user.LevelID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}

	//customize output
	output := UserOutput{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
	}

	return c.JSON(http.StatusOK, output)
}

func (pc PokemonControllers) GetProfileTesting() echo.HandlerFunc {
	return pc.GetProfile
}

func (pc PokemonControllers) Logout(c echo.Context) error {
	//get id user login
	logged_in_user_id := middleware.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}

	//get customer by id
	user, err := pc.Repositories.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}
	user.Token = ""
	c.Bind(&user)
	customer_updated, err := pc.Repositories.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}

	//get level name
	level, err := pc.Repositories.GetLevel(int(customer_updated.LevelID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal Mendapatkan Data")
	}

	//customize output
	output := UserOutput1{
		ID:    user.ID,
		Level: level.Name,
		Email: user.Email,
		Name:  user.Name,
		Token: user.Token,
	}

	return c.JSON(http.StatusOK, output)
}
