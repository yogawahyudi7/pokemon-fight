package routes

import (
	"pokemon-fight/deliveries/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, pokemon *controllers.PokemonControllers) {
	route := e.Group("/v1")
	// route.Use(middleware.RemoveTrailingSlash())

	route.GET("/pokemons", pokemon.GetAll)

	route.POST("/competition", pokemon.AddCompetition)
	route.GET("/competitions", pokemon.GetCompetitions)

	route.GET("/scores", pokemon.GetScores)

	route.DELETE("/pokemon", pokemon.AddBlackList)
	route.GET("/blacklist", pokemon.GetBlackList)

	route.POST("/season", pokemon.AddSeason)
	route.GET("/seasons", pokemon.GetSeasons)
}
