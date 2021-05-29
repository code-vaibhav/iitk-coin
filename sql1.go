package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func insertUser(db *sql.DB, firstname string, lastname string, rollNo int) {
	log.Println("Inserting user record ...")
	statement, err := db.Prepare("INSERT INTO users(firstname, lastname, rollNo) VALUES(?, ?, ?)")
	checkErr(err)
	_, err = statement.Exec(firstname, lastname, rollNo)
	checkErr(err)
	log.Println("User added ...")
}

func displayUsers(db *sql.DB) {
	var rollNo int
	var firstname, lastname string

	rows, err := db.Query("SELECT rollNo , firstname, lastname FROM users")
	checkErr(err)
	for rows.Next() {
		rows.Scan(&rollNo, &firstname, &lastname)
		fmt.Println(strconv.Itoa(rollNo) + ": " + firstname + " " + lastname)
	}
}

func createTable(db *sql.DB) {
	log.Println("Creating users table ...")
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (rollNo INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Users table created")
}

func main() {
	database, err := sql.Open("sqlite3", "./testdata.db")
	checkErr(err)
	createTable(database)
	insertUser(database, "vaibhav", "goyal", 201075)
	displayUsers(database)
}
