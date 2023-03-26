package get

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"markup2/markupapi/api/http/v1/response"
	"markup2/markupapi/core/interactors"
	"markup2/markupapi/core/interactors/files"
)

type Handler struct {
	files files.Interactor
}

func New(files files.Interactor) Handler {
	return Handler{files: files}
}

type Request struct {
	ID     string `json:"id" query:"id"`
	Format string `json:"format" query:"format"`
}

func (h *Handler) Handle(c echo.Context) error {
	req := new(Request)
	if err := c.Bind(req); err != nil {
		log.Warnf("bad request: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	errs := echo.Map{}
	if req.ID == "" {
		errs["id"] = response.StatusEmpty
	}

	contentType := "text/html"
	switch req.Format {
	case "html":
		contentType = "text/html"
	case "md":
		contentType = "text/markdown"
	case "plain":
		contentType = "text/plain"
	default:
		req.Format = "html"
	}

	if len(errs) != 0 {
		log.Warnf("failed to get file: %v", errs)
		resp := response.Response{Errors: errs}

		return c.JSON(http.StatusOK, resp)
	}

	data, err := h.files.Get(files.Opts{Format: req.Format})
	if err != nil {
		log.Warnf("failed to get file info: %v", errs)

		desc := "failed to get file info"
		if errors.Is(err, interactors.ErrNotFound) {
			desc = "user doesn't exist"
		}
		resp := response.Response{Errors: echo.Map{
			"default": desc,
		}}

		return c.JSON(http.StatusOK, resp)
	}

	return c.Blob(http.StatusOK, contentType, data)
}
