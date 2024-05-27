package handlers

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/auth"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/mappers"
)

type TaskReadAllHandler struct {
	taskReadAllUseCase usecases.TaskReadAllUseCaseInterface
}

func NewTaskReadAllHandler(TaskReadAllUseCase usecases.TaskReadAllUseCaseInterface) *TaskReadAllHandler {
	return &TaskReadAllHandler{
		taskReadAllUseCase: TaskReadAllUseCase,
	}
}

func (h *TaskReadAllHandler) Handle(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*auth.JwtCustomClaims)

	user := entities.User{
		ID:    claims.ID,
		Name:  claims.Name,
		Email: claims.Email,
		Role:  entities.UserRole(claims.Role),
	}

	input := mappers.TaskReadAllRequestToTaskReadAllInput(user)
	output, err := h.taskReadAllUseCase.Execute(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := mappers.TaskReadAllOutputToTaskReadResponse(output)

	return c.JSON(http.StatusOK, response)
}
