package routes

import (
	"shopbee/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("v1/auth")
	{
		v1.POST("sign-up", handler.Signup(db))
		v1.POST("sign-in", handler.Signin(db))
	}

	return router
}
