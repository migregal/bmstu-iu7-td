package api

import (
	"markup2/markupapi/api/http"
	"markup2/markupapi/api/http/godraft"
	v1 "markup2/markupapi/api/http/v1"
	"markup2/markupapi/config"

	"github.com/labstack/gommon/log"
)

type API struct {
	http     http.Server
	draftAPI godraft.Documentation
}

func New(cfg config.Config) API {
	s := API{}

	s.http = v1.New(v1.Config(cfg.HTTP))

	godraft.Init()
	s.draftAPI = godraft.New(godraft.Config(cfg.Docs))

	return s
}

func (s *API) Run() {
	errs := make(chan error, 1)

	go func() {
		errs <- s.http.ListenAndServe()
	}()
	go func() {
		errs <- s.draftAPI.ListenAndServe()
	}()

	err := <-errs

	log.Warn("Terminating application", err)
}
