package entities

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	t.Run("should create a task and return it with status open", func(t *testing.T) {
		task, err := NewTask("test task", "test task summary", "test owner")

		assert.Nil(t, err)
		assert.NotNil(t, task.ID)
		assert.Equal(t, "test task", task.Title)
		assert.Equal(t, "test task summary", task.Summary)
		assert.Equal(t, "test owner", task.Owner)
		assert.Equal(t, Open, task.Status)
	})

	t.Run("should return when do not validate parameters", func(t *testing.T) {
		_, err := NewTask("", "test task summary", "test owner")

		assert.NotNil(t, err)
	})
}

func TestValidateTaskParameters(t *testing.T) {
	t.Run("should return an error when title is empty", func(t *testing.T) {
		err := ValidateTaskParameters("", "test task summary")

		assert.NotNil(t, err)
		assert.Equal(t, errorTitleRequired, err.Error())
	})

	t.Run("should return an error when summary is empty", func(t *testing.T) {
		err := ValidateTaskParameters("test task", "")

		assert.NotNil(t, err)
		assert.Equal(t, errorSummaryRequired, err.Error())
	})

	t.Run("should return an error when summary is longer than 2500 characters", func(t *testing.T) {
		err := ValidateTaskParameters("test task", strings.Repeat("a", 2501))

		assert.NotNil(t, err)
		assert.Equal(t, errorSummaryMaxLength, err.Error())
	})
}
