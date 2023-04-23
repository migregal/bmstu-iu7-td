package logout

import (
	"markup2/markupapi/core/interactors/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func New(_ user.Interactor) Handler {
	return Handler{}
}

func (h *Handler) Handle(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
