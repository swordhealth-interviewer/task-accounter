package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
	"github.com/uiansol/task-accounter.git/internal/domain/mocks"
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
)

func TestNewTaskCreateUseCase(t *testing.T) {
	t.Run("should return a task create use case", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskCreateUsecase := usecases.NewTaskCreateUseCase(taskRepositoryMock)

		assert.NotNil(t, taskCreateUsecase)
		assert.NotNil(t, taskCreateUsecase.TaskRepository)
	})
}

func TestTaskCreateUseCaseExecute(t *testing.T) {
	t.Run("should create a task and return it with error nil", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskCreateUsecase := usecases.NewTaskCreateUseCase(taskRepositoryMock)

		task := usecases.TaskCreateInput{
			Title:   "Task Title",
			Summary: "Task Description",
			User: entities.User{
				ID:   "user-id",
				Role: entities.UserRoleTechnician,
			},
		}

		taskRepositoryMock.On("Create", mock.Anything).Return("task-id", nil)

		output, err := taskCreateUsecase.Execute(task)

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, "task-id", output.TaskID)
	})

	t.Run("should return an error when user is not a technician", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskCreateUsecase := usecases.NewTaskCreateUseCase(taskRepositoryMock)

		task := usecases.TaskCreateInput{
			User: entities.User{
				Role: entities.UserRoleManager,
			},
		}

		output, err := taskCreateUsecase.Execute(task)

		assert.NotNil(t, err)
		assert.Equal(t, string(usecases.ErrorTechnicianRoleRequired), err.Error())
		assert.Equal(t, usecases.TaskCreateOutput{}, output)
	})

	t.Run("should return an error when task repository fails to save", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskCreateUsecase := usecases.NewTaskCreateUseCase(taskRepositoryMock)

		taskRepositoryMock.On("Create", mock.Anything).Return("", assert.AnError)

		input := usecases.TaskCreateInput{
			Title:   "Task Title",
			Summary: "Task Description",
			User: entities.User{
				ID:   "user-id",
				Role: entities.UserRoleTechnician,
			},
		}

		output, err := taskCreateUsecase.Execute(input)

		assert.NotNil(t, err)
		assert.Equal(t, string(usecases.ErrorSaveTask)+": assert.AnError general error for testing", err.Error())
		assert.Equal(t, usecases.TaskCreateOutput{}, output)
	})
}
