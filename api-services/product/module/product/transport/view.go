package producttransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	productbiz "shopbee/module/product/biz"
	productstorage "shopbee/module/product/storage"

	"github.com/gin-gonic/gin"
)

func ViewProduct(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := productstorage.NewSQLStore(db)
		biz := productbiz.NewViewProductBiz(store)

		result, err := biz.ViewProduct(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		result.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSucessResponse(result))
	}
}
