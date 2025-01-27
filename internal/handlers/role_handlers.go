package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/qsmsoft/todo/internal/models"
	"github.com/qsmsoft/todo/internal/services"
	"net/http"
	"strconv"
)

type RoleHandler struct {
	roleService services.RoleService
}

func NewRoleHandler(roleService services.RoleService) *RoleHandler {
	return &RoleHandler{roleService: roleService}
}

func (h *RoleHandler) Create(c echo.Context) error {
	role := &models.RoleCreateRequest{}
	if err := c.Bind(role); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdRole, err := h.roleService.Create(role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, createdRole)
}

func (h *RoleHandler) Index(c echo.Context) error {
	roles, err := h.roleService.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, roles)
}

func (h *RoleHandler) Show(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	role, err := h.roleService.Get(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, role)
}

func (h *RoleHandler) Edit(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	role := new(models.RoleUpdateRequest)
	if err := c.Bind(role); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot parse JSON"})
	}

	updatedRole, err := h.roleService.Update(id, role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedRole)
}

func (h *RoleHandler) Destroy(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.roleService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
