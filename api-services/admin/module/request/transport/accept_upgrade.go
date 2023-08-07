package reqtransport

import (
	"fmt"
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	reqbiz "shopbee/module/request/biz"
	reqmodel "shopbee/module/request/model"
	reqstorage "shopbee/module/request/storage"

	"github.com/gin-gonic/gin"
)

func AcceptRequestUpgrade(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMainDBConnection()

		var data reqmodel.RequestUpgrade

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		fmt.Print(data)
		store := reqstorage.NewSQLStore(db)
		biz := reqbiz.NewRequestUpgradeBiz(store, requester)

		if err := biz.AcceptRequestUpgrade(c.Request.Context(), data.UserId); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.UserId))
	}
}
