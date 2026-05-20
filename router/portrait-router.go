package router

import (
	"github.com/QuantumNous/new-api/controller"
	"github.com/QuantumNous/new-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetPortraitRouter(router *gin.Engine) {
	portraitRouter := router.Group("/v1/portrait")
	portraitRouter.Use(middleware.TokenAuth())
	{
		portraitRouter.POST("/groups", controller.PortraitCreateGroup)
		portraitRouter.GET("/groups", controller.PortraitListGroups)

		portraitRouter.POST("/assets", controller.PortraitCreateAsset)
		portraitRouter.GET("/assets", controller.PortraitListAssets)
		portraitRouter.GET("/assets/:assetId", controller.PortraitGetAsset)
	}
}
