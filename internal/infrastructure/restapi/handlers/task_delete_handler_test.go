package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/auth"
	"github.com/uiansol/task-accounter.git/mocks"
)

func TestNewTaskDeleteHandler(t *testing.T) {
	t.Run("should return new task delete handler", func(t *testing.T) {
		useCaseMock := mocks.NewTaskDeleteUseCaseInterface(t)

		h := NewTaskDeleteHandler(useCaseMock)

		assert.NotNil(t, h)
		assert.NotNil(t, h.taskDeleteUseCase)
	})
}

func TestTaskDeleteHandle(t *testing.T) {
	t.Run("should process request and return ok", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/v2/task/01", nil)
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

		useCaseMock := mocks.NewTaskDeleteUseCaseInterface(t)
		useCaseMock.On("Execute", mock.Anything).Return(nil)

		h := NewTaskDeleteHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return internal server error when use case returns error", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/v2/task/01", nil)
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

		useCaseMock := mocks.NewTaskDeleteUseCaseInterface(t)
		useCaseMock.On("Execute", mock.Anything).Return(assert.AnError)

		h := NewTaskDeleteHandler(useCaseMock)
		h.Handle(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
