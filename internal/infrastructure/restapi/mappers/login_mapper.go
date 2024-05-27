package mappers

import (
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/dto"
)

func MapAuthInputToLoginUseCase(input dto.AuthInput) usecases.LoginInput {
	return usecases.LoginInput{
		Username: input.Username,
		Password: input.Password,
	}
}
