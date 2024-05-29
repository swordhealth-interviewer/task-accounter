package restapi

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/uiansol/task-accounter.git/internal/domain/adapters"
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
	dbmysql "github.com/uiansol/task-accounter.git/internal/infrastructure/db/mysql"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/encrypt"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/redis"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/auth"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/handlers"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func configEncrypter() adapters.EncrypterInterface {
	return encrypt.NewEncrypterService(os.Getenv("SUMMARY_SECRET"))
}

func configPublisher(redisClient redis.Redis) adapters.MessagePublisherInterface {
	return redis.NewMessagePublisher(redisClient)
}

func configJwt() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
}

func ConnectToMysql(debug bool) *gorm.DB {
	var err error
	var dsn string

	if debug {
		dsn = fmt.Sprintf("root:%s@tcp(%s:3306)/%s?parseTime=true", os.Getenv("MYSQL_ROOT_PASSWORD"), "127.0.0.1", os.Getenv("MYSQL_DATABASE"))
	} else {
		dsn = fmt.Sprintf("root:%s@tcp(%s:3306)/%s?parseTime=true", os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_ROOT_HOST"), os.Getenv("MYSQL_DATABASE"))
	}

	db, err := gorm.Open(gormmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting with database")
	}

	return db
}

func ConnectToRedis() redis.Redis {
	return redis.NewRedis()
}

func configHandlers(usecases *AppUseCases) *AppHandlers {
	loginHandler := handlers.NewLoginHandler(usecases.loginUseCase)
	pingHandler := handlers.NewPingHandler()
	taskCreateHandler := handlers.NewTaskCreateHandler(usecases.taskCreateUseCase)
	taskReadHandler := handlers.NewTaskReadHandler(usecases.taskReadUseCase)
	taskReadAllHandler := handlers.NewTaskReadAllHandler(usecases.taskReadAllUseCase)
	taskUpdateHandler := handlers.NewTaskUpdateHandler(usecases.taskUpdateUseCase)
	taskDeleteHandler := handlers.NewTaskDeleteHandler(usecases.taskDeleteUseCase)

	return &AppHandlers{
		loginHandler:       loginHandler,
		pingHandler:        pingHandler,
		taskCreateHandler:  taskCreateHandler,
		taskReadHandler:    taskReadHandler,
		taskReadAllHandler: taskReadAllHandler,
		taskUpdateHandler:  taskUpdateHandler,
		taskDeleteHandler:  taskDeleteHandler,
	}
}

func configUseCases(repositories *AppRepositories, encrypter adapters.EncrypterInterface, publisher adapters.MessagePublisherInterface) *AppUseCases {
	loginUseCase := usecases.NewLoginUseCase(repositories.userRepository)
	taskCreateUseCase := usecases.NewTaskCreateUseCase(repositories.taskRepository, encrypter)
	taskReadUsecase := usecases.NewTaskReadUseCase(repositories.taskRepository, encrypter)
	taskReadAllUsecase := usecases.NewTaskReadAllUseCase(repositories.taskRepository, encrypter)
	taskUpdateUsecase := usecases.NewTaskUpdateUseCase(repositories.taskRepository, encrypter, publisher)
	taskDeleteUsecase := usecases.NewTaskDeleteUseCase(repositories.taskRepository)

	return &AppUseCases{
		loginUseCase:       loginUseCase,
		taskCreateUseCase:  taskCreateUseCase,
		taskReadUseCase:    taskReadUsecase,
		taskReadAllUseCase: taskReadAllUsecase,
		taskUpdateUseCase:  taskUpdateUsecase,
		taskDeleteUseCase:  taskDeleteUsecase,
	}
}

func configRepositories(db *gorm.DB) *AppRepositories {
	userRepository := dbmysql.NewUserRepository(db)
	taskRepository := dbmysql.NewTaskRepository(db)

	return &AppRepositories{
		userRepository: userRepository,
		taskRepository: taskRepository,
	}
}
