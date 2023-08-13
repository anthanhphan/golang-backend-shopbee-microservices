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

func UpdateCart(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data cartmodel.CartUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := cartstorage.NewSQLStore(db)
		biz := cartbiz.NewUpdateCartBiz(store, requester)

		data.UserId = requester.GetUserId()

		productId, err := common.FromBase58(data.ProductUID)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data.ProductId = int(productId.GetLocalID())

		if err := biz.UpdateCart(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}
