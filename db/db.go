package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func IntiDatabase() {
	db, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic(err)
	}

	DB = db

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {

	queryCreateUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id	INTEGER PRIMARY KEY AUTOINCREMENT, 
		email TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL,  
		password TEXT NO NULL
	)
	`

	_, err := DB.Exec(queryCreateUsersTable)

	if err != nil {
		panic(err)
	}

	queryCreateEventsTable := `CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT, 
	name TEXT NOT NULL, 
	description TEXT NOT NULL, 
	location TEXT NOT NULL, 
	price TEXT NOT NULL, 
	user_id INTEGER NOT NULL, 
	FOREIGN KEY (user_id) REFERENCES users (id)
 	)`

	_, err = DB.Exec(queryCreateEventsTable)

	if err != nil {
		panic(err)
	}

	queryCreateRegisterationTable := `
		CREATE TABLE IF NOT EXISTS registration (
		id	INTEGER PRIMARY KEY AUTOINCREMENT, 
		event_id	INTEGER NOT NULL, 
		user_id	INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users (id),
		FOREIGN KEY (event_id) REFERENCES events (id)
		) 
	`

	_, err = DB.Exec(queryCreateRegisterationTable)

	if err != nil {
		panic(err)
	}

}
