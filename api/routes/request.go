package route

import (
	"shopbee/component/appctx"
	"shopbee/middleware"
	reqtransport "shopbee/module/request/transport"

	"github.com/gin-gonic/gin"
)

func RequestRouterInit(router *gin.Engine, appCtx appctx.AppContext) {

	reqApi := router.Group("api/v1")
	{
		reqApi.POST("user/upgrade", middleware.RequireAuth(appCtx), reqtransport.CreateRequestUpgrade(appCtx))
		reqApi.POST("report", middleware.RequireAuth(appCtx), reqtransport.CreateRequestBanUser(appCtx))
		reqApi.POST("upgrade/accept", middleware.RequireAuth(appCtx), reqtransport.AcceptRequestUpgrade(appCtx))
		reqApi.POST("upgrade/deny", middleware.RequireAuth(appCtx), reqtransport.DenyRequestUpgrade(appCtx))

		reqApi.GET("upgrade/list", middleware.RequireAuth(appCtx), reqtransport.ListRequetUpgrade(appCtx))
		reqApi.GET("report/list", middleware.RequireAuth(appCtx), reqtransport.ListReport(appCtx))
	}
}
