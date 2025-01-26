package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/qsmsoft/todo/internal/models"
	"github.com/qsmsoft/todo/internal/repositories"
	"github.com/qsmsoft/todo/internal/utils"
)

type UserService interface {
	Create(user *models.UserCreateRequest) (*models.User, error)
	List() ([]*models.User, error)
	Get(id uuid.UUID) (*models.User, error)
	Update(id uuid.UUID, user *models.UserUpdateRequest) (*models.User, error)
	Delete(id uuid.UUID) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{userRepository: repository}
}

func (s *userService) Create(user *models.UserCreateRequest) (*models.User, error) {
	if user.Email == "" {
		return nil, errors.New("email is required")
	}

	if !utils.IsValidEmail(user.Email) {
		return nil, errors.New("email format is invalid")
	}

	if user.Password == "" {
		return nil, errors.New("password is required")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("password not hashed")
	}

	user.Password = hashedPassword

	return s.userRepository.Create(user)
}

func (s *userService) List() ([]*models.User, error) {
	return s.userRepository.List()
}

func (s *userService) Get(id uuid.UUID) (*models.User, error) {
	return s.userRepository.Get(id)
}

func (s *userService) Update(id uuid.UUID, user *models.UserUpdateRequest) (*models.User, error) {
	existingUser, _ := s.Get(id)
	if existingUser == nil {
		return nil, errors.New("user not found")
	}

	return s.userRepository.Update(id, user)
}

func (s *userService) Delete(id uuid.UUID) error {
	existingUser, _ := s.Get(id)
	if existingUser == nil {
		return errors.New("user not found")
	}

	return s.userRepository.Delete(id)
}
