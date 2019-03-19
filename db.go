package main

import (
	"database/sql"
	"os"

	"github.com/asticode/go-astilog"
	_ "github.com/mattn/go-sqlite3"
)

var index = 0
var db *sql.DB
var err error

func initDB() {
	os.Remove("./todos.db")

	db, err = sql.Open("sqlite3", "./todos.db")
	if err != nil {
		astilog.Error(err)
	}

	// Can't close it here as the db type is used globally.
	// defer db.Close()

	sqlStmt := `
	create table todos (id integer not null primary key, title text, done integer);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		astilog.Error("%q: %s\n", err, sqlStmt)
		return
	}

	// insert("Hallo Welt")
	// insert("Hello World!")
	// insert("Bonjour le monde")
}

func insert(title string) {
	tx, err := db.Begin()
	if err != nil {
		astilog.Error(err)
	}
	stmt, err := tx.Exec("insert into todos(id, title, done) values($1, $2, 0)", index, title)
	if err != nil {
		astilog.Error(err)
	}
	astilog.Error(stmt)
	tx.Commit()
	index++
}
