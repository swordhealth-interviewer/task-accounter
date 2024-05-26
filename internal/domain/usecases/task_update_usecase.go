package usecases

import (
	"errors"

	"github.com/uiansol/task-accounter.git/internal/domain/adapters"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
)

type TaskUpdateInput struct {
	Title   string
	Summary string
	Status  entities.TaskStatus
	User    entities.User
}

type TaskUpdateOutput struct {
	Task entities.Task
}

type TaskUpdateUseCaseInterface interface {
	Execute(input TaskUpdateInput) (TaskUpdateOutput, error)
}

type TaskUpdateUseCase struct {
	TaskRepository adapters.TaskRepositoryInterface
}

func NewTaskUpdateUseCase(taskRepository adapters.TaskRepositoryInterface) TaskUpdateUseCase {
	return TaskUpdateUseCase{
		TaskRepository: taskRepository,
	}
}

func (u TaskUpdateUseCase) Execute(input TaskUpdateInput) (TaskUpdateOutput, error) {
	if input.User.Role != entities.UserRoleTechnician {
		return TaskUpdateOutput{}, errors.New(string(ErrorTechnicianRoleRequired))
	}

	return TaskUpdateOutput{}, nil
}
