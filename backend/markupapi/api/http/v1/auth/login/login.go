package login

import (
	"net/http"
	"net/mail"

	"github.com/labstack/echo/v4"

	"markup2/markupapi/api/http/v1/response"
	"markup2/markupapi/core/interactors/user"
	"markup2/pkg/jwt"
)

type Handler struct {
	user user.Interactor
}

func New(user user.Interactor) Handler {
	return Handler{user: user}
}

type Request struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) Handle(c echo.Context) error {
	req := new(Request)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	errs := echo.Map{}
	if req.Login == "" {
		errs["login"] = response.StatusEmpty
	}
	if _, err := mail.ParseAddress(req.Login); err != nil {
		errs["login"] = response.StatusInvalid
	}
	if req.Password == "" {
		errs["password"] = response.StatusEmpty
	}

	if len(errs) != 0 {
		resp := response.Response{Errors: errs}

		return c.JSON(http.StatusOK, resp)
	}

	user, err := h.user.Get(req.Login)
	if err != nil {
		resp := response.Response{Errors: echo.Map{
			"default": "failed to get user info",
		}}

		return c.JSON(http.StatusOK, resp)
	}

	if req.Password != user.PasswordHash {
		resp := response.Response{Errors: echo.Map{
			"password": response.StatusIncorrect,
		}}

		return c.JSON(http.StatusOK, resp)
	}

	t, err := jwt.NewToken([]byte("secret"), req.Login, user.ID)
	if err != nil {
		resp := response.Response{Errors: echo.Map{
			"default": "failed to create token",
		}}

		return c.JSON(http.StatusOK, resp)
	}

	resp := response.Response{Data: echo.Map{
		"id":    user.ID,
		"token": t,
	}}

	return c.JSON(http.StatusOK, resp)
}
