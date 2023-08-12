package route

import (
	"shopbee/component/appctx"
	"shopbee/middleware"
	reqtransport "shopbee/module/request/transport"

	"github.com/gin-gonic/gin"
)

func RequestRouterInit(router *gin.Engine, appCtx appctx.AppContext) {

	reqApi := router.Group("api/v1/admin")
	{
		reqApi.POST("upgrade/accept/:id", middleware.RequireAuth(appCtx), reqtransport.AcceptRequestUpgrade(appCtx))
		reqApi.POST("upgrade/deny/:id", middleware.RequireAuth(appCtx), reqtransport.DenyRequestUpgrade(appCtx))

		reqApi.GET("upgrade/list", middleware.RequireAuth(appCtx), reqtransport.ListRequetUpgrade(appCtx))
		reqApi.GET("report/list", middleware.RequireAuth(appCtx), reqtransport.ListReport(appCtx))
	}
}
