package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/qsmsoft/todo/internal/models"
	"github.com/qsmsoft/todo/internal/services"
	"net/http"
)

type CommentHandler struct {
	service services.CommentService
}

func NewCommentHandler(service services.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

func (h *CommentHandler) Store(c echo.Context) error {
	comment := new(models.CommentCreateRequest)
	if err := c.Bind(comment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot parse JSON"})
	}

	createdComment, err := h.service.Create(comment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdComment)
}

func (h *CommentHandler) Index(c echo.Context) error {
	comments, err := h.service.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) Show(c echo.Context) error {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid UUID format"})
	}

	comment, err := h.service.Get(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, comment)
}

func (h *CommentHandler) Edit(c echo.Context) error {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid UUID format"})
	}

	comment := new(models.CommentUpdateRequest)
	if err := c.Bind(comment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot parse JSON"})
	}

	updatedComment, err := h.service.Update(id, comment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedComment)
}

func (h *CommentHandler) Destroy(c echo.Context) error {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid UUID format"})
	}

	err = h.service.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}
