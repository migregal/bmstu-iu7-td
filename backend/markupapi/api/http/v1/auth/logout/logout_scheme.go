package logout

import (
	"net/http"

	"github.com/gothing/draft"
	"github.com/gothing/draft/types"

	"markup2/pkg/godraft"
)

type AuthLogout struct {
	draft.Endpoint
}

type AuthLogoutParams struct {
}

type AuthLogoutResponse struct {
	UserID types.UserID
}

func (a *AuthLogout) InitEndpointScheme(s *draft.Scheme) {
	s.Project("auth")

	s.Access(draft.Access.Auth)

	s.Method(draft.Method.POST)

	s.Name("«Выход»")

	s.URL("/api/v1/auth/logout")

	// 200 OK
	s.Case(godraft.HTTPStatus(http.StatusOK), "Успех", func() {
		s.Body(AuthLogoutResponse{
			UserID: types.GenUserID(),
		})
	})
}
