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
		portraitRouter.PUT("/groups/:groupId", controller.PortraitUpdateGroup)

		portraitRouter.POST("/assets", controller.PortraitCreateAsset)
		portraitRouter.GET("/assets", controller.PortraitListAssets)
		portraitRouter.GET("/assets/:assetId", controller.PortraitGetAsset)
		portraitRouter.PUT("/assets/:assetId", controller.PortraitUpdateAsset)
	}

	// 管理端：把缺少渠道绑定的素材补建到所有参与渠道（幂等）
	portraitAdminRouter := router.Group("/api/portrait")
	portraitAdminRouter.Use(middleware.RootAuth())
	{
		portraitAdminRouter.POST("/backfill", controller.PortraitBackfill)
	}
}
