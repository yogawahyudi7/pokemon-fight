package routes

import (
	"pokemon-fight/deliveries/controllers"
	"pokemon-fight/deliveries/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterPathBlacklist(e *echo.Group, blacklist *controllers.BlacklistControllers) {

	e.DELETE("/pokemon", blacklist.AddBlackList, echojwt.WithConfig(echojwt.Config(middleware.MiddlewareConfig()))) //BOS
	e.GET("/blacklist", blacklist.GetBlackList, echojwt.WithConfig(echojwt.Config(middleware.MiddlewareConfig())))  //BOS

}
