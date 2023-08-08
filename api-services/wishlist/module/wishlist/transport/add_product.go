package wishlisttransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	wishlistbiz "shopbee/module/wishlist/biz"
	wishliststorage "shopbee/module/wishlist/storage"

	"github.com/gin-gonic/gin"
)

func AddProductToWishList(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMainDBConnection()

		productId, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := wishliststorage.NewSQLStore(db)
		biz := wishlistbiz.NewAddToWishListBiz(store, requester)

		if err := biz.AddToWishList(c.Request.Context(), int(productId.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}
