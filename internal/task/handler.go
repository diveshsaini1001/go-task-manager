package task

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/diveshsaini1001/go-task-manager/internal/logger"
)

type Handler struct {
	service *TaskService
	logger  *logger.Logger
}

func NewHandler(service *TaskService, logger *logger.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Description string `json:"description"`
		OwnerID     string `json:"owner_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		h.logger.Error("invalid request body in CreateTask")
		return
	}

	task := h.service.CreateTask(body.Description, body.OwnerID)
	h.logger.Info("Task created: " + task.ID + " by " + task.OwnerID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	task, err := h.service.GetTask(id)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		h.logger.Error("failed to get task: " + id)
		return
	}
	h.logger.Info("Fetched task: " + id)
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	currentUser := r.Header.Get("X-User-ID")

	existing, err := h.service.GetTask(id)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		h.logger.Error("update failed: task not found - " + id)
		return
	}

	if existing.OwnerID != currentUser {
		http.Error(w, "forbidden: not the owner of this task", http.StatusForbidden)
		if currentUser != "" {
			h.logger.Error("unauthorized update attempt by " + currentUser)
		} else {
			h.logger.Error("unauthorized update attempt by unknown user")
		}
		return
	}

	var body struct {
		Description string `json:"description"`
		IsCompleted bool   `json:"is_completed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		h.logger.Error("invalid request body in UpdateTask")
		return
	}

	updated, err := h.service.UpdateTask(id, body.Description, body.IsCompleted)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		h.logger.Error("update failed for task: " + id)
		return
	}

	h.logger.Info("Task updated: " + id + " by " + currentUser)
	json.NewEncoder(w).Encode(updated)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	currentUser := r.Header.Get("X-User-ID")

	existing, err := h.service.GetTask(id)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		h.logger.Error("delete failed: task not found - " + id)
		return
	}

	if existing.OwnerID != currentUser {
		http.Error(w, "forbidden: not the owner of this task", http.StatusForbidden)
		h.logger.Error("unauthorized delete attempt by " + currentUser)
		return
	}

	err = h.service.DeleteTask(id)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		h.logger.Error("delete failed: " + id)
		return
	}

	h.logger.Info("Task deleted: " + id + " by " + currentUser)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("DONE")
}
