package carttransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	cartbiz "shopbee/module/cart/biz"
	cartstorage "shopbee/module/cart/storage"

	"github.com/gin-gonic/gin"
)

func RemoveProductFromCart(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		productId, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := cartstorage.NewSQLStore(db)
		biz := cartbiz.NewRemoveProductBiz(store, requester)

		if err := biz.RemoveProduct(
			c.Request.Context(),
			int(productId.GetLocalID()),
			requester.GetUserId(),
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse("remove product from cart success"))
	}
}
