package get

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"markup2/markupapi/api/http/v1/response"
)

type Handler struct {
}

func New() Handler {
	return Handler{}
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

	contentType := "html"
	switch req.Format {
	case "":
		req.Format = "html"
	case "html":
		contentType = "text/html"
	case "md", "plain":
		contentType = "text/plain"
	default:
		errs["format"] = response.StatusEmpty
	}

	if len(errs) != 0 {
		log.Warnf("failed to get file: %v", errs)
		resp := response.Response{Errors: errs}

		return c.JSON(http.StatusOK, resp)
	}

	data := []byte(`# Hello, world`)

	return c.Blob(http.StatusOK, contentType, mdToHTML(data))
}
