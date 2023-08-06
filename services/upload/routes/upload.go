package routes

import (
	"shopbee/component/appctx"
	uploadtransport "shopbee/module/upload/transport"

	"github.com/gin-gonic/gin"
)

func UploadRouterInit(router *gin.Engine, appCtx appctx.AppContext) {

	uploadApi := router.Group("api/v1/upload")
	{
		uploadApi.POST("/image", uploadtransport.Upload(appCtx))
	}
}
