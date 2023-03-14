package routes

import (
	"pokemon-fight/deliveries/controllers"
	"pokemon-fight/deliveries/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterPathCompetition(e *echo.Group, competition *controllers.CompetitionControllers) {

	e.POST("/competition", competition.AddCompetition, echojwt.WithConfig(echojwt.Config(middleware.MiddlewareConfig())))  //PENGEDAR
	e.GET("/competitions", competition.GetCompetitions, echojwt.WithConfig(echojwt.Config(middleware.MiddlewareConfig()))) //BOS

}
