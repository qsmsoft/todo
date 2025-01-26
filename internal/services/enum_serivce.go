package services

import "github.com/qsmsoft/todo/internal/enums"

type EnumService interface {
	GetTaskStatuses() map[string]int
}

type enumService struct{}

func NewEnumService() EnumService {
	return &enumService{}
}

func (s *enumService) GetTaskStatuses() map[string]int {
	return map[string]int{
		enums.InProgress.String(): int(enums.InProgress),
		enums.ToDo.String():       int(enums.ToDo),
		enums.Completed.String():  int(enums.Completed),
	}
}
