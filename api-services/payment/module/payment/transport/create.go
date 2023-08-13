package paymenttransport

import (
	"fmt"
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	paymentbiz "shopbee/module/payment/biz"
	paymentmodel "shopbee/module/payment/model"
	paymentstorage "shopbee/module/payment/storage"

	"github.com/gin-gonic/gin"
)

func CreatePayment(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data paymentmodel.Payment

		if err := c.ShouldBind(&data); err != nil {
			fmt.Print("erorr herr")
			panic(common.ErrInvalidRequest(err))
		}

		data.UserId = requester.GetUserId()

		store := paymentstorage.NewSQLStore(db)
		biz := paymentbiz.NewCreatePaymentBiz(store, requester)

		result, err := biz.CreatePayment(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		result.GenUID(1)
		c.JSON(http.StatusOK, common.SimpleSucessResponse(result))
	}
}
