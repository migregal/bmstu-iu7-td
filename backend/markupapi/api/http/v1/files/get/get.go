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
	Style  string `json:"style" query:"style"`
}

type File struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Length int64  `json:"length"`
}

func (h *Handler) Handle(c echo.Context) error {
	req := new(Request)
	if err := c.Bind(req); err != nil {
		log.Warnf("bad request: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	errs := echo.Map{}
	if req.ID == "" && c.Get("user_id") == nil {
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

	if req.ID != "" {
		data, err := h.files.Get(
			c.Request().Context(),
			req.ID,
			files.GetOpts{Format: req.Format, Style: req.Style},
		)
		if err != nil {
			log.Warnf("failed to get file info: %v", err)

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

	ownerID, ok := c.Get("user_id").(uint64)
	if !ok {
		resp := response.Response{Errors: echo.Map{
			"default": "invalid user_id",
		}}

		return c.JSON(http.StatusOK, resp)
	}

	filesInfo, err := h.files.Find(c.Request().Context(), ownerID)
	if err != nil {
		log.Warnf("failed to get files info: %v", err)

		resp := response.Response{Errors: echo.Map{
			"default": "failed to get files info",
		}}

		return c.JSON(http.StatusOK, resp)
	}

	files := make([]File, 0, len(filesInfo))
	for _, info := range filesInfo {
		files = append(files, File{
			ID:     info.ID,
			Title:  info.Title,
			Length: info.Length,
		})
	}

	return c.JSON(http.StatusOK, response.Response{Data: echo.Map{"files": files}})
}
