package v1

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo-contrib/prometheus"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"markup2/markupapi/api/http/v1/auth/login"
	"markup2/markupapi/api/http/v1/auth/logout"
	"markup2/markupapi/api/http/v1/auth/registration"
	"markup2/markupapi/api/http/v1/response"
	"markup2/markupapi/core/interactors/user"
	"markup2/markupapi/core/ports/repositories"
	userRepo "markup2/markupapi/repositories/user"
	pkgjwt "markup2/pkg/jwt"
)

type Config struct {
	Address         string
	GracefulTimeout time.Duration
	UserDB          repositories.UserConfig
}

type Server struct {
	*echo.Echo
	cfg Config
}

func New(cfg Config) (*Server, error) {
	s := Server{cfg: cfg}
	s.Echo = echo.New()

	s.InitMiddleware()

	s.InitHealthCheck()

	err := s.InitAuth()
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (s *Server) InitMiddleware() {
	s.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	s.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		TargetHeader: echo.HeaderXRequestID,
	}))

	p := prometheus.NewPrometheus("markup2", nil)
	p.Use(s.Echo)
}

func (s *Server) InitHealthCheck() {
	s.GET("/readiness", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	s.GET("/liveness", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
}

func (s *Server) InitAuth() error {
	userRepo, err := userRepo.New(s.cfg.UserDB)
	if err != nil {
		return fmt.Errorf("failed to init user repo: %w", err)
	}

	user := user.New(userRepo)

	g := s.Group("/api/v1/auth")
	registration := registration.New(user)
	g.POST("/registration", registration.Handle)

	login := login.New(user)
	g.POST("/login", login.Handle)

	cfg := pkgjwt.NewConfig(
		[]byte("secret"),
		func(c echo.Context, err error) error {
			log.Errorf("unauthorized: %v", err)
			return c.JSON(http.StatusOK, response.Response{Errors: echo.Map{
				"default": "unauthorized",
			}})
		},
	)

	l := s.Group("/api/v1/auth")
	l.Use(echojwt.WithConfig(cfg))
	l.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, ok := c.Get("user").(*jwt.Token)
			if !ok {
				log.Errorf("token not found")
				return c.JSON(http.StatusOK, response.Response{Errors: echo.Map{
					"default": "unauthorized",
				}})
			}
			claims, ok := token.Claims.(*pkgjwt.JWTClaims)
			if !ok {
				log.Errorf("claims not found")
				return c.JSON(http.StatusOK, response.Response{Errors: echo.Map{
					"default": "unauthorized",
				}})
			}

			if claims.ExpiresAt.Before(time.Now()) {
				log.Errorf("expired")
				return c.JSON(http.StatusOK, response.Response{Errors: echo.Map{
					"default": "unauthorized",
				}})
			}

			return next(c)
		}
	})

	logout := logout.New(user)
	l.POST("/logout", logout.Handle)

	return nil
}

func (s *Server) ListenAndServe() error {
	go func() {
		err := s.Start(s.cfg.Address)
		if err != nil && err != http.ErrServerClosed {
			s.Logger.Fatal("shutting down the server")
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	timeout := s.cfg.GracefulTimeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return s.Shutdown(ctx)
}
