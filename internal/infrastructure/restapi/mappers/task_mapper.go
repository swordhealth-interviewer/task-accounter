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
		ID: t.TaskID,
	}
}

func TaskIDRequestToTaskReadInput(t dto.TaskIDRequest, user entities.User) usecases.TaskReadInput {
	return usecases.TaskReadInput{
		ID:   t.ID,
		User: user,
	}
}

func TaskReadOutputToTaskReadResponse(t usecases.TaskReadOutput) dto.TaskReadResponse {
	return dto.TaskReadResponse{
		ID:      t.Task.ID,
		Title:   t.Task.Title,
		Summary: t.Task.Summary,
		OwerID:  t.Task.OwnerID,
		Status:  string(t.Task.Status),
		DoneAt:  t.Task.DoneAt,
	}
}

func TaskReadAllRequestToTaskReadAllInput(user entities.User) usecases.TaskReadAllInput {
	return usecases.TaskReadAllInput{
		User: user,
	}
}

func TaskReadAllOutputToTaskReadResponse(t usecases.TaskReadAllOutput) dto.TaskReadAllResponse {
	var tasks []dto.TaskReadResponse
	for _, task := range t.Tasks {
		tasks = append(tasks, dto.TaskReadResponse{
			ID:      task.ID,
			Title:   task.Title,
			Summary: task.Summary,
			OwerID:  task.OwnerID,
			Status:  string(task.Status),
			DoneAt:  task.DoneAt,
		})
	}
	return dto.TaskReadAllResponse{
		Tasks: tasks,
	}
}

func TaskUpdateRequestToTaskUpdateInput(t dto.TaskUpdateRequest, user entities.User) usecases.TaskUpdateInput {
	return usecases.TaskUpdateInput{
		TaskID:    t.ID,
		Title:     t.Title,
		Summary:   t.Summary,
		CloseTask: t.CloseTask,
		User:      user,
	}
}

func TaskIDRequestToTaskDeleteInput(t dto.TaskIDRequest, user entities.User) usecases.TaskDeleteInput {
	return usecases.TaskDeleteInput{
		TaskID: t.ID,
		User:   user,
	}
}
