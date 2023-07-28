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

func UpdateUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data usermodel.UserUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewUpdateUserBiz(store, requester)

		if err := biz.UpdateUser(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}
