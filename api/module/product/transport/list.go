package producttransport

import (
	"fmt"
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	productbiz "shopbee/module/product/biz"
	productmodel "shopbee/module/product/model"
	productstorage "shopbee/module/product/storage"

	"github.com/gin-gonic/gin"
)

func ListProduct(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		type FilterTemp struct {
			ShopId string `json:"shop_id,omitempty" form:"shop_id"`
		}

		var filterTemp FilterTemp
		var filter productmodel.Filter

		if err := c.ShouldBind(&filterTemp); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		uid, err := common.FromBase58(filterTemp.ShopId)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		fmt.Print(int(uid.GetLocalID()))

		store := productstorage.NewSQLStore(db)
		biz := productbiz.NewListProductBiz(store)

		result, err := biz.ListProduct(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
