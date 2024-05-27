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

func configUseCases(repositories *AppRepositories) *AppUseCases {
	taskCreateUseCase := usecases.NewTaskCreateUseCase(repositories.taskRepository)
	taskReadUsecase := usecases.NewTaskReadUseCase(repositories.taskRepository)
	taskReadAllUsecase := usecases.NewTaskReadAllUseCase(repositories.taskRepository)
	taskUpdateUsecase := usecases.NewTaskUpdateUseCase(repositories.taskRepository)
	taskDeleteUsecase := usecases.NewTaskDeleteUseCase(repositories.taskRepository)

	return &AppUseCases{
		taskCreateUseCase:  taskCreateUseCase,
		taskReadUseCase:    taskReadUsecase,
		taskReadAllUseCase: taskReadAllUsecase,
		taskUpdateUseCase:  taskUpdateUsecase,
		taskDeleteUseCase:  taskDeleteUsecase,
	}
}

func configRepositories() *AppRepositories {
	taskRepository := dbmysql.NewTaskRepository()

	return &AppRepositories{
		taskRepository: taskRepository,
	}
}
