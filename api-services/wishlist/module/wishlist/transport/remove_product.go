package wishlisttransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	wishlistbiz "shopbee/module/wishlist/biz"
	wishliststorage "shopbee/module/wishlist/storage"

	"github.com/gin-gonic/gin"
)

func RemoveProductFromWishList(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		productId, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := wishliststorage.NewSQLStore(db)
		biz := wishlistbiz.NewRemoveProductBiz(store, requester)

		if err := biz.RemoveProduct(c.Request.Context(), int(productId.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse("remove product from wishlist success"))
	}
}
