package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diveshsaini1001/go-task-manager/internal/task"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/health", healthCheckHandler)

	// Task routes
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			task.CreateTask(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			task.GetTask(w, r)
		case http.MethodPut:
			task.UpdateTask(w, r)
		case http.MethodDelete:
			task.DeleteTask(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
