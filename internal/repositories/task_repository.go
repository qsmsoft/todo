package repositories

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/qsmsoft/todo/internal/database"
	"github.com/qsmsoft/todo/internal/models"
)

type TaskRepository interface {
	Create(task *models.TaskCreateRequest) (*models.Task, error)
	List() ([]*models.Task, error)
	Get(id uuid.UUID) (*models.Task, error)
	Update(id uuid.UUID, task *models.TaskUpdateRequest) (*models.Task, error)
	Delete(id uuid.UUID) error
	UpdateStatus(id uuid.UUID, status int) (*models.Task, error)
}

type taskRepository struct {
	db *database.Database
}

func NewTaskRepository(db *database.Database) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *models.TaskCreateRequest) (*models.Task, error) {
	_, err := r.db.Conn.Exec("INSERT INTO tasks (title, description, user_id) VALUES ($1, $2, $3)",
		task.Title, task.Description, task.UserID)
	if err != nil {
		return nil, errors.New("task not created")
	}

	var createdTask models.Task
	err = r.db.Conn.Get(&createdTask, "SELECT * FROM tasks WHERE title = $1", task.Title)
	if err != nil {
		return nil, errors.New("task not fetched")
	}

	return &createdTask, nil
}

func (r *taskRepository) List() ([]*models.Task, error) {
	var tasks []*models.Task

	err := r.db.Conn.Select(&tasks, "SELECT * FROM tasks")
	if err != nil {
		return nil, errors.New("tasks not fetched")
	}

	return tasks, nil
}

func (r *taskRepository) Get(id uuid.UUID) (*models.Task, error) {
	var task models.Task

	err := r.db.Conn.Get(&task, "SELECT * FROM tasks WHERE uuid = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("task not found")
		}
		return nil, errors.New("task not fetched")
	}
	return &task, nil
}

func (r *taskRepository) Update(id uuid.UUID, task *models.TaskUpdateRequest) (*models.Task, error) {
	var (
		title, description string
		status             int
	)

	if task.Title != "" {
		title = task.Title
	} else {
		err := r.db.Conn.Get(&title, "SELECT title FROM tasks WHERE uuid = $1", id)
		if err != nil {
			return nil, errors.New("title not fetched")
		}
	}

	if task.Description != "" {
		description = task.Description
	} else {
		err := r.db.Conn.Get(&description, "SELECT description FROM tasks WHERE uuid = $1", id)
		if err != nil {
			return nil, errors.New("description not fetched")
		}
	}

	if task.Status != 0 {
		status = task.Status
	} else {
		err := r.db.Conn.Get(&status, "SELECT status FROM tasks WHERE uuid = $1", id)
		if err != nil {
			return nil, errors.New("status not fetched")
		}
	}

	_, err := r.db.Conn.Exec("UPDATE tasks SET title = $1, description = $2, status = $3 WHERE uuid = $4",
		title, description, status, id)
	if err != nil {
		return nil, errors.New("task not updated")
	}

	var updatedTask models.Task

	err = r.db.Conn.Get(&updatedTask, "SELECT * FROM tasks WHERE uuid = $1", id)
	if err != nil {
		return nil, errors.New("task not fetched")
	}

	return &updatedTask, nil
}

func (r *taskRepository) Delete(id uuid.UUID) error {
	_, err := r.db.Conn.Exec("DELETE FROM tasks WHERE uuid = $1", id)
	if err != nil {
		return errors.New("task not deleted")
	}

	return nil
}

func (r *taskRepository) UpdateStatus(id uuid.UUID, status int) (*models.Task, error) {
	_, err := r.db.Conn.Exec("UPDATE tasks SET status = $1 WHERE uuid = $2", status, id)
	if err != nil {
		return nil, errors.New("task not updated")
	}

	var updatedTask models.Task

	err = r.db.Conn.Get(&updatedTask, "SELECT * FROM tasks WHERE uuid = $1", id)
	if err != nil {
		return nil, errors.New("task not fetched")
	}

	return &updatedTask, nil
}
