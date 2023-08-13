package producttransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	productbiz "shopbee/module/product/biz"
	productstorage "shopbee/module/product/storage"

	"github.com/gin-gonic/gin"
)

func RemoveProduct(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := productstorage.NewSQLStore(db)
		biz := productbiz.NewRemoveProductBiz(store, requester)

		if err := biz.RemoveProduct(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}
