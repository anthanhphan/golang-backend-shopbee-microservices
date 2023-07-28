package route

import (
	"os"
	"shopbee/component/appctx"
	"shopbee/middleware"
	producttransport "shopbee/module/product/transport"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRouterInit(router *gin.Engine, db *gorm.DB) {
	SECRETKEY := os.Getenv("SECRET_KEY")

	appCtx := appctx.NewAppContext(db, SECRETKEY)

	productApi := router.Group("api/v1/product", middleware.RequireAuth(appCtx))
	{
		productApi.POST("/create", producttransport.CreateProduct(appCtx))
	}
}
