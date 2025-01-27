package services

import (
	"errors"
	"github.com/qsmsoft/todo/internal/models"
	"github.com/qsmsoft/todo/internal/repositories"
)

type RoleService interface {
	Create(role *models.RoleCreateRequest) (*models.Role, error)
	List() ([]*models.Role, error)
	Get(id int) (*models.Role, error)
	Update(id int, role *models.RoleUpdateRequest) (*models.Role, error)
	Delete(id int) error
}

type roleService struct {
	roleRepository repositories.RoleRepository
}

func NewRoleService(repository repositories.RoleRepository) RoleService {
	return &roleService{roleRepository: repository}
}

func (s *roleService) Create(role *models.RoleCreateRequest) (*models.Role, error) {
	if role.Name == "" {
		return nil, errors.New("role name is required")
	}

	return s.roleRepository.Create(role)
}

func (s *roleService) List() ([]*models.Role, error) {
	return s.roleRepository.List()
}

func (s *roleService) Get(id int) (*models.Role, error) {
	return s.roleRepository.Get(id)
}

func (s *roleService) Update(id int, role *models.RoleUpdateRequest) (*models.Role, error) {
	existingRole, _ := s.roleRepository.Get(id)
	if existingRole == nil {
		return nil, errors.New("role not found")
	}

	return s.roleRepository.Update(id, role)
}

func (s *roleService) Delete(id int) error {
	existingRole, _ := s.roleRepository.Get(id)
	if existingRole == nil {
		return errors.New("role not found")
	}
	return s.roleRepository.Delete(id)
}
