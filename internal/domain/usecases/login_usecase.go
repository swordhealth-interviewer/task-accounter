package usecases

import (
	"errors"

	"github.com/uiansol/task-accounter.git/internal/domain/adapters"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string
	Password string
}

type LoginOutput struct {
	User *entities.User
}

type LoginUseCaseInterface interface {
	Execute(input LoginInput) (LoginOutput, error)
}

type LoginUseCase struct {
	UserRepository adapters.UserRepositoryInterface
}

func NewLoginUseCase(userRepository adapters.UserRepositoryInterface) *LoginUseCase {
	return &LoginUseCase{
		UserRepository: userRepository,
	}
}

func (uc *LoginUseCase) Execute(input LoginInput) (LoginOutput, error) {
	user, password, err := uc.UserRepository.FindByUsername(input.Username)
	if err != nil {
		return LoginOutput{}, errors.New(string(ErrUserNotFound))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(input.Password)); err != nil {
		return LoginOutput{}, errors.New(string(ErrInvalidCredentials))
	}

	output := LoginOutput{
		User: user,
	}

	return output, nil
}
