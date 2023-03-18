package registration

import (
	"net/http"

	"github.com/gothing/draft"
	"github.com/gothing/draft/types"

	"markup2/pkg/godraft"
)

type AuthRegistration struct {
	draft.Endpoint
}

type AuthRegistrationParams struct {
	Login    types.Email    `required:"true"`
	Password types.Password `required:"true"`
}

type AuthRegistrationResponse struct {
	UserID types.UserID
	Token  types.Token
}

func (a *AuthRegistration) InitEndpointScheme(s *draft.Scheme) {
	s.Project("markup2")

	s.Access(draft.Access.Auth)

	s.Method(draft.Method.POST)

	s.Name("«Регистрация»")

	s.URL("/api/v1/auth/registration")

	// 200 OK
	s.Case(godraft.HTTPStatus(http.StatusOK), "Успех", func() {
		s.Body(AuthRegistrationResponse{
			UserID: types.GenUserID(),
			Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
		})
	})

	s.Case(godraft.HTTPStatus(http.StatusBadRequest), "Некорректные логин или пароль", func() {
		s.Params(AuthRegistrationParams{
			Login:    types.GenEmail(),
			Password: "",
		})
		s.Body(godraft.InvalidCaseBody("password", "empty", ""))
	})

	s.Case(godraft.HTTPStatus(http.StatusConflict), "Существующий логин", func() {
		login := types.GenEmail()

		s.Params(AuthRegistrationParams{
			Login:    login,
			Password: types.GenPassword(),
		})

		s.Body(godraft.InvalidCaseBody("login", "used", login))
	})
}
