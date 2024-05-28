package restapi

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/uiansol/task-accounter.git/internal/domain/adapters"
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
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
	taskReadHandler    *handlers.TaskReadHandler
	taskReadAllHandler *handlers.TaskReadAllHandler
	taskUpdateHandler  *handlers.TaskUpdateHandler
	taskDeleteHandler  *handlers.TaskDeleteHandler
}

type AppUseCases struct {
	loginUseCase       usecases.LoginUseCaseInterface
	taskCreateUseCase  usecases.TaskCreateUseCaseInterface
	taskReadUseCase    usecases.TaskReadUseCaseInterface
	taskReadAllUseCase usecases.TaskReadAllUseCaseInterface
	taskUpdateUseCase  usecases.TaskUpdateUseCaseInterface
	taskDeleteUseCase  usecases.TaskDeleteUseCaseInterface
}

type AppRepositories struct {
	userRepository adapters.UserRepositoryInterface
	taskRepository adapters.TaskRepositoryInterface
}

func NewRestService(router *echo.Echo, appHandler *AppHandlers) *RestServer {
	return &RestServer{
		router:     router,
		appHandler: appHandler,
	}
}

func StartServer() {
	LoadEnvs()
	jwtConfig := configJwt()
	db := ConnectToMysql()

	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	repositories := configRepositories(db)
	usecases := configUseCases(repositories)
	handlers := configHandlers(usecases)

	server := NewRestService(router, handlers)
	server.SetUpRoutes(jwtConfig)

	server.router.Start(":8080")
}
