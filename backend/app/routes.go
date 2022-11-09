package app

import (
	"backend/app/controller"
	"backend/app/middleware"
)

func addRoutes() {
	api := e.Group("api")
	api.Use(middleware.Auth)
	api.GET("/ping", controller.Ping)
}
