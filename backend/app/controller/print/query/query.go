package query

import (
	"backend/app/response"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// 读入query并原样输出
func query(c echo.Context) error {
	input := c.QueryParam("name")
	fmt.Println(input)
	return response.SendResponse(c, http.StatusOK, "", "query done!")
}
