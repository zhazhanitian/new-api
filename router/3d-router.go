package router

import (
	"github.com/QuantumNous/new-api/controller"
	"github.com/QuantumNous/new-api/middleware"

	"github.com/gin-gonic/gin"
)

func Set3DRouter(router *gin.Engine) {
	g := router.Group("/v1")
	g.Use(middleware.RouteTag("relay"))
	g.Use(middleware.TokenAuth(), middleware.Distribute())
	{
		g.POST("/3d/generations", controller.RelayTask)
		g.GET("/3d/generations/:task_id", controller.RelayTaskFetch)
	}
}
