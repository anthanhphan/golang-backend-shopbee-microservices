package route

import (
	"shopbee/component/appctx"
	categorytransport "shopbee/module/category/transport"

	"github.com/gin-gonic/gin"
)

func CategoryRouterInit(router *gin.Engine, appCtx appctx.AppContext) {

	productApi := router.Group("api/v1/category")
	{
		productApi.GET("/list", categorytransport.ListCategory(appCtx))
	}
}
