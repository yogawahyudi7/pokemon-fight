package routes

import (
	"pokemon-fight/deliveries/controllers"
	"pokemon-fight/deliveries/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterPathUser(e *echo.Group, user *controllers.UserControllers) {
	e.POST("/bos/register", user.RegisterBos)                                                        //Register Bos
	e.POST("/operasional/register", user.RegisterOperasional)                                        //Register Operasional
	e.POST("/pengedar/register", user.RegisterPengedar)                                              //Register Pengedar
	e.POST("/login", user.Login)                                                                     //Login
	e.PUT("/logout", user.Logout, echojwt.WithConfig(echojwt.Config(middleware.MiddlewareConfig()))) //Login

}
