package main

import (
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
)

// Define the data structure of our task
type task struct {
	ID   int64
	Body string
}

// Define the route handler for adding tasks
func addTask(r *http.Request) []byte {
	newTask := insertDatabase(r.FormValue("id"), r.FormValue("body"))
	return getTaskJSON(strconv.FormatInt(newTask, 10))
}

// Define the route handler for retrieving tasks
func retrieveTask(params martini.Params) []byte {
	return getTaskJSON(params["id"])
}

// Defining the handler for our subgroup which we initialized in main.go
func tasks(r martini.Router) {
	// Define the GET request. It will be available at GET http://localhost:3030/tasks/:id
	r.Get("/:id", retrieveTask)

	// Define the POST request. It will be available at POST http://localhost:3030/tasks/
	r.Post("/", addTask)
}
