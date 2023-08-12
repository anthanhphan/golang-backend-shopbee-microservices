package reqtransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	reqbiz "shopbee/module/request/biz"
	reqstorage "shopbee/module/request/storage"

	"github.com/gin-gonic/gin"
)

func AcceptRequestUpgrade(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := reqstorage.NewSQLStore(db)
		biz := reqbiz.NewRequestUpgradeBiz(store, requester)

		if err := biz.AcceptRequestUpgrade(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}
