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
	"github.com/uiansol/task-accounter.git/internal/domain/mocks"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/auth"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/dto"
)

func TestNewTaskUpdateHandler(t *testing.T) {
	t.Run("should return new task update handler", func(t *testing.T) {
		useCaseMock := mocks.NewTaskUpdateUseCaseInterface(t)

		h := NewTaskUpdateHandler(useCaseMock)

		assert.NotNil(t, h)
		assert.NotNil(t, h.taskUpdateUseCase)
	})
}

func TestTaskUpdateHandle(t *testing.T) {
	t.Run("should process request and return ok", func(t *testing.T) {
		validBody := dto.TaskUpdateRequest{
			ID:        "01",
			Title:     "title",
			Summary:   "summary",
			CloseTask: false,
		}
		validBodyJSON, _ := json.Marshal(validBody)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/v2/task:01", strings.NewReader(string(validBodyJSON)))
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

		useCaseMock := mocks.NewTaskUpdateUseCaseInterface(t)
		useCaseMock.On("Execute", mock.Anything).Return(nil)

		h := NewTaskUpdateHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return bad request when body is invalid", func(t *testing.T) {
		invalidBody := "{"
		invalidBodyJSON := []byte(invalidBody)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/v2/task:01", strings.NewReader(string(invalidBodyJSON)))
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

		useCaseMock := mocks.NewTaskUpdateUseCaseInterface(t)

		h := NewTaskUpdateHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return internal server error when use case returns error", func(t *testing.T) {
		validBody := dto.TaskUpdateRequest{
			ID:        "01",
			Title:     "title",
			Summary:   "summary",
			CloseTask: false,
		}
		validBodyJSON, _ := json.Marshal(validBody)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/v2/task:01", strings.NewReader(string(validBodyJSON)))
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

		useCaseMock := mocks.NewTaskUpdateUseCaseInterface(t)
		useCaseMock.On("Execute", mock.Anything).Return(assert.AnError)

		h := NewTaskUpdateHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
