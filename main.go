package main

import (
	"pokemon-fight/configs"
	"pokemon-fight/deliveries/controllers"
	"pokemon-fight/deliveries/middleware"
	"pokemon-fight/deliveries/routes"
	"pokemon-fight/repositories"
	"pokemon-fight/seeders"
	"pokemon-fight/utils"

	"github.com/labstack/echo/v4"
	middlewares "github.com/labstack/echo/v4/middleware"
)

func main() {

	configs := configs.Get()
	db := utils.InitDB(configs)

	utils.InitialMigrate(configs, db)

	seeders.LevelSeeder(db)
	seeders.SeasonSeeder(db)
	seeders.CompetitionSeeder(db)
	seeders.ScoreSeeder(db)

	userRepositories := repositories.NewUserRepositories(db)
	pokemonRepositories := repositories.NewPokemonRepositories(db)
	competitionRepositories := repositories.NewCompetitionRepositories(db)
	scoreRepositories := repositories.NewScoreRepositories(db)
	seasonRepositories := repositories.NewSeasonRepositories(db)
	blacklistRepositories := repositories.NewBlacklistRepositories(db)

	userControllers := controllers.NewUserControllers(userRepositories)
	pokemonControllers := controllers.NewPokemonControllers(pokemonRepositories, userRepositories)
	competitionControllers := controllers.NewCompetitionControllers(competitionRepositories, pokemonRepositories, seasonRepositories, userRepositories, blacklistRepositories)
	scoreControllers := controllers.NewScoreControllers(blacklistRepositories, userRepositories, pokemonRepositories, seasonRepositories, scoreRepositories)
	seasonControllers := controllers.NewSeasonControllers(blacklistRepositories, userRepositories, pokemonRepositories, seasonRepositories, scoreRepositories)
	blacklistControllers := controllers.NewBlacklistControllers(blacklistRepositories, userRepositories, pokemonRepositories, seasonRepositories)

	//echo package
	e := echo.New()
	middleware.LogMiddleware(e)
	e.Pre(middlewares.RemoveTrailingSlash())

	path := e.Group("/v1")

	routes.RegisterPathUser(path, userControllers)
	routes.RegisterPathPokemon(path, pokemonControllers)
	routes.RegisterPathCompetition(path, competitionControllers)
	routes.RegisterPathScore(path, scoreControllers)
	routes.RegisterPathSeason(path, seasonControllers)
	routes.RegisterPathBlacklist(path, blacklistControllers)

	e.Logger.Fatal(e.Start(":" + configs.Port))

}
