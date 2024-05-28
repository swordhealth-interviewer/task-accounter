package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
	"github.com/uiansol/task-accounter.git/internal/domain/mocks"
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
)

func TestNewTaskReadUseCase(t *testing.T) {
	t.Run("should return a task read use case", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		encrypterMock := mocks.NewEncrypterInterface(t)
		TaskReadUseCase := usecases.NewTaskReadUseCase(taskRepositoryMock, encrypterMock)

		assert.NotNil(t, TaskReadUseCase)
		assert.NotNil(t, TaskReadUseCase.TaskRepository)
		assert.NotNil(t, TaskReadUseCase.Encrypter)
	})
}

func TestTaskReadUseCaseExecute(t *testing.T) {
	t.Run("should return a task when user is the owner", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		encrypterMock := mocks.NewEncrypterInterface(t)
		TaskReadUseCase := usecases.NewTaskReadUseCase(taskRepositoryMock, encrypterMock)

		task := &entities.Task{
			ID:      "task-id",
			Title:   "Task Title",
			Summary: "Task Description",
			OwnerID: "user-id",
			Status:  entities.Open,
		}

		taskRepositoryMock.On("FindByID", "task-id").Return(task, nil)
		encrypterMock.On("Decrypt", mock.Anything).Return("Task Description", nil)

		output, err := TaskReadUseCase.Execute(usecases.TaskReadInput{
			ID: "task-id",
			User: entities.User{
				ID:   "user-id",
				Role: entities.UserRoleManager,
			},
		})

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, task, output.Task)
	})

	t.Run("should return a task when user is a manager", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		encrypterMock := mocks.NewEncrypterInterface(t)
		TaskReadUseCase := usecases.NewTaskReadUseCase(taskRepositoryMock, encrypterMock)

		task := &entities.Task{
			ID:      "task-id",
			Title:   "Task Title",
			Summary: "Task Description",
			OwnerID: "user-id",
			Status:  entities.Open,
		}

		taskRepositoryMock.On("FindByID", "task-id").Return(task, nil)
		encrypterMock.On("Decrypt", mock.Anything).Return("Task Description", nil)

		output, err := TaskReadUseCase.Execute(usecases.TaskReadInput{
			ID: "task-id",
			User: entities.User{
				ID:   "other-id",
				Role: entities.UserRoleManager,
			},
		})

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, task, output.Task)
	})

	t.Run("should return an error when task is not owned by user", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		encrypterMock := mocks.NewEncrypterInterface(t)
		TaskReadUseCase := usecases.NewTaskReadUseCase(taskRepositoryMock, encrypterMock)

		task := &entities.Task{
			ID:      "task-id",
			Title:   "Task Title",
			Summary: "Task Description",
			OwnerID: "user-id",
			Status:  entities.Open,
		}

		taskRepositoryMock.On("FindByID", "task-id").Return(task, nil)

		output, err := TaskReadUseCase.Execute(usecases.TaskReadInput{
			ID: "task-id",
			User: entities.User{
				ID:   "user-id-2",
				Role: entities.UserRoleTechnician,
			},
		})

		assert.NotNil(t, err)
		assert.Equal(t, usecases.TaskReadOutput{}, output)
		assert.Equal(t, string(usecases.ErrorTaskNotOwnedByUser), err.Error())
	})
}
