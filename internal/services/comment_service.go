package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/qsmsoft/todo/internal/models"
	"github.com/qsmsoft/todo/internal/repositories"
)

type CommentService interface {
	Create(comment *models.CommentCreateRequest) (*models.Comment, error)
	List() ([]*models.Comment, error)
	Get(id uuid.UUID) (*models.Comment, error)
	Update(id uuid.UUID, comment *models.CommentUpdateRequest) (*models.Comment, error)
	Delete(id uuid.UUID) error
}

type commentService struct {
	commentRepository repositories.CommentRepository
}

func NewCommentService(repository repositories.CommentRepository) CommentService {
	return &commentService{commentRepository: repository}
}

func (s *commentService) Create(comment *models.CommentCreateRequest) (*models.Comment, error) {
	if comment.Content == "" {
		return nil, errors.New("content is required")
	}

	if comment.TaskID == 0 {
		return nil, errors.New("task_id is required")
	}

	if comment.UserID == 0 {
		return nil, errors.New("user_id is required")
	}

	return s.commentRepository.Create(comment)
}

func (s *commentService) List() ([]*models.Comment, error) {
	return s.commentRepository.List()
}

func (s *commentService) Get(id uuid.UUID) (*models.Comment, error) {
	return s.commentRepository.Get(id)
}

func (s *commentService) Update(id uuid.UUID, comment *models.CommentUpdateRequest) (*models.Comment, error) {
	existingComment, _ := s.Get(id)
	if existingComment == nil {
		return nil, errors.New("comment not found")
	}

	return s.commentRepository.Update(id, comment)
}

func (s *commentService) Delete(id uuid.UUID) error {
	existingComment, _ := s.Get(id)
	if existingComment == nil {
		return errors.New("comment not found")
	}

	return s.commentRepository.Delete(id)
}
