package usecases

import (
	"errors"
	"time"

	"github.com/uiansol/task-accounter.git/internal/domain/adapters"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
)

type TaskUpdateInput struct {
	TaskID    string
	Title     string
	Summary   string
	CloseTask bool
	User      entities.User
}

type TaskUpdateUseCaseInterface interface {
	Execute(input TaskUpdateInput) error
}

type TaskUpdateUseCase struct {
	TaskRepository adapters.TaskRepositoryInterface
	Encrypter      adapters.EncrypterInterface
	Publisher      adapters.MessagePublisherInterface
}

func NewTaskUpdateUseCase(taskRepository adapters.TaskRepositoryInterface, encrypter adapters.EncrypterInterface, publisher adapters.MessagePublisherInterface) TaskUpdateUseCase {
	return TaskUpdateUseCase{
		TaskRepository: taskRepository,
		Encrypter:      encrypter,
		Publisher:      publisher,
	}
}

func (u TaskUpdateUseCase) Execute(input TaskUpdateInput) error {
	if input.User.Role != entities.UserRoleTechnician {
		return errors.New(string(ErrorTechnicianRoleRequired))
	}

	task, err := u.TaskRepository.FindByID(input.TaskID)
	if err != nil {
		return errors.New(string(ErrorFindTaskByID) + ": " + err.Error())
	}

	if task.OwnerID != input.User.ID {
		return errors.New(string(ErrorTaskNotOwnedByUser))
	}

	if task.Status == entities.Closed {
		return errors.New(string(ErrorTaskClosed))
	}

	err = entities.ValidateTaskParameters(input.Title, input.Summary)
	if err != nil {
		return errors.New(string(ErrorInvalidTaskData) + ": " + err.Error())
	}

	task.Title = input.Title

	encText, err := u.Encrypter.Encrypt(input.Summary)
	if err != nil {
		return errors.New(string(ErrorCryptSummary) + ": " + err.Error())
	}
	task.Summary = encText

	if input.CloseTask {
		task.Status = entities.Closed
		task.DoneAt = time.Now()
	}

	_, err = u.TaskRepository.Save(*task)
	if err != nil {
		return errors.New(string(ErrorSaveTask) + ": " + err.Error())
	}

	if input.CloseTask {
		go func() {
			userPrint := input.User.Name + "<" + input.User.Email + ">"
			taskPrint := task.Title + "<" + task.ID + ">"
			msg := "The tech " + userPrint + " performed the task" + taskPrint + " on date " + task.DoneAt.String()
			u.Publisher.PublishMessages(msg, "task-done-queue")
		}()

	}

	return nil
}
