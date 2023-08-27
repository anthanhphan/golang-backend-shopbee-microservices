package ordertransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	orderbiz "shopbee/module/order/biz"
	orderstorage "shopbee/module/order/storage"

	"github.com/gin-gonic/gin"
)

func ViewShopOrder(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := orderstorage.NewSQLStore(db)
		biz := orderbiz.NewViewShopOrderBiz(store, requester)

		result, err := biz.ViewShopOrder(c.Request.Context(), requester.GetUserId())

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].GenUID(common.DbTypeProduct)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(result))
	}
}

func ViewShopOrderDetail(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		orderId, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := orderstorage.NewSQLStore(db)
		biz := orderbiz.NewViewShopOrderBiz(store, requester)
		result, err := biz.ViewShopOrderDetail(c.Request.Context(), int(orderId.GetLocalID()))

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].GenUID(1)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(result))
	}
}
