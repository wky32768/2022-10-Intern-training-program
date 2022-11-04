package pong

import (
	"backend/app/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 检验是否活动, 输出pong
func Ping(c echo.Context) error {
	// just a demo
	return response.SendResponse(c, http.StatusOK, "", "pong!")
}
