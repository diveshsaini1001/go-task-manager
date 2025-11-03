package task

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Handler struct {
	service *TaskService
}

func NewHandler(service *TaskService) *Handler {
	return &Handler{service: service}
}

// CreateTask handles POST /tasks
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Description string `json:"description"`
		OwnerID     string `json:"owner_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	task := h.service.CreateTask(body.Description, body.OwnerID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// GetTask handles GET /tasks/{id}
func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	task, err := h.service.GetTask(id)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

// UpdateTask handles PUT /tasks/{id}
func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	var body struct {
		Description string `json:"Description"`
		IsCompleted bool   `json:"is_completed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	updated, err := h.service.UpdateTask(id, body.Description, body.IsCompleted)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

// DeleteTask handles DELETE /tasks/{id}
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	err := h.service.DeleteTask(id)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
