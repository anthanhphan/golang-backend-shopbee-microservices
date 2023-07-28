package usertransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	"shopbee/component/hasher"
	userbiz "shopbee/module/user/biz"
	usermodel "shopbee/module/user/model"
	userstorage "shopbee/module/user/storage"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.FakeId.String()))
	}
}
