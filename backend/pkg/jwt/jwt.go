package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewConfig(
	secret []byte,
	errHandler func(c echo.Context, err error) error,
) echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTClaims)
		},
		ErrorHandler:           errHandler,
		SigningKey:             secret,
		ContinueOnIgnoredError: true,
	}
}

type JWTClaims struct {
	Login  string `json:"login"`
	UserID uint64 `json:"id"`
	jwt.RegisteredClaims
}

func NewToken(secret []byte, login string, id uint64) (string, error) {
	claims := &JWTClaims{
		login,
		id,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}
