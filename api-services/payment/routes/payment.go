package route

import (
	"shopbee/component/appctx"
	"shopbee/middleware"
	paymenttransport "shopbee/module/payment/transport"

	"github.com/gin-gonic/gin"
)

func PaymentRouterInit(router *gin.Engine, appCtx appctx.AppContext) {

	paymentApi := router.Group("api/v1/payment")
	{
		paymentApi.POST("/create", middleware.RequireAuth(appCtx), paymenttransport.CreatePayment(appCtx))
	}
}
