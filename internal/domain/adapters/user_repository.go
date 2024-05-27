package adapters

import "github.com/uiansol/task-accounter.git/internal/domain/entities"

type UserRepositoryInterface interface {
	FindByUsername(username string) (*entities.User, string, error)
}
