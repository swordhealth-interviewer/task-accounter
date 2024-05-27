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

func TestNewTaskReadAllHandler(t *testing.T) {
	t.Run("should return new task read all handler", func(t *testing.T) {
		useCaseMock := mocks.NewTaskReadAllUseCaseInterface(t)

		h := NewTaskReadAllHandler(useCaseMock)

		assert.NotNil(t, h)
		assert.NotNil(t, h.taskReadAllUseCase)
	})
}

func TestTaskReadAllHandle(t *testing.T) {
	t.Run("should process request and return ok with tasks", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/v2/task", nil)
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

		useCaseMock := mocks.NewTaskReadAllUseCaseInterface(t)
		useCaseMock.On("Execute", mock.Anything).Return(usecases.TaskReadAllOutput{
			Tasks: []*entities.Task{
				{
					ID:      "123",
					Title:   "title",
					Summary: "summary",
					OwnerID: "test-user-id",
					Status:  entities.TaskStatus("Open"),
					DoneAt:  date,
				},
				{
					ID:      "456",
					Title:   "title2",
					Summary: "summary2",
					OwnerID: "test-user-id",
					Status:  entities.TaskStatus("Closed"),
					DoneAt:  date,
				},
			},
		}, nil)

		h := NewTaskReadAllHandler(useCaseMock)
		err := h.Handle(c)

		var res dto.TaskReadAllResponse
		json.NewDecoder(rec.Body).Decode(&res)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, 2, len(res.Tasks))

		assert.Equal(t, "123", res.Tasks[0].ID)
		assert.Equal(t, "title", res.Tasks[0].Title)
		assert.Equal(t, "summary", res.Tasks[0].Summary)
		assert.Equal(t, "test-user-id", res.Tasks[0].OwerID)
		assert.Equal(t, "Open", res.Tasks[0].Status)

		assert.Equal(t, "456", res.Tasks[1].ID)
		assert.Equal(t, "title2", res.Tasks[1].Title)
		assert.Equal(t, "summary2", res.Tasks[1].Summary)
		assert.Equal(t, "test-user-id", res.Tasks[1].OwerID)
		assert.Equal(t, "Closed", res.Tasks[1].Status)
	})

	t.Run("should process request and return internal server error when use case fails", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/v2/task", nil)
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

		useCaseMock := mocks.NewTaskReadAllUseCaseInterface(t)
		useCaseMock.On("Execute", mock.Anything).Return(usecases.TaskReadAllOutput{}, assert.AnError)

		h := NewTaskReadAllHandler(useCaseMock)
		err := h.Handle(c)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
