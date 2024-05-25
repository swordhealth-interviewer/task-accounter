package adapters

import "github.com/uiansol/task-accounter.git/internal/domain/entities"

type TaskRepositoryInterface interface {
	Save(task entities.Task) (entities.Task, error)
}
