package TodoController

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"todo/interface"
)
// storing todos in the todosslice memory
var todosSlice []todoInterface.Todo


func ListTodosHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todosSlice)
}


func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]

	// Find the TODO item with the specified ID
	for _, todo := range todosSlice {
		if todo.ID == todoID {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	// If the TODO item is not found, return a 404 status
	http.NotFound(w, r)
}


func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo todoInterface.Todo

	// Decode the JSON request body
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTodo.ID = fmt.Sprintf("%d", len(todosSlice)+1)

	// Add the new TODO item to the todosSlice memory
	todosSlice = append(todosSlice, newTodo)
	json.NewEncoder(w).Encode(newTodo)
}



func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]

	// Find the index of the TODO item with the specified ID
	for i, todo := range todosSlice {
		if todo.ID == todoID {
			// Decode the JSON request body into a Todo struct
			var updatedTodo todoInterface.Todo
			err := json.NewDecoder(r.Body).Decode(&updatedTodo)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Update the TODO item with the new values
			todosSlice[i] = updatedTodo

			// Return the updated TODO item as JSON
			json.NewEncoder(w).Encode(updatedTodo)
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]

	// find the todo by id
	for i, todo := range todosSlice {
		if todo.ID == todoID {
			// Remove the TODO item 
			todosSlice = append(todosSlice[:i], todosSlice[i+1:]...)
			w.Write([]byte("TODO item deleted successfully"))
			return
		}
	}
	http.NotFound(w, r)
}