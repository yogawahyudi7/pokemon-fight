package routes

import (
	"pokemon-fight/deliveries/controllers"
	"pokemon-fight/deliveries/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterPathScore(e *echo.Group, score *controllers.ScoreControllers) {
	e.GET("/scores", score.GetScores, echojwt.WithConfig(echojwt.Config(middleware.MiddlewareConfig()))) //BOS

}
