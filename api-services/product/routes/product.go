package route

import (
	"shopbee/component/appctx"
	"shopbee/middleware"
	producttransport "shopbee/module/product/transport"

	"github.com/gin-gonic/gin"
)

func ProductRouterInit(router *gin.Engine, appCtx appctx.AppContext) {

	productApi := router.Group("api/v1/product")
	{
		productApi.POST("/create", middleware.RequireAuth(appCtx), producttransport.CreateProduct(appCtx))
		productApi.GET("/list", producttransport.ListProduct(appCtx))
		productApi.GET("/view/:id", producttransport.ViewProduct(appCtx))
	}
}
