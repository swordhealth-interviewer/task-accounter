package usecases

import (
	"errors"

	"github.com/google/uuid"
	"github.com/uiansol/task-accounter.git/internal/domain/adapters"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
)

type TaskCreateUseCaseError string

const (
	summaryMaxLength int = 2500

	technicianRoleRequiredError TaskCreateUseCaseError = "only technicians can create tasks"
	emptyTitleError             TaskCreateUseCaseError = "title is required"
	emptySummaryError           TaskCreateUseCaseError = "summary is required"
	summaryMaxLengthError       TaskCreateUseCaseError = "summary must have a maximum of 2500 characters"
	saveError                   TaskCreateUseCaseError = "usecase error saving task"
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
		return TaskCreateOutput{}, errors.New(string(technicianRoleRequiredError))
	}

	if input.Title == "" {
		return TaskCreateOutput{}, errors.New(string(emptyTitleError))
	}

	if input.Summary == "" {
		return TaskCreateOutput{}, errors.New(string(emptySummaryError))
	}

	if len(input.Summary) > summaryMaxLength {
		return TaskCreateOutput{}, errors.New(string(summaryMaxLengthError))
	}

	task := entities.NewTask(input.Title, input.Summary, input.User.ID)

	uuid := uuid.NewString()
	task.ID = uuid

	createdTask, err := u.TaskRepository.Save(task)
	if err != nil {
		return TaskCreateOutput{}, errors.New(string(saveError) + ": " + err.Error())
	}

	return TaskCreateOutput{Task: createdTask}, nil
}
