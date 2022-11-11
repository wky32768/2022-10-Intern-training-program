package app

import (
	"backend/app/controller/pong"
	"backend/app/controller/print/query"
	"backend/app/middleware"
)

func AddRoutes() {
	//api := e.Group("api")
	//api.Use(middleware.Auth)
	//api.GET("/ping", pong.Ping)

	e.Use(middleware.Auth)
	e.GET("/ping", pong.Ping)
	e.POST("/print/query", query.Query)
}
