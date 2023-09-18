package middleware

import (
	"pokemon-fight/constants"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type CustomClaims struct {
	Authorized bool
	UserId     int
	jwt.RegisteredClaims
}

func MiddlewareConfig() echojwt.Config {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(CustomClaims)
		},
		SigningKey: []byte(constants.SecretJWT),
	}

	return config
}

// create token with adding limit time
func CreateToken(userId int) (string, error) {
	// claims := jwt.MapClaims{}
	// claims["authorized"] = true
	// claims["userId"] = userId
	// claims["exp"] = time.Now().Add(time.Hour * 3).Unix()

	claims := CustomClaims{
		Authorized: true,
		UserId:     userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SecretJWT))
}

func ExtractToken(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)

	if user.Valid {
		// claims := user.Claims.(jwt.MapClaims)
		// userId := int(claims["userId"].(float64))
		claims := user.Claims.(*CustomClaims)
		userId := claims.UserId
		return userId
	}
	return 0
}
