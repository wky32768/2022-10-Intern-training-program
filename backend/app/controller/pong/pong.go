package pong

import (
	"backend/app/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	// just a demo
	return response.SendResponse(c, http.StatusOK, "", "pong!")
}
