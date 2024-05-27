package restapi

import (
	"github.com/uiansol/task-accounter.git/internal/domain/usecases"
	dbmysql "github.com/uiansol/task-accounter.git/internal/infrastructure/db/mysql"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/handlers"
)

func configHandlers(usecases *AppUseCases) *AppHandlers {
	loginHandler := handlers.NewLoginHandler()
	pingHandler := handlers.NewPingHandler()
	taskCreateHandler := handlers.NewTaskCreateHandler(usecases.taskCreateUseCase)
	taskReadAllHandler := handlers.NewTaskReadAllHandler(usecases.taskReadAllUseCase)

	return &AppHandlers{
		loginHandler:       loginHandler,
		pingHandler:        pingHandler,
		taskCreateHandler:  taskCreateHandler,
		taskReadAllHandler: taskReadAllHandler,
	}
}

func configUseCases(repositories *AppRepositories) *AppUseCases {
	taskCreateUseCase := usecases.NewTaskCreateUseCase(repositories.taskRepository)
	taskReadAllUsecase := usecases.NewTaskReadAllUseCase(repositories.taskRepository)

	return &AppUseCases{
		taskCreateUseCase:  taskCreateUseCase,
		taskReadAllUseCase: taskReadAllUsecase,
	}
}

func configRepositories() *AppRepositories {
	taskRepository := dbmysql.NewTaskRepository()

	return &AppRepositories{
		taskRepository: taskRepository,
	}
}
