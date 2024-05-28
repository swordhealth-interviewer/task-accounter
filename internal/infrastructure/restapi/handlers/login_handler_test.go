package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
	"github.com/uiansol/task-accounter.git/internal/domain/mocks"
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/dto"
)

func TestNewLoginHandler(t *testing.T) {
	t.Run("should return new login handler", func(t *testing.T) {
		useCaseMock := mocks.NewLoginUseCaseInterface(t)

		h := NewLoginHandler(useCaseMock)

		assert.NotNil(t, h)
		assert.NotNil(t, h.loginUseCase)
	})
}

func TestLoginHandle(t *testing.T) {
	t.Run("should process request and return ok with token", func(t *testing.T) {
		validBody := dto.AuthInput{
			Username: "test-username",
			Password: "test-password",
		}
		validBodyJSON, _ := json.Marshal(validBody)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/v1/login", strings.NewReader(string(validBodyJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := mocks.NewLoginUseCaseInterface(t)
		useCaseMock.On("Execute", mock.Anything).Return(usecases.LoginOutput{
			User: &entities.User{
				ID:    "test-user-id",
				Name:  "test-name",
				Email: "test-email",
				Role:  entities.UserRoleTechnician,
			},
		}, nil)
		h := NewLoginHandler(useCaseMock)

		err := h.Handle(c)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	})

	t.Run("should return bad request when body is invalid", func(t *testing.T) {
		invalidBody := "{"
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/v1/login", strings.NewReader(invalidBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := mocks.NewLoginUseCaseInterface(t)
		h := NewLoginHandler(useCaseMock)

		err := h.Handle(c)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	})

	t.Run("should return unauthorized when use case returns error", func(t *testing.T) {
		validBody := dto.AuthInput{
			Username: "test-username",
			Password: "test-password",
		}
		validBodyJSON, _ := json.Marshal(validBody)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/v1/login", strings.NewReader(string(validBodyJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		useCaseMock := mocks.NewLoginUseCaseInterface(t)
		useCaseMock.On("Execute", mock.Anything).Return(usecases.LoginOutput{}, assert.AnError)
		h := NewLoginHandler(useCaseMock)

		err := h.Handle(c)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})
}
