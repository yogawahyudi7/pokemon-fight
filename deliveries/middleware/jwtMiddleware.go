package middleware

import (
	"time"

	"pokemon-fight/constants"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// create token with adding limit time
func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SecretJWT))
}

func ExtractToken(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)

	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := int(claims["userId"].(float64))
		return userId
	}
	return 0
}
