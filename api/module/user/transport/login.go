package usertransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	"shopbee/component/hasher"
	"shopbee/component/tokenprovider/jwt"
	userbiz "shopbee/module/user/biz"
	usermodel "shopbee/module/user/model"
	userstorage "shopbee/module/user/storage"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(err)
		}

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewLoginBiz(store, tokenProvider, md5, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(account))
	}
}
