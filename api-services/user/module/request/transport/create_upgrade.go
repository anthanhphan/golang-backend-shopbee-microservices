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

func CreateRequestUpgrade(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMainDBConnection()

		var data reqmodel.RequestUpgrade

		data.UserId = requester.GetUserId()

		store := reqstorage.NewSQLStore(db)
		biz := reqbiz.NewRequestUpgradeBiz(store, requester)

		if err := biz.CreateRequestUpgrade(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}
