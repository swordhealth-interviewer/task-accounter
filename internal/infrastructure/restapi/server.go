package restapi

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uiansol/task-accounter.git/internal/domain/adapters"
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
	dbmysql "github.com/uiansol/task-accounter.git/internal/infrastructure/db/mysql"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/handlers"
	"golang.org/x/crypto/bcrypt"
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

func StartServer(debug bool) {
	if debug {
		LoadEnvs()
	}

	jwtConfig := configJwt()
	db := ConnectToMysql(debug)
	redisClient := ConnectToRedis()

	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	publisher := configPublisher(redisClient)
	encrypter := configEncrypter()
	repositories := configRepositories(db)
	usecases := configUseCases(repositories, encrypter, publisher)
	handlers := configHandlers(usecases)

	server := NewRestService(router, handlers)
	server.SetUpRoutes(jwtConfig)

	server.router.Start(":8080")
}

func MigrateDB(debug bool) {
	db := ConnectToMysql(debug)
	db.AutoMigrate(&dbmysql.User{}, &dbmysql.Task{})
}

func PopulateDB(debug bool) {
	db := ConnectToMysql(debug)

	users := []string{"tech-1", "tech-2", "manager-1"}
	passwords := []string{"tech-1", "tech-2", "manager-1"}
	roles := []string{"technician", "technician", "manager"}

	for i := 0; i < 3; i++ {
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(passwords[i]), bcrypt.DefaultCost)

		user := dbmysql.User{
			ID:       uuid.New(),
			Username: users[i],
			Email:    users[i] + "@task-accounter.com",
			Role:     roles[i],
			Password: string(passwordHash),
		}

		db.Create(&user)
	}
}
