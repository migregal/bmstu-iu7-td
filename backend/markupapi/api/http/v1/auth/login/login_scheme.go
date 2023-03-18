package login

import (
	"github.com/gothing/draft"
	"github.com/gothing/draft/types"
)

type AuthLogin struct {
	draft.Endpoint
}

type AuthLoginParams struct {
	Login    types.Login    `required:"true"`
	Password types.Password `required:"true"`
	Remember bool           `comment:"Запомнить сессию"` // инлайн комментарий
}

type AuthLoginResponse struct {
	UserID types.UserID `comment:"{super}, авторизованного пользователя"`
}

func (a *AuthLogin) InitEndpointScheme(s *draft.Scheme) {
	s.Project("auth")

	s.Access(draft.Access.All)

	s.Method(draft.Method.POST)

	s.Name("«Вход»")

	s.URL("/api/v1/auth/login")

	s.Params(AuthLoginParams{
		Login:    types.GenLogin(),
		Password: types.GenPassword(),
	})

	// 200 OK
	s.Case(draft.Status.OK, "Успешная авторизация", func() {
		s.Body(AuthLoginResponse{
			UserID: types.GenUserID(),
		})
	})

	// 403 OK
	s.Case(draft.Status.Denied, "Неправильный Логин или Пароль", func() {
		// Переопределяем базовые параметры запроса
		s.Params(AuthLoginParams{
			Login:    "not-exists-login",
			Password: types.GenPassword(),
		})
	})
}
