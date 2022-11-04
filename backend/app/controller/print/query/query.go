package query

import (
	"backend/app/response"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 读入query并原样输出
func query(c echo.Context) error {
	input := c.queryParam("name")
	fmt.Println(input)
	return response.SendResponse(c, http.StatusOK, "", "query done!")
}
