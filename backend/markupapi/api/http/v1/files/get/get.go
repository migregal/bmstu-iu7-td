package get

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"markup2/markupapi/api/http/v1/response"
	"markup2/markupapi/core/interactors"
	"markup2/markupapi/core/interactors/files"
	"markup2/pkg/shortener"
	"markup2/pkg/validation"
)

type Config struct {
	RedirectHost string
}

type Handler struct {
	cfg   Config
	files files.Interactor
}

func New(cfg Config, files files.Interactor) Handler {
	return Handler{cfg: cfg, files: files}
}

type Request struct {
	ID     string `param:"id"`
	Format string `query:"format"`
	Style  string `query:"style"`
}

type File struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Length int64  `json:"length"`
	URL    string `json:"url"`
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
	var fullID string
	if req.ID != "" {
		decoded, err := shortener.Decode([]byte(req.ID))
		if err != nil || !validation.IsHex(string(decoded)) {
			return c.Redirect(http.StatusTemporaryRedirect, "/404")
		}

		fullID = string(decoded)
	}

	if len(errs) != 0 {
		log.Warnf("failed to get file: %v", errs)
		resp := response.Response{Errors: errs}

		return c.JSON(http.StatusOK, resp)
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

	if fullID != "" {
		data, err := h.files.Get(
			c.Request().Context(),
			fullID,
			files.GetOpts{Format: req.Format, Style: req.Style},
		)
		if err != nil {
			log.Warnf("failed to get file info: %v", err)

			desc := "failed to get file info"
			if errors.Is(err, interactors.ErrNotFound) {
				return c.Redirect(http.StatusTemporaryRedirect, "/404")
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
		encoded, _ := shortener.Encode([]byte(info.ID))
		shortID := string(encoded)

		files = append(files, File{
			ID:     string(shortID),
			Title:  info.Title,
			Length: info.Length,
			URL:    fmt.Sprintf("https://%s/pages/%s", h.cfg.RedirectHost, shortID),
		})
	}

	return c.JSON(http.StatusOK, response.Response{Data: echo.Map{"files": files}})
}
