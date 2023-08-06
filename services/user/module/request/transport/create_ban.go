package reqtransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	reqbiz "shopbee/module/request/biz"
	reqmodel "shopbee/module/request/model"
	reqstorage "shopbee/module/request/storage"

	"github.com/gin-gonic/gin"
)

func CreateRequestBanUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMainDBConnection()

		type ReportUser struct {
			ShopId string `json:"shop_id" gorm:"column:shop_id;"`
		}

		var rpUser ReportUser

		if err := c.ShouldBind(&rpUser); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		uid, err := common.FromBase58(rpUser.ShopId)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data reqmodel.RequestBanUser

		data.UserId = requester.GetUserId()
		data.ShopId = int(uid.GetLocalID())

		store := reqstorage.NewSQLStore(db)
		biz := reqbiz.NewCreateRequestBanUserBiz(store, requester)

		if err := biz.CreateRequestBanUser(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}
