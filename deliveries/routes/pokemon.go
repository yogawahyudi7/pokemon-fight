package routes

import (
	"pokemon-fight/deliveries/controllers"
	"pokemon-fight/deliveries/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterPathPokemon(e *echo.Group, pokemon *controllers.PokemonControllers) {

	e.GET("/pokemons", pokemon.GetPokemons, echojwt.WithConfig(echojwt.Config(middleware.MiddlewareConfig()))) //OPERASIONAL
	e.GET("/pokemon", pokemon.GetPokemon, echojwt.WithConfig(echojwt.Config(middleware.MiddlewareConfig())))   //OPERASIONAL
	e.POST("/pokemon-upload", pokemon.UploadImagePokemon)                                                      //OPERASIONAL

}
