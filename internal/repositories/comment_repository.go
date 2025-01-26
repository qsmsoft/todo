package repositories

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/qsmsoft/todo/internal/database"
	"github.com/qsmsoft/todo/internal/models"
)

type CommentRepository interface {
	Create(comment *models.CommentCreateRequest) (*models.Comment, error)
	List() ([]*models.Comment, error)
	Get(id uuid.UUID) (*models.Comment, error)
	Update(id uuid.UUID, comment *models.CommentUpdateRequest) (*models.Comment, error)
	Delete(id uuid.UUID) error
}

type commentRepository struct {
	db *database.Database
}

func NewCommentRepository(db *database.Database) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment *models.CommentCreateRequest) (*models.Comment, error) {
	_, err := r.db.Conn.Exec("INSERT INTO comments (content, user_id, task_id, parent_id) VALUES ($1, $2, $3, $4)",
		comment.Content, comment.UserID, comment.TaskID, comment.ParentID)
	if err != nil {
		return nil, errors.New("comment not created")
	}

	var createdComment models.Comment
	err = r.db.Conn.Get(&createdComment, "SELECT * FROM comments WHERE content = $1", comment.Content)
	if err != nil {
		return nil, errors.New("comment not fetched")
	}

	return &createdComment, nil
}

func (r *commentRepository) List() ([]*models.Comment, error) {
	var comments []*models.Comment

	err := r.db.Conn.Select(&comments, "SELECT * FROM comments")
	if err != nil {
		return nil, errors.New("comments not fetched " + err.Error())
	}

	return comments, nil
}

func (r *commentRepository) Get(id uuid.UUID) (*models.Comment, error) {
	var comment models.Comment

	err := r.db.Conn.Get(&comment, "SELECT * FROM comments WHERE uuid = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("comment not found")
		}
		return nil, errors.New("comment not fetched")
	}

	return &comment, nil
}

func (r *commentRepository) Update(id uuid.UUID, comment *models.CommentUpdateRequest) (*models.Comment, error) {
	var (
		content  string
		parentID *int
	)

	if comment.Content != "" {
		content = comment.Content
	} else {
		err := r.db.Conn.Get(&content, "SELECT content FROM comments WHERE uuid = $1", id)
		if err != nil {
			return nil, errors.New("content not fetched")
		}
	}

	if comment.ParentID != nil {
		parentID = comment.ParentID
	} else {
		err := r.db.Conn.Get(&parentID, "SELECT parent_id FROM comments WHERE uuid = $1", id)
		if err != nil {
			return nil, errors.New("parent_id not fetched")
		}
	}

	_, err := r.db.Conn.Exec("UPDATE comments SET content = $1, parent_id = $2 WHERE uuid = $3",
		content, parentID, id)
	if err != nil {
		return nil, errors.New("comment not updated")
	}

	var updatedComment models.Comment
	err = r.db.Conn.Get(&updatedComment, "SELECT * FROM comments WHERE uuid = $1", id)
	if err != nil {
		return nil, errors.New("comment not fetched")
	}

	return &updatedComment, nil
}

func (r *commentRepository) Delete(id uuid.UUID) error {
	_, err := r.db.Conn.Exec("DELETE FROM comments WHERE uuid = $1", id)
	if err != nil {
		return errors.New("comment not deleted")
	}
	return nil
}
