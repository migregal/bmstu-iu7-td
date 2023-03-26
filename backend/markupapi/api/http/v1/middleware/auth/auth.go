package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"markup2/markupapi/api/http/v1/response"
	pkgjwt "markup2/pkg/jwt"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			log.Error("token not found")
			return c.JSON(http.StatusOK, response.Response{Errors: echo.Map{
				"default": "unauthorized",
			}})
		}
		claims, ok := token.Claims.(*pkgjwt.JWTClaims)
		if !ok {
			log.Error("claims not found")
			return c.JSON(http.StatusOK, response.Response{Errors: echo.Map{
				"default": "unauthorized",
			}})
		}

		if claims.ExpiresAt.Before(time.Now()) {
			log.Error("expired")
			return c.JSON(http.StatusOK, response.Response{Errors: echo.Map{
				"default": "unauthorized",
			}})
		}

		c.Set("user_id", claims.UserID)

		return next(c)
	}
}

func OptionalAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			log.Warn("token not found")
			return next(c)
		}

		claims, ok := token.Claims.(*pkgjwt.JWTClaims)
		if !ok {
			log.Error("claims not found")
			return c.JSON(http.StatusOK, response.Response{Errors: echo.Map{
				"default": "unauthorized",
			}})
		}

		if claims.ExpiresAt.Before(time.Now()) {
			log.Error("expired")
			return c.JSON(http.StatusOK, response.Response{Errors: echo.Map{
				"default": "unauthorized",
			}})
		}

		c.Set("user_id", claims.UserID)

		return next(c)
	}
}
