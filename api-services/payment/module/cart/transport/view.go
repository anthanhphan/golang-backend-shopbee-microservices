package carttransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	cartbiz "shopbee/module/cart/biz"
	cartstorage "shopbee/module/cart/storage"

	"github.com/gin-gonic/gin"
)

func ViewMyCart(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := cartstorage.NewSQLStore(db)
		biz := cartbiz.NewViewMyCartBiz(store, requester)

		data, err := biz.ViewMyCart(c.Request.Context(), requester.GetUserId())

		if err != nil {
			panic(err)
		}

		for _, item := range data {
			uid := common.NewUID(uint32(item["shop_id"].(int)), 1, 1)
			fakeId := &uid
			item["shop_id"] = fakeId.String()

			encryptedProductIDs := make([]string, 0)
			for _, x := range item["product_ids"].([]int) {
				uid := common.NewUID(uint32(x), 1, 1)
				fakeId := &uid
				encryptedProductIDs = append(encryptedProductIDs, fakeId.String())
			}

			item["product_ids"] = encryptedProductIDs
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data))
	}
}
