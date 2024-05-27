package handlers

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/auth"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/dto"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/mappers"
)

type TaskUpdateHandler struct {
	taskUpdateUseCase usecases.TaskUpdateUseCaseInterface
}

func NewTaskUpdateHandler(taskUpdateUseCase usecases.TaskUpdateUseCaseInterface) *TaskUpdateHandler {
	return &TaskUpdateHandler{
		taskUpdateUseCase: taskUpdateUseCase,
	}
}

func (h *TaskUpdateHandler) Handle(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*auth.JwtCustomClaims)

	user := entities.User{
		ID:    claims.ID,
		Name:  claims.Name,
		Email: claims.Email,
		Role:  entities.UserRole(claims.Role),
	}

	var task dto.TaskUpdateRequest
	if err := c.Bind(&task); err != nil {
		return c.String(http.StatusBadRequest, badRequestMessage)
	}

	input := mappers.TaskUpdateRequestToTaskUpdateInput(task, user)
	err := h.taskUpdateUseCase.Execute(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
