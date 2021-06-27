package main

import (
	"github.com/code-vaibhav/iitk-coin/controllers"
	"github.com/code-vaibhav/iitk-coin/sqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var router *gin.Engine

func main() {
	sqldb.ConnectDB()

	router = gin.Default()
	controllers.SetUpRoutes(router)
	router.Run(":8080")
}
