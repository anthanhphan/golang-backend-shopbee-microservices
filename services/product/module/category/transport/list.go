package categorytransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	categorybiz "shopbee/module/category/biz"
	categorystorage "shopbee/module/category/storage"

	"github.com/gin-gonic/gin"
)

func ListCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		store := categorystorage.NewSQLStore(db)
		biz := categorybiz.NewListReportBiz(store)

		result, err := biz.ListCategory(c.Request.Context(), &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].RId = result[i].Id
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, nil))
	}
}
