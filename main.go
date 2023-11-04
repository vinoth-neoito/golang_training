package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"todo/routes"
)

func main() {
	r := mux.NewRouter()

	// Route
	todosRoute.SetupTodoRoutes(r.PathPrefix("/todos").Subrouter())

	// Start the server
	http.Handle("/", r)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}