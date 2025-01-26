package repositories

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/qsmsoft/todo/internal/database"
	"github.com/qsmsoft/todo/internal/models"
	"github.com/qsmsoft/todo/internal/utils"
)

type UserRepository interface {
	Create(user *models.UserCreateRequest) (*models.User, error)
	List() ([]*models.User, error)
	Get(id uuid.UUID) (*models.User, error)
	Update(id uuid.UUID, user *models.UserUpdateRequest) (*models.User, error)
	Delete(id uuid.UUID) error
}

type userRepository struct {
	db *database.Database
}

func NewUserRepository(db *database.Database) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.UserCreateRequest) (*models.User, error) {
	existingUser, _ := r.FindByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	_, err := r.db.Conn.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		user.Name, user.Email, user.Password)
	if err != nil {
		return nil, errors.New("user not created")
	}

	var createdUser models.User
	err = r.db.Conn.Get(&createdUser, "SELECT * FROM users WHERE email = $1", user.Email)
	if err != nil {
		return nil, errors.New("user not fetched")
	}

	return &createdUser, nil
}

func (r *userRepository) List() ([]*models.User, error) {
	var users []*models.User

	err := r.db.Conn.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, errors.New("users not fetched")
	}

	return users, nil
}

func (r *userRepository) Get(id uuid.UUID) (*models.User, error) {
	var user models.User

	err := r.db.Conn.Get(&user, "SELECT * FROM users WHERE uuid = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("user not fetched")
	}

	return &user, nil
}

func (r *userRepository) Update(id uuid.UUID, user *models.UserUpdateRequest) (*models.User, error) {
	var name, email string

	checkingUser, _ := r.Get(id)
	if checkingUser == nil {
		return nil, errors.New("user not found")
	}

	if user.Name != "" {
		name = user.Name
	} else {
		err := r.db.Conn.Get(&name, "SELECT name FROM users WHERE uuid = $1", id)
		if err != nil {
			return nil, errors.New("name not fetched")
		}
	}

	if user.Email != "" {
		email = user.Email

		if existingUser, _ := r.FindByEmail(email); existingUser != nil {
			return nil, errors.New("user already exists")
		}

		if !utils.IsValidEmail(email) {
			return nil, errors.New("email format is invalid")
		}
	} else {
		err := r.db.Conn.Get(&email, "SELECT email FROM users WHERE uuid = $1", id)
		if err != nil {
			return nil, errors.New("email not fetched")
		}
	}

	_, err := r.db.Conn.Exec("UPDATE users SET name = $1, email = $2 WHERE uuid = $3",
		name, email, id)
	if err != nil {
		return nil, errors.New("user not updated")
	}

	var updatedUser models.User

	err = r.db.Conn.Get(&updatedUser, "SELECT * FROM users WHERE uuid = $1", id)
	if err != nil {
		return nil, errors.New("user not fetched")
	}

	return &updatedUser, nil
}

func (r *userRepository) Delete(id uuid.UUID) error {
	_, err := r.db.Conn.Exec("DELETE FROM users WHERE uuid = $1", id)
	if err != nil {
		return errors.New("user not deleted")
	}

	return nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.db.Conn.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
	}
	return &user, nil
}
