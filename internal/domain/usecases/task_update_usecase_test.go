package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
