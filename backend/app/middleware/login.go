package middleware

import (
	"github.com/labstack/echo/v4"
)

// 设置为默认用户
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("User", "default")
		return next(c)
	}
}
