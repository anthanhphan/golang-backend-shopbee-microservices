package route

import (
	"shopbee/component/appctx"
	"shopbee/middleware"
	wishlisttransport "shopbee/module/wishlist/transport"

	"github.com/gin-gonic/gin"
)

func WishListRouterInit(router *gin.Engine, appCtx appctx.AppContext) {

	wishListApi := router.Group("api/v1/wishlist")
	{
		wishListApi.POST("/addproduct/:id", middleware.RequireAuth(appCtx), wishlisttransport.AddProductToWishList(appCtx))
		wishListApi.POST("/removeproduct/:id", middleware.RequireAuth(appCtx), wishlisttransport.RemoveProductFromWishList(appCtx))
		wishListApi.GET("/view", middleware.RequireAuth(appCtx), wishlisttransport.ViewMyWishList(appCtx))
	}
}
