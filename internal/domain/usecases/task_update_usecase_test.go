package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
	"github.com/uiansol/task-accounter.git/internal/domain/mocks"
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
)

func TestNewTaskUpdateUseCase(t *testing.T) {
	t.Run("should return a task update use case", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskUpdateUsecase := usecases.NewTaskUpdateUseCase(taskRepositoryMock)

		assert.NotNil(t, taskUpdateUsecase)
		assert.NotNil(t, taskUpdateUsecase.TaskRepository)
	})
}

func TestTaskUpdateUseCaseExecute(t *testing.T) {
	t.Run("should update a task and return it with error nil", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskUpdateUsecase := usecases.NewTaskUpdateUseCase(taskRepositoryMock)

		task := usecases.TaskUpdateInput{
			TaskID:    "task-id",
			Title:     "Task Title 2",
			Summary:   "Task Description 2",
			CloseTask: false,
			User: entities.User{
				ID:   "user-id",
				Role: entities.UserRoleTechnician,
			},
		}

		taskRepositoryMock.On("FindByID", task.TaskID).Return(&entities.Task{
			ID:      "task-id",
			Title:   "Task Title",
			Summary: "Task Description",
			OwnerID: "user-id",
			Status:  entities.Open,
		}, nil)

		taskRepositoryMock.On("Save", mock.Anything).Return(&entities.Task{
			ID:      "task-id",
			Title:   "Task Title 2",
			Summary: "Task Description 2",
			OwnerID: "user-id",
			Status:  entities.Open,
		}, nil)

		output, err := taskUpdateUsecase.Execute(task)

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, task.TaskID, output.Task.ID)
		assert.Equal(t, task.Title, output.Task.Title)
		assert.Equal(t, task.Summary, output.Task.Summary)
		assert.Equal(t, task.User.ID, output.Task.OwnerID)
		assert.Equal(t, entities.Open, output.Task.Status)
	})

	t.Run("should close a task and return it with error nil", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskUpdateUsecase := usecases.NewTaskUpdateUseCase(taskRepositoryMock)

		task := usecases.TaskUpdateInput{
			TaskID:    "task-id",
			Title:     "Task Title 2",
			Summary:   "Task Description 2",
			CloseTask: true,
			User: entities.User{
				ID:   "user-id",
				Role: entities.UserRoleTechnician,
			},
		}

		taskRepositoryMock.On("FindByID", task.TaskID).Return(&entities.Task{
			ID:      "task-id",
			Title:   "Task Title",
			Summary: "Task Description",
			OwnerID: "user-id",
			Status:  entities.Open,
		}, nil)

		taskRepositoryMock.On("Save", mock.Anything).Return(&entities.Task{
			ID:      "task-id",
			Title:   "Task Title 2",
			Summary: "Task Description 2",
			OwnerID: "user-id",
			Status:  entities.Closed,
		}, nil)

		output, err := taskUpdateUsecase.Execute(task)

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, task.TaskID, output.Task.ID)
		assert.Equal(t, task.Title, output.Task.Title)
		assert.Equal(t, task.Summary, output.Task.Summary)
		assert.Equal(t, task.User.ID, output.Task.OwnerID)
		assert.Equal(t, entities.Closed, output.Task.Status)
	})

	t.Run("should return an error when user is not a technician", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskUpdateUsecase := usecases.NewTaskUpdateUseCase(taskRepositoryMock)

		task := usecases.TaskUpdateInput{
			User: entities.User{
				Role: entities.UserRoleManager,
			},
		}

		output, err := taskUpdateUsecase.Execute(task)

		assert.NotNil(t, err)
		assert.Equal(t, string(usecases.ErrorTechnicianRoleRequired), err.Error())
		assert.Equal(t, usecases.TaskUpdateOutput{}, output)
	})

	t.Run("should return an error when task is not found", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskUpdateUsecase := usecases.NewTaskUpdateUseCase(taskRepositoryMock)

		task := usecases.TaskUpdateInput{
			TaskID: "task-id",
			User: entities.User{
				ID:   "user-id",
				Role: entities.UserRoleTechnician,
			},
		}

		taskRepositoryMock.On("FindByID", task.TaskID).Return(nil, assert.AnError)

		output, err := taskUpdateUsecase.Execute(task)

		assert.NotNil(t, err)
		assert.Equal(t, string(usecases.ErrorFindTaskByID)+": assert.AnError general error for testing", err.Error())
		assert.Equal(t, usecases.TaskUpdateOutput{}, output)
	})

	t.Run("should return an error when task is not owned by user", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskUpdateUsecase := usecases.NewTaskUpdateUseCase(taskRepositoryMock)

		task := usecases.TaskUpdateInput{
			TaskID: "task-id",
			User: entities.User{
				ID:   "user-id",
				Role: entities.UserRoleTechnician,
			},
		}

		taskRepositoryMock.On("FindByID", task.TaskID).Return(&entities.Task{
			ID:      "task-id",
			Title:   "Task Title",
			Summary: "Task Description",
			OwnerID: "another-user-id",
			Status:  entities.Open,
		}, nil)

		output, err := taskUpdateUsecase.Execute(task)

		assert.NotNil(t, err)
		assert.Equal(t, string(usecases.ErrorTaskNotOwnedByUser), err.Error())
		assert.Equal(t, usecases.TaskUpdateOutput{}, output)
	})

	t.Run("should return an error when task is closed", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskUpdateUsecase := usecases.NewTaskUpdateUseCase(taskRepositoryMock)

		task := usecases.TaskUpdateInput{
			TaskID: "task-id",
			User: entities.User{
				ID:   "user-id",
				Role: entities.UserRoleTechnician,
			},
		}

		taskRepositoryMock.On("FindByID", task.TaskID).Return(&entities.Task{
			ID:      "task-id",
			Title:   "Task Title",
			Summary: "Task Description",
			OwnerID: "user-id",
			Status:  entities.Closed,
		}, nil)

		output, err := taskUpdateUsecase.Execute(task)

		assert.NotNil(t, err)
		assert.Equal(t, string(usecases.ErrorTaskClosed), err.Error())
		assert.Equal(t, usecases.TaskUpdateOutput{}, output)
	})

	t.Run("should return an error when task data is invalid", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskUpdateUsecase := usecases.NewTaskUpdateUseCase(taskRepositoryMock)

		task := usecases.TaskUpdateInput{
			TaskID: "task-id",
			Title:  "Task Title",
			User: entities.User{
				ID:   "user-id",
				Role: entities.UserRoleTechnician,
			},
		}

		taskRepositoryMock.On("FindByID", task.TaskID).Return(&entities.Task{
			ID:      "task-id",
			Title:   "Task Title",
			Summary: "Task Description",
			OwnerID: "user-id",
			Status:  entities.Open,
		}, nil)

		output, err := taskUpdateUsecase.Execute(task)

		assert.NotNil(t, err)
		assert.Equal(t, string(usecases.ErrorInvalidTaskData)+": summary is required", err.Error())
		assert.Equal(t, usecases.TaskUpdateOutput{}, output)
	})

	t.Run("should return an error when task repository save fails", func(t *testing.T) {
		taskRepositoryMock := mocks.NewTaskRepositoryInterface(t)
		taskUpdateUsecase := usecases.NewTaskUpdateUseCase(taskRepositoryMock)

		task := usecases.TaskUpdateInput{
			TaskID:  "task-id",
			Title:   "Task Title",
			Summary: "Task Description",
			User: entities.User{
				ID:   "user-id",
				Role: entities.UserRoleTechnician,
			},
		}

		taskRepositoryMock.On("FindByID", task.TaskID).Return(&entities.Task{
			ID:      "task-id",
			Title:   "Task Title",
			Summary: "Task Description",
			OwnerID: "user-id",
			Status:  entities.Open,
		}, nil)

		taskRepositoryMock.On("Save", mock.Anything).Return(nil, assert.AnError)

		output, err := taskUpdateUsecase.Execute(task)

		assert.NotNil(t, err)
		assert.Equal(t, string(usecases.ErrorSaveTask)+": assert.AnError general error for testing", err.Error())
		assert.Equal(t, usecases.TaskUpdateOutput{}, output)
	})
}
