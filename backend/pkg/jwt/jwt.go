package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	Login string `json:"login"`
	ID    uint64 `json:"id"`
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
