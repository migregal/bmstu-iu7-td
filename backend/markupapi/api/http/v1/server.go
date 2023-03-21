package v1

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	userRepo "markup2/markupapi/core/adapters/repositories/user"
	"markup2/markupapi/core/interactors/user"
	"markup2/markupapi/core/ports/repositories"
)

type Config struct {
	Address         string
	GracefulTimeout time.Duration
	UserDB          repositories.UserConfig
}

type Server struct {
	*echo.Echo
	cfg  Config
	user user.Interactor
}

func New(cfg Config) *Server {
	s := Server{cfg: cfg}
	s.Echo = echo.New()

	s.InitMiddleware()

	s.InitHealthCheck()

	s.InitAuth()

	return &s
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

	_ = user.New(userRepo)

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
