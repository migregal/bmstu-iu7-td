package auth

import (
	"github.com/gothing/draft"

	"markup2/markupapi/api/http/v1/auth/login"
	"markup2/markupapi/api/http/v1/auth/logout"
	"markup2/markupapi/api/http/v1/auth/registration"
)

var Service = draft.Compose(
	"Авторизация",
	login.Endpoint,
	logout.Endpoint,
	registration.Endpoint,
)
