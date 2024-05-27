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

type TaskCreateHandler struct {
	taskCreateUseCase usecases.TaskCreateUseCaseInterface
}

func NewTaskCreateHandler(taskCreateUseCase usecases.TaskCreateUseCaseInterface) *TaskCreateHandler {
	return &TaskCreateHandler{
		taskCreateUseCase: taskCreateUseCase,
	}
}

func (h *TaskCreateHandler) Handle(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*auth.JwtCustomClaims)

	user := entities.User{
		ID:    claims.ID,
		Name:  claims.Name,
		Email: claims.Email,
		Role:  entities.UserRole(claims.Role),
	}

	var task dto.TaskCreateRequest
	if err := c.Bind(&task); err != nil {
		return c.String(http.StatusBadRequest, badRequestMessage)
	}

	input := mappers.TaskCreateRequestToTaskCreateInput(task, user)
	output, err := h.taskCreateUseCase.Execute(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := mappers.TaskCreateOutputToTaskCreateResponse(output)

	return c.JSON(http.StatusCreated, response)
}
