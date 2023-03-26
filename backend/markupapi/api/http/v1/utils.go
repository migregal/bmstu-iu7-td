package v1

import (
	"fmt"
	"markup2/markupapi/api/http/v1/response"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func IgnoreError(_ echo.Context, err error) error {
	log.Warnf("unauthorized: %v", err)
	return nil
}

func ForceAuthError(c echo.Context, err error) error {
	log.Errorf("unauthorized: %v", err)

	err = c.JSON(http.StatusOK, response.Response{Errors: echo.Map{
		"default": "unauthorized",
	}})
	if err != nil {
		log.Errorf("failed to send rresponse to user: %v", err)
	}

	return fmt.Errorf("unauthorized")
}
