package repositories

import (
	"database/sql"
	"errors"
	"github.com/qsmsoft/todo/internal/database"
	"github.com/qsmsoft/todo/internal/models"
)

type RoleRepository interface {
	Create(role *models.RoleCreateRequest) (*models.Role, error)
	List() ([]*models.Role, error)
	Get(id int) (*models.Role, error)
	Update(id int, role *models.RoleUpdateRequest) (*models.Role, error)
	Delete(id int) error
}

type roleRepository struct {
	db *database.Database
}

func NewRoleRepository(db *database.Database) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(role *models.RoleCreateRequest) (*models.Role, error) {
	existingRole, _ := r.GetByName(role.Name)
	if existingRole != nil {
		return nil, errors.New("role already exists")
	}

	_, err := r.db.Conn.Exec("INSERT INTO roles (name) VALUES ($1)", role.Name)
	if err != nil {
		return nil, errors.New("role not created")
	}

	var createdRole models.Role
	err = r.db.Conn.Get(&createdRole, "SELECT * FROM roles WHERE name = $1", role.Name)
	if err != nil {
		return nil, errors.New("role not fetched")
	}

	return &createdRole, nil
}

func (r *roleRepository) List() ([]*models.Role, error) {
	var roles []*models.Role

	err := r.db.Conn.Select(&roles, "SELECT * FROM roles")
	if err != nil {
		return nil, errors.New("roles not fetched")
	}

	return roles, nil
}

func (r *roleRepository) Get(id int) (*models.Role, error) {
	var role models.Role

	err := r.db.Conn.Get(&role, "SELECT * FROM roles WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("role not found")
		}
		return nil, errors.New("role not fetched")
	}

	return &role, nil
}

func (r *roleRepository) Update(id int, role *models.RoleUpdateRequest) (*models.Role, error) {
	existingRole, _ := r.GetByName(role.Name)
	if existingRole != nil {
		return nil, errors.New("role already exists")
	}

	_, err := r.db.Conn.Exec("UPDATE roles SET name = $1 WHERE id = $2", role.Name, id)
	if err != nil {
		return nil, errors.New("role not updated")
	}

	var updatedRole models.Role
	err = r.db.Conn.Get(&updatedRole, "SELECT * FROM roles WHERE id = $1", id)
	if err != nil {
		return nil, errors.New("role not fetched")
	}

	return &updatedRole, nil
}

func (r *roleRepository) Delete(id int) error {
	_, err := r.db.Conn.Exec("DELETE FROM roles WHERE id = $1", id)
	if err != nil {
		return errors.New("role not deleted")
	}

	return nil
}

func (r *roleRepository) GetByName(name string) (*models.Role, error) {
	var role models.Role

	err := r.db.Conn.Get(&role, "SELECT * FROM roles WHERE name = $1", name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("role not found")
		}
	}
	return &role, nil
}
