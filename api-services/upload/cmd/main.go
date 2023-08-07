package main

import (
	"log"
	"os"
	"shopbee/common"
	"shopbee/component/appctx"
	"shopbee/component/uploadprovider"
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

	common.RegisDiscovery("upload", PORT)
}

func main() {
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appCtx := appctx.NewAppContext(s3Provider)

	router := gin.Default()
	router.Use(middleware.Recover(appCtx))
	router.Use(cors.New(
		cors.Config{
			AllowAllOrigins: true,
			AllowHeaders:    []string{"Origin", "Content-Type", "Accept", "Authorization"},
		},
	))

	routes.UploadRouterInit(router, appCtx)

	if err := router.Run(":" + PORT); err != nil {
		log.Fatalln(err.Error())
	}
}
