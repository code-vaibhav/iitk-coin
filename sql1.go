package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	firstname string
	lastname  string
	rollNo    int
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func insertUser(db *sql.DB, stu *user) {
	log.Println("Inserting user record ...")
	statement, err := db.Prepare("INSERT INTO users(firstname, lastname, rollNo) VALUES(?, ?, ?)")
	checkErr(err)
	_, err = statement.Exec(stu.firstname, stu.lastname, stu.rollNo)
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
	os.Remove("testdata.db")
	log.Println("Creating users table ...")
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (rollNo INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Users table created")
}

func main() {
	var users = []user{}
	users = append(users, user{"vaibhav", "goyal", 201075})
	users = append(users, user{"Nikhil", "Rathore", 200635})
	users = append(users, user{"Aastha", "Sitpal", 200134})
	users = append(users, user{"Harshal", "Mehta", 200845})
	users = append(users, user{"Rose", "Aggarwal", 200983})
	users = append(users, user{"Rini", "Sahu", 209983})

	fmt.Println(users)
	database, err := sql.Open("sqlite3", "./testdata.db")
	checkErr(err)
	createTable(database)
	for _, stu := range users {
		insertUser(database, &stu)
	}
	displayUsers(database)
}
