package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	t.Run("should create a task and return it with status open", func(t *testing.T) {
		task := NewTask("test task", "test task summary", User{})

		assert.NotNil(t, task.ID)
		assert.Equal(t, "test task", task.Title)
		assert.Equal(t, "test task summary", task.Summary)
		assert.Equal(t, User{}, task.Owner)
		assert.Equal(t, Open, task.Status)
	})
}
