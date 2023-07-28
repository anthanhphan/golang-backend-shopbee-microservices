package usertransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	userbiz "shopbee/module/user/biz"
	usermodel "shopbee/module/user/model"
	userstorage "shopbee/module/user/storage"

	"github.com/gin-gonic/gin"
)

func ListUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter usermodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewListUserBiz(store, requester)

		result, err := biz.ListUser(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
