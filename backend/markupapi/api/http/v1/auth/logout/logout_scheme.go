package logout

import (
	"github.com/gothing/draft"
	"github.com/gothing/draft/types"
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

	s.Method(draft.Method.DELETE)

	s.Name("«Выход»")

	s.URL("/api/v1/auth/logout")

	// 200 OK
	s.Case(draft.Status.OK, "Успех", func() {
		s.Body(AuthLogoutResponse{
			UserID: types.GenUserID(),
		})
	})
}
