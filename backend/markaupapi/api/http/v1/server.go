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
)

type Config struct {
	Port            uint
	GracefulTimeout time.Duration
}

type Server struct {
	*echo.Echo
	cfg Config
}

func New(cfg Config) *Server {
	s := Server{cfg: cfg}
	s.Echo = echo.New()

	s.InitMiddleware()

	s.InitHealthCheck()

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

func (s *Server) ListenAndServe() error {
	go func() {
		err := s.Start(fmt.Sprintf(":%d", s.cfg.Port))
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
