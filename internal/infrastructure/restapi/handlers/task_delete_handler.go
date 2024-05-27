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

type TaskDeleteHandler struct {
	taskDeleteUseCase usecases.TaskDeleteUseCaseInterface
}

func NewTaskDeleteHandler(taskDeleteUseCase usecases.TaskDeleteUseCaseInterface) *TaskDeleteHandler {
	return &TaskDeleteHandler{
		taskDeleteUseCase: taskDeleteUseCase,
	}
}

func (h *TaskDeleteHandler) Handle(c echo.Context) error {
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

	input := mappers.TaskIDRequestToTaskDeleteInput(request, user)
	err := h.taskDeleteUseCase.Execute(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
