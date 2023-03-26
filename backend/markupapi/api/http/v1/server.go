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
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"markup2/markupapi/api/http/v1/auth/login"
	"markup2/markupapi/api/http/v1/auth/logout"
	"markup2/markupapi/api/http/v1/auth/registration"
	"markup2/markupapi/api/http/v1/files/add"
	"markup2/markupapi/api/http/v1/files/del"
	"markup2/markupapi/api/http/v1/files/get"
	"markup2/markupapi/api/http/v1/files/upd"
	"markup2/markupapi/api/http/v1/middleware/auth"
	"markup2/markupapi/core/interactors/files"
	"markup2/markupapi/core/interactors/user"
	"markup2/markupapi/core/ports/repositories"
	filesRepo "markup2/markupapi/repositories/files"
	userRepo "markup2/markupapi/repositories/user"
	pkgjwt "markup2/pkg/jwt"
)

type Config struct {
	Address         string
	GracefulTimeout time.Duration
	UserDB          repositories.UserConfig
	Render          files.Config
	FilesDB         repositories.FilesConfig
	Secret          string
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

	err = s.InitFiles()
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

	login := login.New(user, s.cfg.Secret)
	g.POST("/login", login.Handle)

	cfg := pkgjwt.NewConfig([]byte(s.cfg.Secret), ForceAuthError)

	l := s.Group("/api/v1/auth")
	l.Use(echojwt.WithConfig(cfg))
	l.Use(auth.AuthMiddleware)

	logout := logout.New(user)
	l.POST("/logout", logout.Handle)

	return nil
}

func (s *Server) InitFiles() error {
	filesRepo, err := filesRepo.New(s.cfg.FilesDB)
	if err != nil {
		return fmt.Errorf("failed to init user repo: %w", err)
	}

	files, err := files.New(s.cfg.Render, filesRepo)
	if err != nil {
		return fmt.Errorf("failed to init interactor: %w", err)
	}

	optAuthCfg := pkgjwt.NewConfig([]byte(s.cfg.Secret), IgnoreError)
	optAuth := s.Group("/api/v1/files")
	optAuth.Use(echojwt.WithConfig(optAuthCfg))
	optAuth.Use(auth.OptionalAuthMiddleware)

	get := get.New(files)
	optAuth.GET("/get/:id", get.Handle)
	optAuth.GET("/get", get.Handle)

	authCfg := pkgjwt.NewConfig([]byte(s.cfg.Secret), ForceAuthError)
	authed := s.Group("/api/v1/files")
	authed.Use(echojwt.WithConfig(authCfg))
	authed.Use(auth.AuthMiddleware)

	add := add.New(files)
	authed.POST("/add", add.Handle)

	upd := upd.New(files)
	authed.PUT("/upd/:id", upd.Handle)

	del := del.New(files)
	authed.DELETE("/del/:id", del.Handle)

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
