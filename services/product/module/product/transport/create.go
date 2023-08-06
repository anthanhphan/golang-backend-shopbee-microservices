package producttransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	productbiz "shopbee/module/product/biz"
	productmodel "shopbee/module/product/model"
	productstorage "shopbee/module/product/storage"

	"github.com/gin-gonic/gin"
)

func CreateProduct(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMainDBConnection()

		var data productmodel.Product

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data.ShopId = requester.GetUserId()

		store := productstorage.NewSQLStore(db)
		biz := productbiz.NewCreateProductBiz(store, requester)

		if err := biz.CreateProduct(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.GenUID(common.DbTypeProduct)
		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.FakeId.String()))
	}
}
