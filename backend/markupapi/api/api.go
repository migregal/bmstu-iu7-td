package api

import (
	"fmt"

	"github.com/labstack/gommon/log"

	"markup2/markupapi/api/http"
	v1 "markup2/markupapi/api/http/v1"
	"markup2/markupapi/config"
	filesInteractor "markup2/markupapi/core/interactors/files"
	"markup2/markupapi/core/ports/repositories"
)

type API struct {
	http     http.Server
}

func New(cfg config.Config) (API, error) {
	s := API{}

	var err error
	s.http, err = v1.New(v1.Config{
		Address:         cfg.HTTP.Address,
		GracefulTimeout: cfg.HTTP.GracefulTimeout,
		UserDB:          repositories.UserConfig(cfg.UserDB),
		Render:          filesInteractor.Config(cfg.Render),
		FilesDB:         repositories.FilesConfig(cfg.FilesDB),
	})
	if err != nil {
		return API{}, fmt.Errorf("failed to init http api: %w", err)
	}

	return s, nil
}

func (s *API) Run() {
	errs := make(chan error, 1)

	go func() {
		errs <- s.http.ListenAndServe()
	}()

	err := <-errs

	log.Warn("Terminating application", err)
}
