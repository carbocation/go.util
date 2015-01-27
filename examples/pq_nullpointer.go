package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Thing struct {
	Id int64
}

var (
	db     *sql.DB                       //Global db connection
	router *mux.Router = mux.NewRouter() //Global router
)

func main() {
	db, err := initDb()

	id, err := getter()
}

func getter() (*Thing, error) {
	stmt, err := db.Prepare(`SELECT id FROM entries WHERE id=$1`)
	defer stmt.Close()

	if err != nil {
		return New(), err
	}
	
	rows, err := stmt.Query(root, user.GetId())
	if err != nil {
		return New(), err
	}
	
	// Iterate over the rows
	for rows.Next() {
		var e *Thing = New()
		err = rows.Scan(&e.Id)
		if err != nil {
			return e, err
		}

		entries[e.Id] = e
		relationships = append(relationships, map[string]int64{"Parent": ancestor, "Child": e.Id})
	}

	return New(), err
}

func New() *Thing {
	return &Thing{}
}

func initDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s user=%s password=%s port=%s sslmode=disable",
		"projects",
		"askbitcoin",
		"xnkxglie",
		"5432"))
	db.SetMaxIdleConns(95)
	if err != nil {
		fmt.Println("Panic: " + err.Error())
		panic(err)
	}

	return db, err
}
