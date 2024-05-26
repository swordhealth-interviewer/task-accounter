package usecases

import (
	"errors"

	"github.com/uiansol/task-accounter.git/internal/domain/adapters"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
)

type TaskDeleteInput struct {
	TaskID string
	User   entities.User
}

type TaskDeleteUseCaseInterface interface {
	Execute(input TaskDeleteInput) error
}

type TaskDeleteUseCase struct {
	TaskRepository adapters.TaskRepositoryInterface
}

func NewTaskDeleteUseCase(taskRepository adapters.TaskRepositoryInterface) TaskDeleteUseCase {
	return TaskDeleteUseCase{
		TaskRepository: taskRepository,
	}
}

func (u TaskDeleteUseCase) Execute(input TaskDeleteInput) error {
	if input.User.Role != entities.UserRoleManager {
		return errors.New(string(ErrorManagerRoleRequired))
	}

	err := u.TaskRepository.Delete(input.TaskID)
	if err != nil {
		return errors.New(string(ErrorDeleteTask) + ": " + err.Error())
	}

	return nil
}
