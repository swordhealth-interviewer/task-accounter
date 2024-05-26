package restapi

import (
	"github.com/labstack/echo/v4"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/handlers"
)

type RestServer struct {
	router     *echo.Echo
	appHandler *AppHandlers
}

type AppHandlers struct {
	pingHandler *handlers.PingHandler
}

func NewRestService(router *echo.Echo, appHandler *AppHandlers) *RestServer {
	return &RestServer{
		router:     router,
		appHandler: appHandler,
	}
}

func StartServer() {
	router := echo.New()
	handlers := configHandlers()

	server := NewRestService(router, handlers)
	server.SetUpRoutes()

	server.router.Logger.Fatal(router.Start(":8080"))
}
