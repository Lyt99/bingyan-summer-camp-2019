package context

import (
	"encoding/base64"
	"net/http"

	"github.com/labstack/echo"
)

// Response
type Response struct {
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
	Success bool        `json:"success"`
	ErrHint string      `json:"errHint,omitempty"`
}

// Success
func Success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Data:    data,
		Error:   "",
		Success: true,
	})
}

// Error
func Error(c echo.Context, status int, data string, err error) error {
	ret := Response{
		Data:    nil,
		Error:   data,
		Success: false,
	}

	if err != nil {
		ret.ErrHint = base64.StdEncoding.EncodeToString([]byte(err.Error()))
	}

	return c.JSON(status, ret)
}
