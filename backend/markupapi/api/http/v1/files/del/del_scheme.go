package del

import (
	"markup2/pkg/godraft"
	"net/http"

	"github.com/gothing/draft"
	"github.com/gothing/draft/types"
)

type FilesDelete struct {
	draft.Endpoint
}

type FilesDelParams struct {
	ID types.ID `required:"true" comment:"передается в path"`
}

func (a *FilesDelete) InitEndpointScheme(s *draft.Scheme) {
	s.Project("markup2")

	s.Access(draft.Access.Auth)

	s.Method(draft.Method.POST)

	s.Name("«Удаление файла»")

	s.URL("/api/v1/files/del")

	s.Params(FilesDelParams{
		ID: types.GenID(),
	})

	// 200 OK
	s.Case(godraft.HTTPStatus(http.StatusOK), "Файл существует", func() {
	})

	// 403 OK
	s.Case(godraft.HTTPStatus(http.StatusForbidden), "Файл не принадлежит пользователю", func() {
		s.Params(FilesDelParams{
			ID: types.GenID(),
		})
	})

	// 404 OK
	s.Case(godraft.HTTPStatus(http.StatusNotFound), "Файла не существует", func() {
		id := types.GenID()

		s.Params(FilesDelParams{
			ID: id,
		})

		s.Body(godraft.InvalidCaseBody("id", "not found", id))
	})
}
