package main

import (
	"database/sql"
	"log"

	"github.com/code-vaibhav/iitk-coin/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func createTable(db *sql.DB) {
	log.Println("Creating users table ...")
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (rollNo INTEGER PRIMARY KEY, name TEXT NOT NULL, password TEXT NOT NULL, coins INTEGER NOT NULL)")
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Users table created")
}

var router *gin.Engine

func ApiMiddleWare(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConn", db)
		c.Next()
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./testdata.db")
	controllers.CheckErr(err)
	createTable(db)

	router = gin.Default()
	router.Use(ApiMiddleWare(db))
	controllers.SetUpRoutes(router)
	router.Run(":8080")
}
