package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/auth"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/dto"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/mappers"
)

type LoginHandler struct {
	loginUseCase usecases.LoginUseCaseInterface
}

func NewLoginHandler(loginUseCase usecases.LoginUseCaseInterface) *LoginHandler {
	return &LoginHandler{
		loginUseCase: loginUseCase,
	}

}

func (h *LoginHandler) Handle(c echo.Context) error {
	var authInput dto.AuthInput
	if err := c.Bind(&authInput); err != nil {
		return c.String(http.StatusBadRequest, badRequestMessage+err.Error())
	}

	input := mappers.MapAuthInputToLoginUseCase(authInput)
	output, err := h.loginUseCase.Execute(input)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	claims := &auth.JwtCustomClaims{
		output.User.ID,
		output.User.Name,
		output.User.Email,
		string(output.User.Role),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
