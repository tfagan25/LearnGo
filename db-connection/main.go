package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test-db.db")
	if err != nil {
		log.Fatal("Error setting up db")
	}
	defer db.Close()

	sqlStmt := `CREATE TABLE IF NOT EXISTS users (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			email TEXT
		);
		`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

	// addUser := `
	// INSERT INTO users (name, email) VALUES(
	// 	"Trevor", 
	// 	"trevorfagan77@gmail.com"
	// );
	// `
	// _, err = db.Exec(addUser)
	// if err != nil {
	// 	log.Fatalf("%q: %s\n", err, addUser)
	// }

	showUsers := `SELECT * FROM users;`
	rows, err := db.Query(showUsers)
	if err != nil {
		log.Fatalf("%q: %s\n", err, showUsers)
	}
	defer rows.Close()
	
	fmt.Println("Users:")
	for rows.Next() {
		var id int
		var name, email string
		rows.Scan(&id, &name, &email)
		fmt.Printf(" - [%d] %s <%s>\n", id, name, email)
	}
}
