package route

import (
	"shopbee/component/appctx"
	"shopbee/middleware"
	carttransport "shopbee/module/cart/transport"

	"github.com/gin-gonic/gin"
)

func CartRouterInit(router *gin.Engine, appCtx appctx.AppContext) {

	cartApi := router.Group("api/v1/cart")
	{
		cartApi.POST("/addproduct/:id", middleware.RequireAuth(appCtx), carttransport.AddToCart(appCtx))
		cartApi.GET("/view", middleware.RequireAuth(appCtx), carttransport.ViewMyCart(appCtx))
	}
}
