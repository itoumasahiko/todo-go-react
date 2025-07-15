package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/rs/cors"
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

func main() {
    mux := http.NewServeMux()

	mux.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getTodos(w, r)
		} else if r.Method == http.MethodPost {
			postTodo(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	handler := cors.Default().Handler(mux)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
