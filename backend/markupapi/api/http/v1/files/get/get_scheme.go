package get

import (
	"markup2/pkg/godraft"
	"net/http"

	"github.com/gothing/draft"
	"github.com/gothing/draft/types"
)

type FilesGet struct {
	draft.Endpoint
}

type FilesGetParams struct {
	ID types.ID `required:"true" comment:"передается в path"`
	Format string `comment:"формат файла для ответа: md, html, plain"`
}

func (a *FilesGet) InitEndpointScheme(s *draft.Scheme) {
	s.Project("markup2")

	s.Access(draft.Access.All)

	s.Method(draft.Method.GET)

	s.Name("«Получение файла»")

	s.URL("/api/v1/files/get")

	s.Params(FilesGetParams{
		ID: types.GenID(),
	})

	// 200 OK
	s.Case(godraft.HTTPStatus(http.StatusOK), "Файл существует", func() {
		s.Body("index.html")
	})

	// 400 OK
	s.Case(godraft.HTTPStatus(http.StatusNotFound), "Файл не существует", func() {
		s.Params(FilesGetParams{
			ID: 0,
		})
		s.Body(godraft.InvalidCaseBody("id", "invalid", ""))
	})
}
