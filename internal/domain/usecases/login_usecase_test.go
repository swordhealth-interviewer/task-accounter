package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
	"github.com/uiansol/task-accounter.git/internal/domain/mocks"
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
)

func TestNewLoginUseCase(t *testing.T) {
	t.Run("should return a login use case", func(t *testing.T) {
		userRepositoryMock := mocks.NewUserRepositoryInterface(t)
		loginUsecase := usecases.NewLoginUseCase(userRepositoryMock)

		assert.NotNil(t, loginUsecase)
		assert.NotNil(t, loginUsecase.UserRepository)
	})
}

func TestLoginUseCaseExecute(t *testing.T) {
	t.Run("should return an error when user is not found", func(t *testing.T) {
		userRepositoryMock := mocks.NewUserRepositoryInterface(t)
		loginUsecase := usecases.NewLoginUseCase(userRepositoryMock)

		input := usecases.LoginInput{
			Username: "test-user",
			Password: "password",
		}

		userRepositoryMock.On("FindByUsername", input.Username).Return(nil, "", assert.AnError)

		output, err := loginUsecase.Execute(input)

		assert.NotNil(t, err)
		assert.Equal(t, usecases.LoginOutput{}, output)
		assert.Equal(t, string(usecases.ErrUserNotFound), err.Error())
	})

	t.Run("should return an error when password is invalid", func(t *testing.T) {
		userRepositoryMock := mocks.NewUserRepositoryInterface(t)
		loginUsecase := usecases.NewLoginUseCase(userRepositoryMock)

		input := usecases.LoginInput{
			Username: "test-user",
			Password: "password",
		}

		userRepositoryMock.On("FindByUsername", input.Username).Return(&entities.User{
			ID:    "user-id",
			Name:  "test-user",
			Email: "test@test.com",
			Role:  entities.UserRoleTechnician,
		}, "invalid", nil)

		output, err := loginUsecase.Execute(input)

		assert.NotNil(t, err)
		assert.Equal(t, usecases.LoginOutput{}, output)
		assert.Equal(t, string(usecases.ErrInvalidCredentials), err.Error())
	})
}
