package adapters

import "github.com/uiansol/task-accounter.git/internal/domain/entities"

type TaskRepositoryInterface interface {
	Save(task entities.Task) (*entities.Task, error)
	FindByID(id string) (*entities.Task, error)
	FindAll() ([]*entities.Task, error)
	FindByUserID(userID string) ([]*entities.Task, error)
	Delete(id string) error
}
