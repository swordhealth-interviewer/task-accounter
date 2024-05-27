package usecases

import (
	"errors"

	"github.com/uiansol/task-accounter.git/internal/domain/adapters"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
)

type TaskReadAllInput struct {
	User entities.User
}

type TaskReadAllOutput struct {
	Tasks []*entities.Task
}

type TaskReadAllUseCaseInterface interface {
	Execute(input TaskReadAllInput) (TaskReadAllOutput, error)
}

type TaskReadAllUseCase struct {
	TaskRepository adapters.TaskRepositoryInterface
}

func NewTaskReadAllUseCase(taskRepository adapters.TaskRepositoryInterface) TaskReadAllUseCase {
	return TaskReadAllUseCase{
		TaskRepository: taskRepository,
	}
}

func (u TaskReadAllUseCase) Execute(input TaskReadAllInput) (TaskReadAllOutput, error) {
	var tasks []*entities.Task
	var err error

	if input.User.Role == entities.UserRoleManager {
		tasks, err = u.TaskRepository.FindAll()
		if err != nil {
			return TaskReadAllOutput{}, errors.New(string(ErrorFindAllTasks) + ": " + err.Error())
		}
	} else {
		tasks, err = u.TaskRepository.FindByUserID(input.User.ID)
		if err != nil {
			return TaskReadAllOutput{}, errors.New(string(ErrorFindTasksByUser) + ": " + err.Error())
		}
	}

	output := TaskReadAllOutput{
		Tasks: tasks,
	}

	return output, nil
}
