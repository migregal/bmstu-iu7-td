package registration

import (
	"errors"
	"markup2/markupapi/api/http/v1/response"
	"markup2/markupapi/core/interactors"
	"markup2/markupapi/core/interactors/user"
	"net/http"
	"net/mail"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
		log.Warnf("bad request: %v", err)
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
		log.Warnf("bad request: %v", errs)
		resp := response.Response{Errors: errs}

		return c.JSON(http.StatusOK, resp)
	}

	user := user.UserInfo{Login: req.Login, Password: req.Password}
	if _, err := h.user.Register(user) ; err != nil {
		log.Errorf("failed to register new user: %v", err)

		desc := "failed to register new user"
		if errors.Is(err, interactors.ErrExists) {
			desc = "user already exists"
		}

		resp := response.Response{Errors: echo.Map{
			"default": desc,
		}}

		return c.JSON(http.StatusOK, resp)
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/api/v1/auth/login")
}
