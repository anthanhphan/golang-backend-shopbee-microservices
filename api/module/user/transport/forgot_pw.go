package usertransport

import (
	"fmt"
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	userbiz "shopbee/module/user/biz"
	userstorage "shopbee/module/user/storage"

	"github.com/gin-gonic/gin"
)

func ForgotPassword(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		type Reciever struct {
			Email string `json:"email"`
		}

		var data Reciever

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewForgotPWBiz(store)

		fmt.Print(data.Email)
		if err := biz.ForgotPassword(c.Request.Context(), data.Email); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}
