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

}
