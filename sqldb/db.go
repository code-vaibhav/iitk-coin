package sqldb

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("sqlite3", "./testdata.db")
	checkErr(err)

	DB = db
	createTable()
}

func createTable() {
	//users table
	log.Println("Creating users table ...")

	schema := `CREATE TABLE IF NOT EXISTS users (
		rollNo INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		password TEXT NOT NULL,
		coins TEXT NOT NULL,
		isAdmin INTEGER NOT NULL,
		isFreezed INTEGER NOT NULL
	);`

	statement, err := DB.Prepare(schema)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Users table created")

	//transactions table
	log.Println("Creating transactions table ...")

	schema = `CREATE TABLE IF NOT EXISTS transactions (
		sender INTEGER REFERENCES users(rollNo),
		reciever INTEGET NOT NULL REFERENCES users(rollNo),
		amount INTEGER NOT NULL,
		type TEXT NOT NULL,
		madeAt INTEGER NOT NULL
	);`

	statement, err = DB.Prepare(schema)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("transactions table created")

	//items table
	log.Println("Creating transactions table ...")

	schema = `CREATE TABLE IF NOT EXISTS items (
		code INTEGER PRIMARY KEY,
		amount INTEGER NOT NULL,
		name TEXT NOT NULL,
		isAvailable INTEGER NOT NULL
	)`

	statement, err = DB.Prepare(schema)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Users table created")

	// redeem_requests table
	log.Println("Creating redeem_requests table ...")

	schema = `CREATE TABLE IF NOT EXISTS redeem_requests (
		user INTEGER NOT NULL REFERENCES users(rollNo),
		itemCode INTEGER NOT NULL REFERENCES items(code),
		status TEXT NOT NULL
	);`

	statement, err = DB.Prepare(schema)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Users table created")
}
