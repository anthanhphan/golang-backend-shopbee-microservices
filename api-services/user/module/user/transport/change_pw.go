package usertransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	"shopbee/component/hasher"
	userbiz "shopbee/module/user/biz"
	userstorage "shopbee/module/user/storage"

	"github.com/gin-gonic/gin"
)

func ChangePassword(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMainDBConnection()

		type Reciever struct {
			OldPass string `json:"old_pass"`
			NewPass string `json:"new_pass"`
		}

		var r Reciever

		if err := c.ShouldBind(&r); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewChangePasswordBiz(store, requester, md5)

		if err := biz.ChangePassword(
			c.Request.Context(),
			requester.GetUserId(),
			r.OldPass,
			r.NewPass,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}
