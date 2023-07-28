package reqtransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	reqbiz "shopbee/module/request/biz"
	reqmodel "shopbee/module/request/model"
	reqstorage "shopbee/module/request/storage"

	"github.com/gin-gonic/gin"
)

func ListRequetUpgrade(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter reqmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := reqstorage.NewSQLStore(db)
		biz := reqbiz.NewListRequetUpgradeBiz(store, requester)

		result, err := biz.ListRequetUpgrade(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
