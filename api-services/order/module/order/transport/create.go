package ordertransport

import (
	"fmt"
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	orderbiz "shopbee/module/order/biz"
	ordermodel "shopbee/module/order/model"
	orderstorage "shopbee/module/order/storage"

	"github.com/gin-gonic/gin"
)

func CreateOrder(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var recieve ordermodel.OrderCreate

		if err := c.ShouldBind(&recieve); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		shopId, err := common.FromBase58(recieve.ShopId)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paymentId, err := common.FromBase58(recieve.PaymentId)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data ordermodel.Order
		data.UserId = requester.GetUserId()
		data.ShopId = int(shopId.GetLocalID())
		data.PaymentId = int(paymentId.GetLocalID())
		data.ShippingAddr = recieve.ShippingAddr
		data.TotalPrice = recieve.TotalPrice

		store := orderstorage.NewSQLStore(db)
		biz := orderbiz.NewCreateOrderBiz(store, requester)

		fmt.Print(recieve.ProductList)
		if err := biz.CreateOrder(c.Request.Context(), &data, recieve.ProductList); err != nil {
			panic(err)
		}

		data.GenUID(1)
		c.JSON(http.StatusOK, common.SimpleSucessResponse(data))
	}
}
