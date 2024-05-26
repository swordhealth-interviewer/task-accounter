package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
	"github.com/uiansol/task-accounter.git/internal/domain/mocks"
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
)

func TestNewTaskDeleteUseCase(t *testing.T) {
	t.Run("should return a task delete use case", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskDeleteUsecase := usecases.NewTaskDeleteUseCase(taskRepositoryMock)

		assert.NotNil(t, taskDeleteUsecase)
		assert.NotNil(t, taskDeleteUsecase.TaskRepository)
	})
}

func TestTaskDeleteUseCaseExecute(t *testing.T) {
	t.Run("should delete a task and return error nil", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskDeleteUsecase := usecases.NewTaskDeleteUseCase(taskRepositoryMock)

		task := usecases.TaskDeleteInput{
			TaskID: "task-id",
			User: entities.User{
				Role: entities.UserRoleManager,
			},
		}

		taskRepositoryMock.On("Delete", "task-id").Return(nil)

		err := taskDeleteUsecase.Execute(task)

		assert.Nil(t, err)
	})

	t.Run("should return an error when user is not a manager", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskDeleteUsecase := usecases.NewTaskDeleteUseCase(taskRepositoryMock)

		task := usecases.TaskDeleteInput{
			TaskID: "task-id",
			User: entities.User{
				Role: entities.UserRoleTechnician,
			},
		}

		err := taskDeleteUsecase.Execute(task)

		assert.NotNil(t, err)
		assert.Equal(t, string(usecases.ErrorManagerRoleRequired), err.Error())
	})

	t.Run("should return an error when task repository returns an error", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskDeleteUsecase := usecases.NewTaskDeleteUseCase(taskRepositoryMock)

		task := usecases.TaskDeleteInput{
			TaskID: "task-id",
			User: entities.User{
				Role: entities.UserRoleManager,
			},
		}

		taskRepositoryMock.On("Delete", "task-id").Return(assert.AnError)

		err := taskDeleteUsecase.Execute(task)

		assert.NotNil(t, err)
		assert.Equal(t, string(usecases.ErrorDeleteTask)+": assert.AnError general error for testing", err.Error())
	})
}
