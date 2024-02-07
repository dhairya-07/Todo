package middleware

import (
	"net/http"
	"fmt"
	"encoding/json"
	"time"

	"github.com/gorilla/mux"
	// "github.com/google/uuid"
	"github.com/dhairya-07/todo/database"
)

// type TodoService interface {
//     CreateTodo(username, title, description, status string, createdAt, updatedAt time.Time) (*database.Todo, error)
//     GetTodo(id string) (*database.Todo, error)
//     GetAllTodos(username string) ([]*database.Todo, error)
//     UpdateTodo(id, newStatus string) (string, error)
//     DeleteTodo(id string) (string, error)
// }

// type todoService struct{}

// func NewTodoService() database.TodoService {
//     return &todoService{}
// }

type TodoService interface {
	CreateTodoHandler(w http.ResponseWriter, r *http.Request)
	GetAllTodosHandler(w http.ResponseWriter, r *http.Request)
	GetTodoHandler(w http.ResponseWriter, r *http.Request)
	UpdateTodoHandler(w http.ResponseWriter, r *http.Request)
	DeleteTodoHandler(w http.ResponseWriter, r *http.Request)
}

type todoService struct{}

func NewTodoService() TodoService {
    return &todoService{}
}

// func (s *todoService) CreateTodo(username, title, description, status string, createdAt, updatedAt time.Time) (*database.Todo, error) {
//     return database.CreateTodo(username, title, description, status, createdAt, updatedAt)
// }

// func (s *todoService) GetTodo(id string) (*database.Todo, error) {
//     return database.GetTodo(id)
// }

// func (s *todoService) GetAllTodos(username string) ([]*database.Todo, error) {
//     return database.GetAllTodos(username)
// }

// func (s *todoService) UpdateTodo(id, newStatus string) (string, error) {
//     return database.UpdateTodo(id, newStatus)
// }

// func (s *todoService) DeleteTodo(id string) (string, error) {
//     return database.DeleteTodo(id)
// }

func (s *todoService) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo database.Todo
	vars := mux.Vars(r)
	user_id := vars["user_id"]
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err!=nil{
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	newTodo, err := database.CreateTodo(user_id,todo.Title, todo.Description, "Pending", time.Now(), time.Now())
	if err!=nil{
		http.Error(w,fmt.Sprintf("Error in creating todo: %v",err), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(newTodo)
}

func (s *todoService) GetAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]

	todos, err := database.GetAllTodos(user_id)
	if err!=nil{
		http.Error(w, fmt.Sprintf("Internal server error %v",err),http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

func (s *todoService) GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user_id := vars["user_id"]

	todo, err := database.GetTodo(user_id,id)
	if err!=nil{
		if err == database.ErrNotFound{
			http.Error(w, "No todo found", http.StatusNotFound)
		}else{
			http.Error(w, fmt.Sprintf("Error retrieving todo: %v",err), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (s *todoService) UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]
	id := vars["id"]

	var updatedStatus struct {
        Status string `json:"Status"`
    }

	err := json.NewDecoder(r.Body).Decode(&updatedStatus)
	if err!=nil{
		http.Error(w, "Invalid JSON format",http.StatusBadRequest)
	}

	msg, err := database.UpdateTodo(user_id,id, updatedStatus.Status)
	if err!=nil{
		http.Error(w, fmt.Sprintf("Error: %v",err), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(msg))

}

func (s *todoService) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user_id := vars["user_id"]

	msg, err := database.DeleteTodo(user_id,id)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(msg))

}