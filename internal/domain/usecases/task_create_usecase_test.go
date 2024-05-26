package usecases_test

import (
	"strings"
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

		taskRepositoryMock.On("Save", mock.Anything).Return(entities.Task{
			ID:      "task-id",
			Title:   "Task Title",
			Summary: "Task Description",
			Owner:   "user-id",
			Status:  entities.Open,
		}, nil)

		output, err := taskCreateUsecase.Execute(task)

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, "task-id", output.Task.ID)
		assert.Equal(t, task.Title, output.Task.Title)
		assert.Equal(t, task.Summary, output.Task.Summary)
		assert.Equal(t, task.User.ID, output.Task.Owner)
		assert.Equal(t, entities.Open, output.Task.Status)
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
		assert.Equal(t, "only technicians can create tasks", err.Error())
		assert.Equal(t, usecases.TaskCreateOutput{}, output)
	})

	t.Run("should return an error when title is empty", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskCreateUsecase := usecases.NewTaskCreateUseCase(taskRepositoryMock)

		task := usecases.TaskCreateInput{
			User: entities.User{
				Role: entities.UserRoleTechnician,
			},
		}

		output, err := taskCreateUsecase.Execute(task)

		assert.NotNil(t, err)
		assert.Equal(t, "title is required", err.Error())
		assert.Equal(t, usecases.TaskCreateOutput{}, output)
	})

	t.Run("should return an error when summary is empty", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskCreateUsecase := usecases.NewTaskCreateUseCase(taskRepositoryMock)

		task := usecases.TaskCreateInput{
			Title: "Task Title",
			User: entities.User{
				Role: entities.UserRoleTechnician,
			},
		}

		output, err := taskCreateUsecase.Execute(task)

		assert.NotNil(t, err)
		assert.Equal(t, "summary is required", err.Error())
		assert.Equal(t, usecases.TaskCreateOutput{}, output)
	})

	t.Run("should return an error when summary is longer than 2500 characters", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskCreateUsecase := usecases.NewTaskCreateUseCase(taskRepositoryMock)

		task := usecases.TaskCreateInput{
			Title:   "Task Title",
			Summary: strings.Repeat("a", 2501),
			User: entities.User{
				Role: entities.UserRoleTechnician,
			},
		}

		output, err := taskCreateUsecase.Execute(task)

		assert.NotNil(t, err)
		assert.Equal(t, "summary must have a maximum of 2500 characters", err.Error())
		assert.Equal(t, usecases.TaskCreateOutput{}, output)
	})

	t.Run("should return an error when task repository fails to save", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskCreateUsecase := usecases.NewTaskCreateUseCase(taskRepositoryMock)

		taskRepositoryMock.On("Save", mock.Anything).Return(entities.Task{}, assert.AnError)

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
		assert.Equal(t, "error saving task: assert.AnError general error for testing", err.Error())
		assert.Equal(t, usecases.TaskCreateOutput{}, output)
	})
}
