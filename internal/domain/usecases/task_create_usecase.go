package usecases

import (
	"errors"

	"github.com/google/uuid"
	"github.com/uiansol/task-accounter.git/internal/domain/adapters"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
)

type TaskCreateInput struct {
	Title   string
	Summary string
	User    entities.User
}

type TaskCreateOutput struct {
	Task entities.Task
}

type TaskCreateUseCaseInterface interface {
	Execute(input TaskCreateInput) (TaskCreateOutput, error)
}

type TaskCreateUseCase struct {
	TaskRepository adapters.TaskRepositoryInterface
}

func NewTaskCreateUseCase(taskRepository adapters.TaskRepositoryInterface) TaskCreateUseCase {
	return TaskCreateUseCase{
		TaskRepository: taskRepository,
	}
}

func (u TaskCreateUseCase) Execute(input TaskCreateInput) (TaskCreateOutput, error) {
	if input.User.Role != entities.UserRoleTechnician {
		return TaskCreateOutput{}, errors.New(string(ErrorTechnicianRoleRequired))
	}

	task, err := entities.NewTask(input.Title, input.Summary, input.User.ID)
	if err != nil {
		return TaskCreateOutput{}, errors.New(string(ErrorCreateTask) + ": " + err.Error())
	}

	uuid := uuid.NewString()
	task.ID = uuid

	createdTask, err := u.TaskRepository.Save(task)
	if err != nil {
		return TaskCreateOutput{}, errors.New(string(ErrorSaveTask) + ": " + err.Error())
	}

	output := TaskCreateOutput{
		Task: createdTask,
	}

	return output, nil
}
