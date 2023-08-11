package carttransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	cartbiz "shopbee/module/cart/biz"
	cartmodel "shopbee/module/cart/model"
	cartstorage "shopbee/module/cart/storage"

	"github.com/gin-gonic/gin"
)

func AddToCart(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data cartmodel.CartCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := cartstorage.NewSQLStore(db)
		biz := cartbiz.NewAddToCartBiz(store, requester)

		if err := biz.AddToCart(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}
