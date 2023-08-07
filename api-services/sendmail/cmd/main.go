package main

import (
	"log"
	"os"
	"shopbee/common"
	"shopbee/component/appctx"
	"shopbee/middleware"
	"shopbee/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	PORT string
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalln(err.Error())
	}

	env := os.Getenv("RUN_ENV")

	// Default is DEV environment
	PORT = os.Getenv("APP_DEV_PORT")
	if env == "PRO" {
		PORT = os.Getenv("APP_PRO_PORT")
	}

	common.RegisDiscovery("mail", PORT)
}

func main() {
	appCtx := appctx.NewAppContext()

	router := gin.Default()
	router.Use(middleware.Recover(appCtx))
	router.Use(cors.New(
		cors.Config{
			AllowAllOrigins: true,
			AllowHeaders:    []string{"Origin", "Content-Type", "Accept", "Authorization"},
		},
	))

	routes.MailServiceRouterInit(router, appCtx)

	if err := router.Run(":" + PORT); err != nil {
		log.Fatalln(err.Error())
	}
}
