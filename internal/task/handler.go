package task

import (
	"fmt"
	"net/http"
)

// CreateTask handles POST /tasks
func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create Task - still working on it!!!")
}

// GetTask handles GET /tasks/{id}
func GetTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get Task - still working on it!!!")
}

// UpdateTask handles PUT /tasks/{id}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Update Task - still working on it!!!")
}

// DeleteTask handles DELETE /tasks/{id}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete Task - still working on it!!!")
}
