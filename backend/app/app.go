package app

import (
	"backend/app/controller/pong"
	"backend/app/middleware"
)

func AddRoutes() {
	api := e.Group("api")
	api.Use(middleware.Auth)
	api.GET("/ping", pong.Ping)
}
