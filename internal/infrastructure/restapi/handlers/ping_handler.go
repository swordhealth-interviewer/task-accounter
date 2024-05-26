package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Handle(c echo.Context) error {
	res := map[string]string{
		"message": "pong",
	}

	return c.JSON(http.StatusOK, res)
}
