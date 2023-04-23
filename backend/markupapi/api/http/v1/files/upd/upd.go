package upd

import (
	"bufio"
	"errors"
	"io"
	"markup2/markupapi/api/http/v1/response"
	"markup2/markupapi/core/interactors/files"
	"markup2/pkg/shortener"
	"markup2/pkg/validation"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler struct {
	files files.Interactor
}

func New(files files.Interactor) Handler {
	return Handler{files: files}
}

type Request struct {
	ID    string `param:"id"`
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
	if err != nil && !errors.Is(err, http.ErrMissingFile){
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

	if req.Title == "" && formFile == nil {
		errs["title"] = response.StatusEmpty
		errs["file"] = response.StatusEmpty
	}

	if len(errs) != 0 {
		log.Warnf("failed to upload file: %v", errs)
		resp := response.Response{Errors: errs}

		return c.JSON(http.StatusOK, resp)
	}

	var reader io.Reader
	if formFile != nil {
		file, err := formFile.Open()
		if err != nil {
			log.Errorf("failed to open file: %v", err)
			resp := response.Response{Errors: echo.Map{
				"default": err,
			}}

			return c.JSON(http.StatusOK, resp)
		}

		reader = bufio.NewReader(file)
	}

	id, err := h.files.Update(c.Request().Context(), ownerID, req.Title, string(fullID), reader)
	if err != nil {
		log.Warnf("failed to update file info: %v", err)

		resp := response.Response{Errors: echo.Map{
			"default": "failed to update file info",
		}}

		return c.JSON(http.StatusOK, resp)
	}

	shortID, _ := shortener.Encode([]byte(id))
	resp := response.Response{Data: echo.Map{"id": string(shortID)}}
	return c.JSON(http.StatusOK, resp)
}
