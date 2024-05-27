package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
	"github.com/uiansol/task-accounter.git/internal/domain/mocks"
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/auth"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/dto"
)

func TestNewTaskReadHandler(t *testing.T) {
	t.Run("should return new task read handler", func(t *testing.T) {
		useCaseMock := mocks.NewTaskReadUseCaseInterface(t)

		h := NewTaskReadHandler(useCaseMock)

		assert.NotNil(t, h)
		assert.NotNil(t, h.taskReadUseCase)
	})
}

func TestTaskReadHandle(t *testing.T) {
	t.Run("should process request and return ok with task", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/v2/task/01", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{
			Claims: &auth.JwtCustomClaims{
				ID:    "test-user-id",
				Name:  "test-name",
				Email: "test-email",
				Role:  "technician",
			},
		})

		dateString := "2021-11-22"
		date, _ := time.Parse("2006-01-02", dateString)

		useCaseMock := mocks.NewTaskReadUseCaseInterface(t)
		useCaseMock.On("Execute", mock.Anything).Return(usecases.TaskReadOutput{
			Task: &entities.Task{
				ID:      "123",
				Title:   "title",
				Summary: "summary",
				OwnerID: "test-user-id",
				Status:  entities.TaskStatus("Open"),
				DoneAt:  date,
			},
		}, nil)

		h := NewTaskReadHandler(useCaseMock)
		err := h.Handle(c)

		var res dto.TaskReadResponse
		json.NewDecoder(rec.Body).Decode(&res)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "123", res.ID)
		assert.Equal(t, "title", res.Title)
		assert.Equal(t, "summary", res.Summary)
		assert.Equal(t, "test-user-id", res.OwerID)
		assert.Equal(t, "Open", res.Status)
	})

	t.Run("should process request and return internal server error when use case fails", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/v2/task/01", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{
			Claims: &auth.JwtCustomClaims{
				ID:    "test-user-id",
				Name:  "test-name",
				Email: "test-email",
				Role:  "technician",
			},
		})

		useCaseMock := mocks.NewTaskReadUseCaseInterface(t)
		useCaseMock.On("Execute", mock.Anything).Return(usecases.TaskReadOutput{}, assert.AnError)

		h := NewTaskReadHandler(useCaseMock)
		err := h.Handle(c)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
