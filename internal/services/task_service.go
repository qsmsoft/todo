package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/qsmsoft/todo/internal/models"
	"github.com/qsmsoft/todo/internal/repositories"
)

type TaskService interface {
	Create(task *models.TaskCreateRequest) (*models.Task, error)
	List() ([]*models.Task, error)
	Get(id uuid.UUID) (*models.Task, error)
	Update(id uuid.UUID, task *models.TaskUpdateRequest) (*models.Task, error)
	Delete(id uuid.UUID) error
}

type taskService struct {
	taskRepository repositories.TaskRepository
}

func NewTaskService(repository repositories.TaskRepository) TaskService {
	return &taskService{taskRepository: repository}
}

func (s *taskService) Create(task *models.TaskCreateRequest) (*models.Task, error) {
	if task.Title == "" {
		return nil, errors.New("title is required")
	}

	if task.Description == "" {
		return nil, errors.New("description is required")
	}

	if task.UserID == 0 {
		return nil, errors.New("user_id is required")
	}

	return s.taskRepository.Create(task)
}

func (s *taskService) List() ([]*models.Task, error) {
	return s.taskRepository.List()
}

func (s *taskService) Get(id uuid.UUID) (*models.Task, error) {
	return s.taskRepository.Get(id)
}

func (s *taskService) Update(id uuid.UUID, task *models.TaskUpdateRequest) (*models.Task, error) {
	existingTask, _ := s.Get(id)
	if existingTask == nil {
		return nil, errors.New("task not found")
	}

	return s.taskRepository.Update(id, task)
}

func (s *taskService) Delete(id uuid.UUID) error {
	existingTask, _ := s.Get(id)
	if existingTask == nil {
		return errors.New("task not found")
	}

	return s.taskRepository.Delete(id)
}
