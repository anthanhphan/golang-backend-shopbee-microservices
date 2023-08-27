package ordertransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	orderbiz "shopbee/module/order/biz"
	orderstorage "shopbee/module/order/storage"

	"github.com/gin-gonic/gin"
)

func ChangeOrderStatus(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		orderId, err := common.FromBase58(c.Param("id"))
		status := c.Param("status")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := orderstorage.NewSQLStore(db)
		biz := orderbiz.NewChangeOrderStatusBiz(store, requester)

		if err := biz.ChangeOrderStatus(
			c.Request.Context(),
			int(orderId.GetLocalID()),
			status); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}
