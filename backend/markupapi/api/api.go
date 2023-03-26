package api

import (
	"fmt"

	"github.com/labstack/gommon/log"

	"markup2/markupapi/api/http"
	v1 "markup2/markupapi/api/http/v1"
	"markup2/markupapi/api/http/v1/files"
	"markup2/markupapi/config"
	filesInteractor "markup2/markupapi/core/interactors/files"
	"markup2/markupapi/core/ports/repositories"
	"markup2/pkg/godraft"
)

type API struct {
	http     http.Server
	draftAPI *godraft.Documentation
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

	if cfg.Debug {
		s.draftAPI = setupDocumentation(cfg)
	}

	return s, nil
}

func setupDocumentation(cfg config.Config) *godraft.Documentation {
	godraft.Init()
	draftAPI := godraft.New(godraft.Config(cfg.Docs))
	draftAPI.Add(files.Service)

	return draftAPI
}

func (s *API) Run() {
	errs := make(chan error, 2)

	go func() {
		errs <- s.http.ListenAndServe()
	}()
	go func() {
		if s.draftAPI != nil {
			errs <- s.draftAPI.ListenAndServe()
		}
	}()

	err := <-errs

	log.Warn("Terminating application", err)
}
