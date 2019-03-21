package main

import (
	"database/sql"

	"github.com/asticode/go-astilog"
	_ "github.com/mattn/go-sqlite3"
)

// Todo represents a single todo.
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Todos is a slice of all todos.
var Todos []Todo

var index = 0
var db *sql.DB
var err error

func initDB() {
	db, err = sql.Open("sqlite3", "./todos.db")
	if err != nil {
		astilog.Error(err)
	}

	// Can't close it here as the db type is used globally.
	// defer db.Close()

	sqlStmt := `
	create table if not exists todos (id integer not null primary key, title text, done integer);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		astilog.Error("%q: %s\n", err, sqlStmt)
		return
	}

	// insert("Hallo Welt")
	// insert("Hello World!")
	// insert("Bonjour le monde")
	getAll()
	setIndex()
}

func insert(title string) (id int) {
	tx, err := db.Begin()
	if err != nil {
		astilog.Error(err)
	}
	_, err = tx.Exec("insert into todos(id, title, done) values($1, $2, 0)", index+1, title)
	if err != nil {
		astilog.Error(err)
	}
	tx.Commit()
	index++
	return index
}

func delete(id int) {
	tx, err := db.Begin()
	if err != nil {
		astilog.Error(err)
	}
	_, err = tx.Exec("delete from todos where id = $1", id)
	if err != nil {
		astilog.Error(err)
	}
	tx.Commit()
	index++
}

func update (done int, id int) {
	tx, err := db.Begin()
	if err != nil {
		astilog.Error(err)
	}
	_, err = tx.Exec("Update todos set done = $1 where id = $2", done, id)
	if err != nil {
		astilog.Error(err)
	}
	tx.Commit()
}

func getAll() {
	rows, err := db.Query("select id, title, done from todos")
	if err != nil {
		astilog.Error(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var title string
		var done int
		err := rows.Scan(&id, &title, &done)
		if err != nil {
			astilog.Error(err)
		}
		Todos = append(Todos, Todo{id, title, _intToBool(done)})
	}
}

func setIndex() {
	for _, todo := range Todos {
		index = todo.ID
	}
}

func _boolToInt(boolean interface{}) int {
	return (map[interface{}]int{false: 0, true: 1})[boolean]
}

func _intToBool(integer interface{}) bool {
	return (map[interface{}]bool{0: false, 1: true})[integer]
}