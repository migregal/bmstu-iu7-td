package del

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"markup2/markupapi/api/http/v1/response"
	"markup2/markupapi/core/interactors"
	"markup2/markupapi/core/interactors/files"
	"markup2/pkg/shortener"
	"markup2/pkg/validation"
)

type Handler struct {
	files files.Interactor
}

func New(files files.Interactor) Handler {
	return Handler{files: files}
}

type Request struct {
	ID string `param:"id"`
}

func (h *Handler) Handle(c echo.Context) error {
	ownerID, ok := c.Get("user_id").(uint64)
	if !ok {
		resp := response.Response{Errors: echo.Map{
			"default": "invalid user_id",
		}}

		return c.JSON(http.StatusOK, resp)
	}

	req := new(Request)
	if err := c.Bind(req); err != nil {
		log.Warnf("bad request: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	errs := echo.Map{}
	if req.ID == "" {
		errs["id"] = response.StatusEmpty
	}
	fullID, err := shortener.Decode([]byte(req.ID))
	if err != nil || !validation.IsHex(string(fullID)) {
		errs["id"] = response.StatusInvalid
	}

	if len(errs) != 0 {
		log.Warnf("failed to delete file: %v", errs)
		resp := response.Response{Errors: errs}

		return c.JSON(http.StatusOK, resp)
	}

	err = h.files.Delete(c.Request().Context(), ownerID, string(fullID))
	if err != nil {
		log.Warnf("failed to delete file info: %v", err)

		desc := "failed to delete file info"
		if errors.Is(err, interactors.ErrNotFound) {
			desc = "file not found"
		}

		resp := response.Response{Errors: echo.Map{
			"default": desc,
		}}

		return c.JSON(http.StatusOK, resp)
	}

	return c.JSON(http.StatusOK, response.Response{Data: echo.Map{"id": req.ID}})
}
