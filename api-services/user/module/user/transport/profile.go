package usertransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"

	"github.com/gin-gonic/gin"
)

func Profile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		// fmt.Print(u.GetUserId())
		c.JSON(http.StatusOK, common.SimpleSucessResponse(u))
	}
}
