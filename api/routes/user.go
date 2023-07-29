package route

import (
	"os"
	"shopbee/component/appctx"
	"shopbee/middleware"
	usertransport "shopbee/module/user/transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouterInit(router *gin.Engine, db *gorm.DB) {
	SECRETKEY := os.Getenv("SECRET_KEY")

	appCtx := appctx.NewAppContext(db, SECRETKEY)

	userApi := router.Group("api/v1/user")
	{
		userApi.POST("/register", usertransport.Register(appCtx))
		userApi.POST("/authenticate", usertransport.Login(appCtx))
		userApi.POST("/forgotpw", usertransport.ForgotPassword(appCtx))

		userApi.GET("/profile", middleware.RequireAuth(appCtx), usertransport.Profile(appCtx))
		userApi.DELETE("/delete", middleware.RequireAuth(appCtx), usertransport.DeleteUser(appCtx))
		userApi.DELETE("/delete/:id", middleware.RequireAuth(appCtx), usertransport.DeleteUserWithId(appCtx))
		userApi.GET("/list", middleware.RequireAuth(appCtx), usertransport.ListUser(appCtx))
		userApi.PATCH("/update/:id", middleware.RequireAuth(appCtx), usertransport.UpdateUser(appCtx))
	}
}
