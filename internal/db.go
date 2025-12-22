package internal

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func CreateDB() *sql.DB {
	db, err := sql.Open("sqlite", "./my.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = execScoreTableCreation(db)
	if err != nil {
		fmt.Println("execScoreTableCreation: ", err)
		return nil
	}

	return db
}

func execScoreTableCreation(db *sql.DB) error {
	sqlStmt := `
	create table if not exists scores (id integer not null primary key autoincrement, "user" text, session text unique, value integer, created_at datetime default current_timestamp);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}

	return nil
}
