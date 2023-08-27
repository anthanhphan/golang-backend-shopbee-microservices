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
		orderApi.GET("/list", middleware.RequireAuth(appCtx), ordertransport.ViewOrder(appCtx))
		orderApi.GET("/:id", middleware.RequireAuth(appCtx), ordertransport.ViewOrderDetail(appCtx))

		orderApi.GET("/shop/list", middleware.RequireAuth(appCtx), ordertransport.ViewShopOrder(appCtx))
		orderApi.GET("", middleware.RequireAuth(appCtx), ordertransport.ViewShopOrderDetail(appCtx))
		orderApi.PATCH("/shop/:id/:status", middleware.RequireAuth(appCtx), ordertransport.ChangeOrderStatus(appCtx))
	}
}
