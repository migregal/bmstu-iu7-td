package add

import (
	"fmt"
	"markup2/pkg/godraft"
	"net/http"

	"github.com/gothing/draft"
	"github.com/gothing/draft/types"
)

type FilesAdd struct {
	draft.Endpoint
}

type FilesAddParams struct {
	File string `required:"true" comment:"файл в формате md"`
}

type FilesAddResponse struct {
	ID types.ID `comment:"{super}, загруженного файла"`
	QR types.BackURL
}

func (a *FilesAdd) InitEndpointScheme(s *draft.Scheme) {
	s.Project("markup2")

	s.Access(draft.Access.All)

	s.Method(draft.Method.POST)

	s.Name("«Загрузка файла»")

	s.URL("/api/v1/files/add")

	s.Params(FilesAddParams{
		File: "index.md",
	})

	// 200 OK
	s.Case(godraft.HTTPStatus(http.StatusOK), "Файл существует", func() {
		id := types.GenID()
		s.Body(FilesAddResponse{
			ID: id,
			QR: types.BackURL(fmt.Sprintf("https://markup2.ru/api/v1/files/get/%d", id)),
		})
	})
}
