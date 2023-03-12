package routes

import (
	"pokemon-fight/deliveries/controllers"
	"pokemon-fight/deliveries/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, pokemon *controllers.PokemonControllers) {
	route := e.Group("/v1")
	// route.Use(middleware.RemoveTrailingSlash())

	//-----------Bos
	route.POST("/bos/register", pokemon.RegisterBos)

	//-----------Operasional
	route.POST("/operasional/register", pokemon.RegisterOperasional)

	//-----------Pengedar
	route.POST("/pengedar/register", pokemon.RegisterPengedar)

	//LOGIN
	route.POST("/login", pokemon.Login)

	auth := route.Group("")
	auth.Use(echojwt.WithConfig(echojwt.Config(middleware.MiddlewareConfig())))

	//LOGOUT
	auth.PUT("/logout", pokemon.Logout)

	auth.GET("/pokemons", pokemon.GetPokemons) //OPERASIONAL
	auth.GET("/pokemon", pokemon.GetPokemon)   //OPERASIONAL

	auth.GET("/seasons", pokemon.GetSeasons)          //PENGEDAR
	auth.POST("/season", pokemon.AddSeason)           //PENGEDAR
	auth.POST("/competition", pokemon.AddCompetition) //PENGEDAR

	auth.GET("/competitions", pokemon.GetCompetitions) //BOS
	auth.GET("/scores", pokemon.GetScores)             //BOS

	auth.DELETE("/pokemon", pokemon.AddBlackList) //BOS
	auth.GET("/blacklist", pokemon.GetBlackList)  //BOS

}
