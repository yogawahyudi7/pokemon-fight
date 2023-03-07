package main

import (
	"pokemon-fight/configs"
	"pokemon-fight/deliveries/controllers"
	"pokemon-fight/deliveries/middleware"
	"pokemon-fight/deliveries/routes"
	"pokemon-fight/repositories"
	"pokemon-fight/utils"

	"github.com/labstack/echo/v4"
)

func main() {

	configs := configs.Get()
	db := utils.InitDB(configs)
	pokemonRepositories := repositories.NewPokemonRepositories(db)
	pokemonControllers := controllers.NewPokemonControllers(pokemonRepositories)

	//echo package
	e := echo.New()
	middleware.LogMiddleware(e)

	routes.RegisterPath(e, pokemonControllers)
	// data := pokemonRepositories
	// data.GetAll("3", "0")
	// data.GetByUrl("https://pokeapi.co/api/v2/pokemon/1/")
	// fmt.Println(db)
	// fmt.Println("good")
	//"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	e.Logger.Fatal(e.Start(":" + configs.Port))

}
