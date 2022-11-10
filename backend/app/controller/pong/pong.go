package pong

import (
	"backend/app/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Ping 检验是否活动, 输出pong
func Ping(c echo.Context) error {
	// just a demo
	return response.SendResponse(c, http.StatusOK, "", "pong!")
}
