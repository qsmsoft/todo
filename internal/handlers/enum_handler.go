package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/qsmsoft/todo/internal/services"
	"net/http"
)

type EnumHandler struct {
	service services.EnumService
}

func NewEnumHandler(service services.EnumService) *EnumHandler {
	return &EnumHandler{service: service}
}

func (h *EnumHandler) GetTaskStatuses(c echo.Context) error {
	return c.JSON(http.StatusOK, h.service.GetTaskStatuses())
}
