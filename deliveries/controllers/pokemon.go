package controllers

import (
	"fmt"
	"net/http"
	"pokemon-fight/deliveries/common"
	"pokemon-fight/deliveries/middleware"
	"pokemon-fight/facade"
	"pokemon-fight/helpers"
	"pokemon-fight/repositories"
	"pokemon-fight/services"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

type PokemonControllers struct {
	Repositories repositories.PokemonRepositoriesInterface
	Auth         repositories.UserRepositoriesInterface
	Services     services.ServiceInterface
}

func NewPokemonControllers(repositories repositories.PokemonRepositoriesInterface, auth repositories.UserRepositoriesInterface, services services.ServiceInterface) *PokemonControllers {
	return &PokemonControllers{
		Repositories: repositories,
		Auth:         auth,
		Services:     services,
	}
}

func (pc PokemonControllers) GetPokemons(ctx echo.Context) error {
	response := common.Response{}

	//check otorisasi
	tokenUserId := middleware.ExtractToken(ctx)
	if tokenUserId == 0 {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(1))
	}
	user, err := pc.Auth.GetUserById(tokenUserId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(6))
		}
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("User Id : "+err.Error()))
	}
	if user.Token == "" {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(2))
	}
	if user.LevelId != 2 {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(3))
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
	tokenUserId := middleware.ExtractToken(ctx)
	if tokenUserId == 0 {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(1))
	}
	user, err := pc.Auth.GetUserById(tokenUserId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(6))
		}
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("User Id : "+err.Error()))
	}
	if user.Token == "" {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(2))
	}
	if user.LevelId != 2 {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(3))
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

func (pc PokemonControllers) UploadImagePokemon(ctx echo.Context) error {
	response := common.Response{}

	//check otorisasi
	tokenUserId := middleware.ExtractToken(ctx)
	if tokenUserId == 0 {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(1))
	}
	user, err := pc.Auth.GetUserById(tokenUserId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(6))
		}
		return ctx.JSON(http.StatusInternalServerError, response.InternalServerError("User Id : "+err.Error()))
	}
	if user.Token == "" {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(2))
	}
	if user.LevelId != 2 {
		return ctx.JSON(http.StatusUnauthorized, response.Unauthorized(3))
	}

	// Maksimum ukuran file yang diizinkan (dalam byte)
	maxFileSize := int64(10 * 1024 * 1024) // Contoh: 10MB

	// Dapatkan Id pada form value
	idStr := ctx.FormValue("id")

	id, _ := strconv.Atoi(idStr)

	// Dapatkan file dari permintaan
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Gagal mendapatkan file dari form"})
	}

	// Validasi ukuran file
	if file.Size > maxFileSize {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Ukuran file terlalu besar"})
	}

	// Buka file yang diunggah
	src, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal membuka file"})
	}
	defer src.Close()

	pokemonId := strconv.Itoa(id)

	filePatch, err := facade.ServerUploadFile(pokemonId, file)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error upload to server": err.Error()})
	}

	err = pc.Services.UploadImagePokemonGCS(pokemonId, filePatch)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error upload to gcs": err.Error()})
	}

	err = facade.ServerRemoveFile(filePatch)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error remove file": err.Error()})
	}

	return ctx.JSON(http.StatusOK, response.Saved(nil))
}
