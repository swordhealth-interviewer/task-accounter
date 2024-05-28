package usecases

import (
	"errors"

	"github.com/uiansol/task-accounter.git/internal/domain/adapters"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
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
	Encrypter      adapters.EncrypterInterface
}

func NewTaskReadUseCase(taskRepository adapters.TaskRepositoryInterface, encrypter adapters.EncrypterInterface) TaskReadUseCase {
	return TaskReadUseCase{
		TaskRepository: taskRepository,
		Encrypter:      encrypter,
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

	decText, err := u.Encrypter.Decrypt(task.Summary)
	if err != nil {
		return TaskReadOutput{}, errors.New(string(ErrorCryptSummary))
	}
	task.Summary = decText

	output := TaskReadOutput{
		Task: task,
	}

	return output, nil
}
