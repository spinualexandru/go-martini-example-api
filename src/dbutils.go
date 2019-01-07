package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// General function for inserting into the database a task
func insertDatabase(identifier string, body string) int64 {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Exec(fmt.Sprintf("INSERT INTO test VALUES (%s, '%s')", identifier, body))

	if err != nil {
		panic(err.Error())
	}
	id, err := insert.LastInsertId()
	return id
}

// General function for retrieving a task in SQL Data format
func getTask(id string) *sql.Row {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res := db.QueryRow(fmt.Sprintf("SELECT * FROM test.test WHERE id=%s", id))
	if err != nil {
		panic(err.Error())
	}

	return res
}

// General function for retrieving a task and converting it to JSON String
func getTaskJSON(taskID string) []byte {
	var id int64
	var body string
	getTask(taskID).Scan(&id, &body)
	task := task{id, body}
	encoded, err := json.Marshal(task)
	if err != nil {
		panic(err.Error())
	}
	return encoded
}
