package main

import (
	"fmt"
	"log"
	"os"
	"shopbee/component/appctx"
	"shopbee/component/uploadprovider"
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

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	fmt.Print(s3Domain)
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appCtx := appctx.NewAppContext(DB, s3Provider, SECRETKEY)
	router.Use(middleware.Recover(appCtx))
	router.Use(cors.New(
		cors.Config{
			AllowAllOrigins: true,
			AllowHeaders:    []string{"Origin", "Content-Type", "Accept", "Authorization"},
		},
	))

	route.UserRouterInit(router, appCtx)
	route.ProductRouterInit(router, appCtx)
	route.RequestRouterInit(router, appCtx)
	route.UploadRouterInit(router, appCtx)

	if err := router.Run(":" + PORT); err != nil {
		log.Fatalln(err.Error())
	}
}
