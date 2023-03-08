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
)

func main() {

	configs := configs.Get()
	db := utils.InitDB(configs)

	utils.InitialMigrate(configs, db)
	seeders.SeasonSeeder(db)

	pokemonRepositories := repositories.NewPokemonRepositories(db)
	pokemonControllers := controllers.NewPokemonControllers(pokemonRepositories)

	//echo package
	e := echo.New()
	middleware.LogMiddleware(e)

	routes.RegisterPath(e, pokemonControllers)

	e.Logger.Fatal(e.Start(":" + configs.Port))

}
