package api

import (
	"markup2/markaupapi/api/http"
	v1 "markup2/markaupapi/api/http/v1"
	"markup2/markaupapi/config"

	"github.com/labstack/gommon/log"
)

type API struct {
	http http.Server
}

func New(cfg config.Config) API {
	s := API{}

	s.http = v1.New(v1.Config(cfg.HTTP))

	return s
}

func (s *API) Run() {
	errs := make(chan error, 1)

	go func() {
		errs <- s.http.ListenAndServe()
	}()

	<-errs

	log.Warn("Terminating application")
}
