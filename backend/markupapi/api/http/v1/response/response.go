package response

import "github.com/labstack/echo/v4"

var (
	StatusEmpty   = "empty"
	StatusInvalid = "invalid"
)

type Response struct {
	Data   echo.Map `json:"data,omitempty"`
	Errors echo.Map `json:"errors,omitempty"`
}
