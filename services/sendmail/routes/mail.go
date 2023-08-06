package routes

import (
	"shopbee/component/appctx"
	mailtransport "shopbee/module/sendmail/transport"

	"github.com/gin-gonic/gin"
)

func MailServiceRouterInit(router *gin.Engine, appCtx appctx.AppContext) {

	mailApi := router.Group("api/v1/mail")
	{
		mailApi.POST("/send", mailtransport.SendMail(appCtx))
	}
}
