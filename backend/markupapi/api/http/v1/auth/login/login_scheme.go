package login

import (
	"net/http"

	"github.com/gothing/draft"
	"github.com/gothing/draft/types"

	"markup2/pkg/godraft"
)

type AuthLogin struct {
	draft.Endpoint
}

type AuthLoginParams struct {
	Login    types.Email    `required:"true"`
	Password types.Password `required:"true"`
	Remember bool           `comment:"Запомнить сессию"` // инлайн комментарий
}

type AuthLoginResponse struct {
	UserID types.UserID `comment:"{super}, авторизованного пользователя"`
	Token  types.Token
}

func (a *AuthLogin) InitEndpointScheme(s *draft.Scheme) {
	s.Project("markup2")

	s.Access(draft.Access.All)

	s.Method(draft.Method.POST)

	s.Name("«Вход»")

	s.URL("/api/v1/auth/login")

	s.Params(AuthLoginParams{
		Login:    types.GenEmail(),
		Password: types.GenPassword(),
	})

	// 200 OK
	s.Case(godraft.HTTPStatus(http.StatusOK), "Успешная авторизация", func() {
		s.Body(AuthLoginResponse{
			UserID: types.GenUserID(),
			Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
		})
	})

	// 403 OK
	s.Case(godraft.HTTPStatus(http.StatusForbidden), "Неправильный Логин или Пароль", func() {
		s.Params(AuthLoginParams{
			Login:    "not-exists-login",
			Password: types.GenPassword(),
		})
	})
}
