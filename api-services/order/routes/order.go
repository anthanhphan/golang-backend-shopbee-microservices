package route

import (
	"shopbee/component/appctx"
	"shopbee/middleware"
	ordertransport "shopbee/module/order/transport"

	"github.com/gin-gonic/gin"
)

func OrderRouterInit(router *gin.Engine, appCtx appctx.AppContext) {

	orderApi := router.Group("api/v1/order")
	{
		orderApi.POST("/create", middleware.RequireAuth(appCtx), ordertransport.CreateOrder(appCtx))
	}
}
