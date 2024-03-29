package add

import (
	"bufio"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"markup2/markupapi/api/http/v1/response"
	"markup2/markupapi/core/interactors/files"
	"markup2/pkg/shortener"
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
	Title string `form:"title"`
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
	formFile, err := c.FormFile("file")
	if err != nil {
		log.Warnf("bad request: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	errs := echo.Map{}
	if req.Title == "" {
		errs["title"] = response.StatusEmpty
	}
	if formFile == nil {
		errs["file"] = response.StatusEmpty
	}

	if len(errs) != 0 {
		log.Warnf("failed to upload file: %v", errs)
		resp := response.Response{Errors: errs}

		return c.JSON(http.StatusOK, resp)
	}

	file, err := formFile.Open()
	if err != nil {
		log.Errorf("failed to open file: %v", err)
		resp := response.Response{Errors: echo.Map{
			"default": err,
		}}

		return c.JSON(http.StatusOK, resp)
	}

	id, err := h.files.Add(c.Request().Context(), ownerID, req.Title, bufio.NewReader(file))
	if err != nil {
		log.Warnf("failed to add file info: %v", err)

		resp := response.Response{Errors: echo.Map{
			"default": "failed to add file info",
		}}

		return c.JSON(http.StatusOK, resp)
	}

	encoded, _ := shortener.Encode([]byte(id))
	shortID := string(encoded)
	resp := response.Response{Data: echo.Map{
		"id":  shortID,
		"url": fmt.Sprintf("https://%s/pages/%s", h.cfg.RedirectHost, shortID),
	}}
	return c.JSON(http.StatusOK, resp)
}
