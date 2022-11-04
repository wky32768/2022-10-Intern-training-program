package response

import (
	"github.com/labstack/echo/v4"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func response(c echo.Context, code int, message string, data interface{}) {
	return c.json(
		http.StatusOK, 
		Result {
			code, message, data,
		}
	)
}
