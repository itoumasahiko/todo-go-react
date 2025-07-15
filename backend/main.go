package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/rs/cors"
	"strconv"
)

type Todo struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
}

var nextID = 3

var todos = []Todo{
    {ID: 1, Title: "タスク1"},
    {ID: 2, Title: "タスク2"},
}

func getTodos(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil || newTodo.Title == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	newTodo.ID = nextID
	nextID++
	todos = append(todos, newTodo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTodo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	newTodos := []Todo{}
	found := false
	for _, t := range todos {
		if t.ID != id {
			newTodos = append(newTodos, t)
		} else {
			found = true
		}
	}
	if !found {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	todos = newTodos
	w.WriteHeader(http.StatusNoContent)
}

func main() {
    mux := http.NewServeMux()

	mux.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTodos(w, r)
		case http.MethodPost:
			postTodo(w, r)
		case http.MethodDelete:
			deleteTodo(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	corsHandler := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type"},
        AllowCredentials: false,
    }).Handler(mux)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
