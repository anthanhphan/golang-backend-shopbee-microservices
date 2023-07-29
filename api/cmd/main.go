package main

import (
	"fmt"
	"log"
	"os"
	"shopbee/component/appctx"
	dbconn "shopbee/database"
	"shopbee/middleware"
	route "shopbee/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	PORT      string
	DSN       string
	SECRETKEY string
	DB        *gorm.DB
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalln(err.Error())
	}

	env := os.Getenv("RUN_ENV")

	// Default is DEV environment
	PORT = os.Getenv("APP_DEV_PORT")
	DSN = os.Getenv("DB_DEV_DSN")
	SECRETKEY = os.Getenv("SECRET_KEY")
	if env == "PRO" {
		PORT = os.Getenv("APP_PRO_PORT")
		DSN = os.Getenv("DB_PRO_DSN")
	} else {
		fmt.Println("APP PORT:", PORT)
		fmt.Println("DB DSN:", DSN)
	}

	DB = dbconn.ConnectDB(DSN)
	DB = DB.Debug()
}

func main() {
	router := gin.Default()
	router.Use(middleware.Recover(appctx.NewAppContext(DB, SECRETKEY)))
	router.Use(cors.Default())

	route.UserRouterInit(router, DB)
	route.ProductRouterInit(router, DB)
	route.RequestRouterInit(router, DB)

	if err := router.Run(":" + PORT); err != nil {
		log.Fatalln(err.Error())
	}
}
