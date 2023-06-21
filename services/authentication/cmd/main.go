package main

import (
	"fmt"
	"log"
	"os"
	"shopbee/db"
	"shopbee/routes"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var PORT int
var RUNENV string
var DSN string

func init() {
	// Load local.env file to open and read
	err := godotenv.Load("../local.env")
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}

	// Read port
	p, _ := strconv.Atoi(os.Getenv("PORT"))
	PORT = p

	// Read run environment
	RUNENV = os.Getenv("RUNENV")

	// Read DSN string to connect db
	DSN = os.Getenv("MYSQL_CONN_DEV")

	if RUNENV == "pro" {
		DSN = os.Getenv("MYSQL_CONN_PRO")
	}

	fmt.Println(">>> SERVICE: " + os.Getenv("SERVICE_NAME"))
	fmt.Println(">>> RUNNING ON PORT:", PORT)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	db := db.ConnectDB(DSN)

	routes.SetupAuthRoutes(db).Run(fmt.Sprintf(":%d", PORT))
}
