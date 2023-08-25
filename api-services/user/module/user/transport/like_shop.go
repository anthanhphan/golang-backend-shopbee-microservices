package usertransport

import (
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	userbiz "shopbee/module/user/biz"
	userstorage "shopbee/module/user/storage"

	"github.com/gin-gonic/gin"
)

func LikeShop(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewLikeShopBiz(store, requester)

		if err := biz.LikeShop(
			c.Request.Context(),
			requester.GetUserId(),
			int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}

func DislikeShop(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewLikeShopBiz(store, requester)

		if err := biz.DislikeShop(
			c.Request.Context(),
			requester.GetUserId(),
			int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}

func IsLikedShop(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewLikeShopBiz(store, requester)

		if err := biz.IsLikedShop(
			c.Request.Context(),
			requester.GetUserId(),
			int(uid.GetLocalID())); !err {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}

func CountLikeShop(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewLikeShopBiz(store, requester)

		count := biz.CountLikeShop(c.Request.Context(), int(uid.GetLocalID()))
		c.JSON(http.StatusOK, gin.H{
			"like": count,
		})
	}
}
