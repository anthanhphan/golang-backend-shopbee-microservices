package main

import (
	"log"
	"os"
	"shopbee/db"
	"shopbee/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	LoadEnv()
}

func main() {
	db := db.ConnectDB(os.Getenv("MYSQL_CONN"))

	router := gin.Default()

	v1 := router.Group("v1")
	{
		v1.POST("sign-up", handler.Signup(db))
		v1.POST("sign-in", handler.Signin(db))
	}

	router.Run(":3000")
}

func LoadEnv() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
}
