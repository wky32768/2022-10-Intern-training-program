package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SendResponse(c echo.Context, code int, message string, data ...interface{}) error {
	return c.JSON(
		http.StatusOK,
		Result{
			code, message, data,
		},
	)
}
