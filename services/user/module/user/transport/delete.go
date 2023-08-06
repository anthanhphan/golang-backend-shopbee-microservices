package usertransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	userbiz "shopbee/module/user/biz"
	userstorage "shopbee/module/user/storage"

	"github.com/gin-gonic/gin"
)

func DeleteUserWithId(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMainDBConnection()
		// id, err := strconv.Atoi(c.Param("id"))
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewDeleteUserBiz(store, requester)

		if err := biz.DeleteUser(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}

func DeleteUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMainDBConnection()
		// id, err := strconv.Atoi(c.Param("id"))

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewDeleteUserBiz(store, requester)

		if err := biz.DeleteUser(c.Request.Context(), requester.GetUserId()); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}
