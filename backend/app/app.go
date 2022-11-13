package app

import (
	"backend/app/controller/pong"
	"backend/app/controller/print/query"
	"backend/app/middleware"
	"backend/model"
)

func AddRoutes() {
	//api := e.Group("api")

	e.Use(middleware.Auth)
	e.GET("/ping", pong.Ping)
	e.POST("/print/query", query.Query)
	e.POST("/db/create", model.Add)
	e.POST("/db/search", model.Get)
	e.POST("/db/delete", model.Delete)
}
