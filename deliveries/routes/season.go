package routes

import (
	"pokemon-fight/deliveries/controllers"
	"pokemon-fight/deliveries/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterPathSeason(e *echo.Group, season *controllers.SeasonControllers) {

	e.GET("/seasons", season.GetSeasons, echojwt.WithConfig(echojwt.Config(middleware.MiddlewareConfig()))) //PENGEDAR
	e.POST("/season", season.AddSeason, echojwt.WithConfig(echojwt.Config(middleware.MiddlewareConfig())))  //PENGEDAR

}
