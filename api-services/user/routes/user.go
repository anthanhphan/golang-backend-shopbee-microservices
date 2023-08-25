package route

import (
	"shopbee/component/appctx"
	"shopbee/middleware"
	reqtransport "shopbee/module/request/transport"
	usertransport "shopbee/module/user/transport"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(router *gin.Engine, appCtx appctx.AppContext) {

	userApi := router.Group("api/v1/user")
	{
		userApi.GET("/profile", middleware.RequireAuth(appCtx), usertransport.Profile(appCtx))
		userApi.GET("/list", middleware.RequireAuth(appCtx), usertransport.ListUser(appCtx))
		userApi.GET("/getstatus/:id", usertransport.FindUser(appCtx))
		userApi.GET("/like/:id", middleware.RequireAuth(appCtx), usertransport.IsLikedShop(appCtx))

		userApi.POST("/like/:id", middleware.RequireAuth(appCtx), usertransport.LikeShop(appCtx))
		userApi.POST("/dislike/:id", middleware.RequireAuth(appCtx), usertransport.LikeShop(appCtx))

		userApi.POST("/register", usertransport.Register(appCtx))
		userApi.POST("/authenticate", usertransport.Login(appCtx))
		userApi.POST("/forgotpw", usertransport.ForgotPassword(appCtx))
		userApi.POST("/changepw", middleware.RequireAuth(appCtx), usertransport.ChangePassword(appCtx))

		userApi.POST("/upgrade", middleware.RequireAuth(appCtx), reqtransport.CreateRequestUpgrade(appCtx))
		userApi.POST("/report", middleware.RequireAuth(appCtx), reqtransport.CreateRequestBanUser(appCtx))

		userApi.DELETE("/delete", middleware.RequireAuth(appCtx), usertransport.DeleteUser(appCtx))
		userApi.DELETE("/delete/:id", middleware.RequireAuth(appCtx), usertransport.DeleteUserWithId(appCtx))

		userApi.PATCH("/update/:id", middleware.RequireAuth(appCtx), usertransport.UpdateUser(appCtx))
	}
}
