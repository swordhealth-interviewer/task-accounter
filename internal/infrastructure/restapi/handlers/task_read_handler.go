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

type TaskReadHandler struct {
	taskReadUseCase usecases.TaskReadUseCaseInterface
}

func NewTaskReadHandler(TaskReadUseCase usecases.TaskReadUseCaseInterface) *TaskReadHandler {
	return &TaskReadHandler{
		taskReadUseCase: TaskReadUseCase,
	}
}

func (h *TaskReadHandler) Handle(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*auth.JwtCustomClaims)

	user := entities.User{
		ID:    claims.ID,
		Name:  claims.Name,
		Email: claims.Email,
		Role:  entities.UserRole(claims.Role),
	}

	var request dto.TaskIDRequest
	if err := c.Bind(&request); err != nil {
		return c.String(http.StatusBadRequest, badRequestMessage)
	}

	input := mappers.TaskIDRequestToTaskReadInput(request, user)
	output, err := h.taskReadUseCase.Execute(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := mappers.TaskReadOutputToTaskReadResponse(output)

	return c.JSON(http.StatusOK, response)
}
