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
	taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)

	t.Run("should return a task create use case", func(t *testing.T) {
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
				ID: "user-id",
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

	t.Run("should return an error when task repository fails to save", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskCreateUsecase := usecases.NewTaskCreateUseCase(taskRepositoryMock)

		taskRepositoryMock.On("Save", mock.Anything).Return(entities.Task{}, assert.AnError)

		output, err := taskCreateUsecase.Execute(usecases.TaskCreateInput{})

		assert.NotNil(t, err)
		assert.Equal(t, "usecase error saving task: assert.AnError general error for testing", err.Error())
		assert.Equal(t, usecases.TaskCreateOutput{}, output)
	})
}
