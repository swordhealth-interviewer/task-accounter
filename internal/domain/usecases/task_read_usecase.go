package usecases

import (
	"errors"
	"os"

	"github.com/uiansol/task-accounter.git/internal/domain/adapters"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/encrypt"
)

type TaskReadInput struct {
	ID   string
	User entities.User
}

type TaskReadOutput struct {
	Task *entities.Task
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
	task, err := u.TaskRepository.FindByID(input.ID)
	if err != nil {
		return TaskReadOutput{}, err
	}

	if task.OwnerID != input.User.ID && input.User.Role != entities.UserRoleManager {
		return TaskReadOutput{}, errors.New(string(ErrorTaskNotOwnedByUser))
	}

	decText, err := encrypt.Decrypt(task.Summary, os.Getenv("SUMMARY_SECRET"))
	if err != nil {
		return TaskReadOutput{}, errors.New(string(ErrorCryptSummary))
	}
	task.Summary = decText

	output := TaskReadOutput{
		Task: task,
	}

	return output, nil
}
