package todo

import (
	"encoding/json"
	"net/http"
	"time"
	// "strconv"
)

type APIResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

// Method Get
// GET
func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, _ := loadFromFile()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(APIResponse{
		Message: "success",
		Data:    todos,
	})
}


// Method Post
// POST
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req Todo
	json.NewDecoder(r.Body).Decode(&req)

	if req.CreatedAt == "" {
		req.CreatedAt = time.Now().Format(time.RFC3339)
	}

	Create(req.Title, req.CreatedAt, req.Priority)

	todos, _ := loadFromFile()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(APIResponse{
		Message: "created",
		Data:    todos,
	})
}



// Method Delete
// DELETE
func DeleteTodo(w http.ResponseWriter, r *http.Request, id int) {
	if !DeleteByID(id) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(APIResponse{Message: "not found"})
		return
	}

	todos, _ := loadFromFile()
	json.NewEncoder(w).Encode(APIResponse{
		Message: "deleted",
		Data:    todos,
	})
}

// DELETE ALL
func DeleteAllTodos(w http.ResponseWriter, r *http.Request) {
	todos = []Todo{}
	_ = saveToFile(todos)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(APIResponse{
		Message: "all todos deleted",
		Data:    todos,
	})
}


// Method Put/update
// TOGGLE
func ToggleTodo(w http.ResponseWriter, r *http.Request, id int) {
	if _, success := ToggleDoneByID(id); !success {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(APIResponse{Message: "not found"})
		return
	}

	todos, _ := loadFromFile()
	json.NewEncoder(w).Encode(APIResponse{
		Message: "updated",
		Data:    todos,
	})
}