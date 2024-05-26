package usecases

import (
	"errors"

	"github.com/uiansol/task-accounter.git/internal/domain/adapters"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
)

type TaskReadInput struct {
	User entities.User
}

type TaskReadOutput struct {
	Tasks []*entities.Task
}

type TaskReadUseCaseInterface interface {
	Execute(input TaskReadInput) (TaskReadOutput, error)
}

type TaskReadUseCase struct {
	TaskRepository adapters.TaskRepositoryInterface
}

func NewTaskReadUseCase(taskRepository adapters.TaskRepositoryInterface) TaskReadUseCase {
	return TaskReadUseCase{
		TaskRepository: taskRepository,
	}
}

func (u TaskReadUseCase) Execute(input TaskReadInput) (TaskReadOutput, error) {
	var tasks []*entities.Task
	var err error

	if input.User.Role == entities.UserRoleManager {
		tasks, err = u.TaskRepository.FindAll()
		if err != nil {
			return TaskReadOutput{}, errors.New(string(ErrorReadAllTasks) + ": " + err.Error())
		}
	} else {
		tasks, err = u.TaskRepository.FindByUserID(input.User.ID)
		if err != nil {
			return TaskReadOutput{}, errors.New(string(ErrorReadTasksByUser) + ": " + err.Error())
		}
	}

	output := TaskReadOutput{
		Tasks: tasks,
	}

	return output, nil
}
