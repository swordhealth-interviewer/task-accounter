package restapi

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/uiansol/task-accounter.git/internal/domain/adapters"
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/auth"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/handlers"
)

type RestServer struct {
	router     *echo.Echo
	appHandler *AppHandlers
}

type AppHandlers struct {
	loginHandler       *handlers.LoginHandler
	pingHandler        *handlers.PingHandler
	taskCreateHandler  *handlers.TaskCreateHandler
	taskReadAllHandler *handlers.TaskReadAllHandler
	taskUpdateHandler  *handlers.TaskUpdateHandler
	taskDeleteHandler  *handlers.TaskDeleteHandler
}

type AppUseCases struct {
	taskCreateUseCase  usecases.TaskCreateUseCaseInterface
	taskReadAllUseCase usecases.TaskReadAllUseCaseInterface
	taskUpdateUseCase  usecases.TaskUpdateUseCaseInterface
	taskDeleteUseCase  usecases.TaskDeleteUseCaseInterface
}

type AppRepositories struct {
	taskRepository adapters.TaskRepositoryInterface
}

func NewRestService(router *echo.Echo, appHandler *AppHandlers) *RestServer {
	return &RestServer{
		router:     router,
		appHandler: appHandler,
	}
}

func StartServer() {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}

	repositories := configRepositories()
	usecases := configUseCases(repositories)
	handlers := configHandlers(usecases)

	server := NewRestService(router, handlers)
	server.SetUpRoutes(config)

	server.router.Start(":8080")
}
