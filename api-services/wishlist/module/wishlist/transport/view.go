package wishlisttransport

import (
	"fmt"
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	wishlistbiz "shopbee/module/wishlist/biz"
	wishliststorage "shopbee/module/wishlist/storage"

	"github.com/gin-gonic/gin"
)

func ViewMyWishList(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMainDBConnection()

		store := wishliststorage.NewSQLStore(db)
		biz := wishlistbiz.NewViewMyWishListBiz(store, requester)

		wishlist, err := biz.ViewMyWishList(c.Request.Context(), int(requester.GetUserId()))

		if err != nil {
			panic(err)
		}

		var productIdList []string
		for i := range wishlist {
			fmt.Print(i)
			uid := common.NewUID(uint32(wishlist[i].ProductId), 1, 1)
			fakeId := &uid
			fmt.Println(fakeId.String())
			productIdList = append(productIdList, fakeId.String())
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(productIdList))
	}
}
