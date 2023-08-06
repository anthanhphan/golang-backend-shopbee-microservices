package usertransport

import (
	"net/http"
	"shopbee/component/appctx"
	userbiz "shopbee/module/user/biz"
	userstorage "shopbee/module/user/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewFindUserBiz(store)

		result, err := biz.FindUserById(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"status": result.Status,
		})
	}
}
