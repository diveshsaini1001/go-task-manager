package task

import (
	"errors"
	"sync"
)

type TaskRepository struct {
	mutex sync.RWMutex
	tasks map[string]*Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[string]*Task),
	}
}

func (r *TaskRepository) Create(task *Task) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.tasks[task.ID] = task
}

func (r *TaskRepository) Get(id string) (*Task, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (r *TaskRepository) Update(id string, updated *Task) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	_, exists := r.tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	r.tasks[id] = updated
	return nil
}

func (r *TaskRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	_, exists := r.tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}
