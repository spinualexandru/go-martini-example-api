package main

import "github.com/go-martini/martini"

func main() {
	// Initialize Server
	m := martini.Classic()
	// Initialize the first route group
	m.Group("/tasks", tasks)
	m.Run()
}
