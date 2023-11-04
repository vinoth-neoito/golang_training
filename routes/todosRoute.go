package todosRoute

import (
    "github.com/gorilla/mux"
	"todo/controller"
)

// SetupTodoRoutes sets up routes for the "/todos" path
func SetupTodoRoutes(r *mux.Router) {
	r.HandleFunc("/", TodoController.ListTodosHandler).Methods("GET")
	r.HandleFunc("/{id}", TodoController.GetTodoHandler).Methods("GET")
	r.HandleFunc("/", TodoController.CreateTodoHandler).Methods("POST")
	r.HandleFunc("/{id}",TodoController.UpdateTodoHandler).Methods("PUT")
	r.HandleFunc("/{id}", TodoController.DeleteTodoHandler).Methods("DELETE")
}









