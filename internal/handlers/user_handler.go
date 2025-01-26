package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/qsmsoft/todo/internal/models"
	"github.com/qsmsoft/todo/internal/services"
	"net/http"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Store(c echo.Context) error {
	user := new(models.UserCreateRequest)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot parse JSON"})
	}

	createdUser, err := h.service.Create(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) Index(c echo.Context) error {
	users, err := h.service.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Show(c echo.Context) error {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid UUID format"})

	}

	user, err := h.service.Get(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Edit(c echo.Context) error {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid UUID format"})
	}

	user := new(models.UserUpdateRequest)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot parse JSON"})
	}

	updatedUser, err := h.service.Update(id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) Destroy(c echo.Context) error {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid UUID format"})
	}

	err = h.service.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}
