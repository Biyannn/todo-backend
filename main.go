package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"todo-backend/todo"

	"github.com/joho/godotenv"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// handle preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	godotenv.Load()

	// ✅ health check
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Backend Golang is running 🚀"))
	})

	// ✅ todo API
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todo.GetTodos(w, r)
		case http.MethodPost:
			todo.CreateTodo(w, r)
		case http.MethodDelete:
   	 	todo.DeleteAllTodos(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// ✅ todo delete API dan todo update API
	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}

	switch r.Method {
	case http.MethodDelete:
		todo.DeleteTodo(w, r, id)

	case http.MethodPut:
		todo.ToggleTodo(w, r, id)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("API running on", port)
	http.ListenAndServe(":"+port, enableCORS(http.DefaultServeMux))
}
