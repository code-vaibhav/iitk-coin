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

	log.Println("Creating history table ...")

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
	log.Println("Users table created")
}
