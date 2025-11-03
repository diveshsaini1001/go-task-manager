package task

import "github.com/google/uuid"

type TaskService struct {
	repo *TaskRepository
}

func NewTaskService(repo *TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(description, ownerID string) *Task {
	task := &Task{
		ID:          uuid.NewString(),
		Description: description,
		OwnerID:     ownerID,
		IsCompleted: false,
	}
	s.repo.Create(task)
	return task
}

func (s *TaskService) GetTask(id string) (*Task, error) {
	return s.repo.Get(id)
}

func (s *TaskService) UpdateTask(id, description string, completed bool) (*Task, error) {
	existing, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	existing.Description = description
	existing.IsCompleted = completed
	s.repo.Update(id, existing)
	return existing, nil
}

func (s *TaskService) DeleteTask(id string) error {
	return s.repo.Delete(id)
}
