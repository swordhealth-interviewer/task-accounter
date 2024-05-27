package mappers

import (
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/dto"
)

func TaskCreateRequestToTaskCreateInput(t dto.TaskCreateRequest, user entities.User) usecases.TaskCreateInput {
	return usecases.TaskCreateInput{
		Title:   t.Title,
		Summary: t.Summary,
		User:    user,
	}
}

func TaskCreateOutputToTaskCreateResponse(t usecases.TaskCreateOutput) dto.TaskCreateResponse {
	return dto.TaskCreateResponse{
		ID: t.Task.ID,
	}
}
