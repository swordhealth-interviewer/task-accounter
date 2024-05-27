package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

func TestNewTaskCreateHandler(t *testing.T) {
	t.Run("should return new product create handler", func(t *testing.T) {
		useCaseMock := mocks.NewTaskCreateUseCaseInterface(t)

		h := NewTaskCreateHandler(useCaseMock)

		assert.NotNil(t, h)
		assert.NotNil(t, h.taskCreateUseCase)
	})
}

func TestTaskCreateHandle(t *testing.T) {
	t.Run("should process request and return ok with task id", func(t *testing.T) {
		validBody := dto.TaskCreateRequest{
			Title:   "title",
			Summary: "summary",
		}
		validBodyJSON, _ := json.Marshal(validBody)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/v2/task", strings.NewReader(string(validBodyJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
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

		useCaseMock := mocks.NewTaskCreateUseCaseInterface(t)
		useCaseMock.On("Execute", mock.Anything).Return(usecases.TaskCreateOutput{
			Task: &entities.Task{
				ID: "123",
			},
		}, nil)

		h := NewTaskCreateHandler(useCaseMock)
		h.Handle(c)

		var res dto.TaskCreateResponse
		json.NewDecoder(rec.Body).Decode(&res)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "123", res.ID)
	})

	t.Run("should return bad request when body is invalid", func(t *testing.T) {
		invalidBody := "{"

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/v2/task", strings.NewReader(invalidBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
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

		useCaseMock := mocks.NewTaskCreateUseCaseInterface(t)

		h := NewTaskCreateHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return internal server error when use case returns error", func(t *testing.T) {
		validBody := dto.TaskCreateRequest{
			Title:   "title",
			Summary: "summary",
		}
		validBodyJSON, _ := json.Marshal(validBody)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/v2/task", strings.NewReader(string(validBodyJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
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

		useCaseMock := mocks.NewTaskCreateUseCaseInterface(t)
		useCaseMock.On("Execute", mock.Anything).Return(usecases.TaskCreateOutput{}, assert.AnError)

		h := NewTaskCreateHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
