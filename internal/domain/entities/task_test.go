package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	t.Run("should create a task and return it with status open", func(t *testing.T) {
		task := NewTask("test task", "test task summary", "test owner")

		assert.NotNil(t, task.ID)
		assert.Equal(t, "test task", task.Title)
		assert.Equal(t, "test task summary", task.Summary)
		assert.Equal(t, "test owner", task.Owner)
		assert.Equal(t, Open, task.Status)
	})
}
