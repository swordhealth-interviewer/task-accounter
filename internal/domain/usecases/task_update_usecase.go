package usecases

import (
	"errors"
	"fmt"
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

	task, err := u.TaskRepository.FindByID(input.TaskID)
	if err != nil {
		return TaskUpdateOutput{}, errors.New(string(ErrorFindTaskByID) + ": " + err.Error())
	}

	if task.OwnerID != input.User.ID {
		return TaskUpdateOutput{}, errors.New(string(ErrorTaskNotOwnedByUser))
	}

	if task.Status == entities.Closed {
		return TaskUpdateOutput{}, errors.New(string(ErrorTaskClosed))
	}

	err = entities.ValidateTaskParameters(input.Title, input.Summary)
	if err != nil {
		return TaskUpdateOutput{}, errors.New(string(ErrorInvalidTaskData) + ": " + err.Error())
	}

	task.Title = input.Title
	task.Summary = input.Summary

	if input.CloseTask {
		task.Status = entities.Closed
		task.DoneAt = time.Now()
	}

	updatedTask, err := u.TaskRepository.Save(*task)
	if err != nil {
		return TaskUpdateOutput{}, errors.New(string(ErrorSaveTask) + ": " + err.Error())
	}

	if input.CloseTask {
		userPrint := input.User.Name + "<" + input.User.Email + ">"
		taskPrint := task.Title + "<" + task.ID + ">"
		// TODO: send to message broker
		fmt.Println("The tech " + userPrint + " performed the task" + taskPrint + " on date " + task.DoneAt.String())
	}

	output := TaskUpdateOutput{
		Task: *updatedTask,
	}

	return output, nil
}
