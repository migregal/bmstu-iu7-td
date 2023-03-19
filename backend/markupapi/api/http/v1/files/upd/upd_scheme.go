package upd

import (
	"markup2/pkg/godraft"
	"net/http"

	"github.com/gothing/draft"
	"github.com/gothing/draft/types"
)

type FilesUpdate struct {
	draft.Endpoint
}

type FilesUpdParams struct {
	ID   types.ID `required:"true" comment:"передается в path"`
	File string   `required:"true" comment:"файл в формате md"`
}

func (a *FilesUpdate) InitEndpointScheme(s *draft.Scheme) {
	s.Project("markup2")

	s.Access(draft.Access.Auth)

	s.Method(draft.Method.POST)

	s.Name("«Обновление файла»")

	s.URL("/api/v1/files/upd/")

	s.Params(FilesUpdParams{
		ID: types.GenID(),
	})

	// 200 OK
	s.Case(godraft.HTTPStatus(http.StatusOK), "Файл существует", func() {
		s.Body(FilesUpdParams{
			ID:   types.GenID(),
			File: "index.md",
		})
	})

	// 403 OK
	s.Case(godraft.HTTPStatus(http.StatusForbidden), "Файл не принадлежит пользователю", func() {
		s.Params(FilesUpdParams{
			ID: types.GenID(),
		})
	})

	// 404 OK
	s.Case(godraft.HTTPStatus(http.StatusNotFound), "Файла не существует", func() {
		id := types.GenID()

		s.Params(FilesUpdParams{
			ID: id,
		})

		s.Body(godraft.InvalidCaseBody("id", "not found", id))
	})
}
