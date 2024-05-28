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
	TaskID string
}

type TaskCreateUseCaseInterface interface {
	Execute(input TaskCreateInput) (TaskCreateOutput, error)
}

type TaskCreateUseCase struct {
	TaskRepository adapters.TaskRepositoryInterface
	Encrypter      adapters.EncrypterInterface
}

func NewTaskCreateUseCase(taskRepository adapters.TaskRepositoryInterface, encrypter adapters.EncrypterInterface) TaskCreateUseCase {
	return TaskCreateUseCase{
		TaskRepository: taskRepository,
		Encrypter:      encrypter,
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

	encText, err := u.Encrypter.Encrypt(task.Summary)
	if err != nil {
		return TaskCreateOutput{}, errors.New(string(ErrorCryptSummary) + ": " + err.Error())
	}
	task.Summary = encText

	uuid := uuid.NewString()
	task.ID = uuid

	taskID, err := u.TaskRepository.Create(task)
	if err != nil {
		return TaskCreateOutput{}, errors.New(string(ErrorSaveTask) + ": " + err.Error())
	}

	output := TaskCreateOutput{
		TaskID: taskID,
	}

	return output, nil
}
